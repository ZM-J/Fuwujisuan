package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"os/exec"
	"io"
)

type sp_args struct {
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type int
	print_dest string
}

const INBUFSIZ = 16 * 1024

var progname string

func main() {
    sa := sp_args{
		-1,
		-1,
		"",
		72,
		'l',
		"",
	}
	progname = os.Args[0]
	process_args(len(os.Args), os.Args, &sa)
	process_input(sa)
}

func process_args(ac int, av [] string, psa *sp_args) {
	var (
		s1, s2 string
		argno int
	)
	if (ac < 3) {
		fmt.Fprintln(os.Stderr, progname, ": not enough arguments")
		usage()
		os.Exit(1)
	}
	// handle 1st arg - start page
	s1 = av[1]
	if (s1[:2] != "-s") {
		fmt.Fprintln(os.Stderr, progname, ": 1st arg should be -sstart_page")
		usage()
		os.Exit(2)
	}
	i, error := strconv.Atoi(s1[2:])
	if (error != nil || i < 1) {
		fmt.Fprintln(os.Stderr, progname, ": invalid start page ", s1[2:])
		usage()
		os.Exit(3)
	}
	psa.start_page = i
	// handle 2nd arg - end page
	s1 = av[2]
	if (s1[:2] != "-e") {
		fmt.Fprintln(os.Stderr, progname, ": 2nd arg should be -eend_page")
		usage()
		os.Exit(4)
	}
	i, error = strconv.Atoi(s1[2:])
	if (error != nil || i < 1 || i < psa.start_page) {
		fmt.Fprintln(os.Stderr, progname, ": invalid end page ", s1[2:])
		usage()
		os.Exit(5)
	}
	psa.end_page = i
	// handle optional args
	argno = 3
	for (argno < ac && av[argno][0] == '-') {
		s1 = av[argno]
		switch s1[1] {
			case 'l':
				s2 := s1[2:]
				i, error = strconv.Atoi(s2)
				if (error != nil || i < 1) {
					fmt.Fprintln(os.Stderr, progname, ": invalid page length ", s2)
					usage()
					os.Exit(6)
				}
				psa.page_len = i
				argno++
			case 'f':
				if (s1 != "-f") {
					fmt.Fprintln(os.Stderr, progname, ": option should be \"-f\"")
					usage()
					os.Exit(7)
				}
				psa.page_type = 'f'
				argno++
			case 'd':
				s2 = s1[2:]
				if (len(s2) < 1) {
					fmt.Fprintln(os.Stderr, progname, ": -d option requires a printer destination")
					usage()
					os.Exit(8)
				}
				psa.print_dest = s2
				argno++
			default:
				fmt.Fprintln(os.Stderr, progname, ": unknown option ", s1)
				usage()
				os.Exit(9)
		}
	}
	if (argno < ac) {
		psa.in_filename = av[argno]
		f, e := os.Open(psa.in_filename)
		if (e != nil) {
			panic(e)
		}
		defer f.Close()
	}
	// 
	if !(psa.start_page > 0) {
		os.Exit(88)
	}
	if !(psa.end_page > 0 && psa.end_page >= psa.start_page) {
		os.Exit(88)
	}
	if !(psa.page_len > 1) {
		os.Exit(88)
	}
	if !(psa.page_type == 'l' || psa.page_type == 'f') {
		os.Exit(88)
	}
}

func process_input(sa sp_args) {
	fin, fout := os.Stdin, os.Stdout
	if (len(sa.in_filename) != 0) {
		f, e := os.Open(sa.in_filename)
		if (e != nil) {
			panic(e)
		}
		fin = f
		defer f.Close()
	}
	if (len(sa.print_dest) != 0) {
		sf := bufio.NewWriter(os.Stdout)
		sf.Flush()
		cmd := exec.Command("lp", "-d"+sa.print_dest)
		_, err := cmd.Output()
		if (err != nil) {
			fmt.Fprintln(os.Stderr, progname, ": could not open pipe to \"lp -d", sa.print_dest,"\"")
			os.Exit(13)
		}
	}
	rf := bufio.NewReader(fin)
	wf := bufio.NewWriter(fout)
	fe := false
	var line_ctr, page_ctr int
	if (sa.page_type == 'l') {
		line_ctr, page_ctr = 0, 1
		for {
			crc, err := rf.ReadString('\n')
			if (err == io.EOF) {
				break
			} else if (err != nil) {
				fe = true
				break
			}
			line_ctr++
			if (line_ctr > sa.page_len) {
				page_ctr++
				line_ctr = 1
			}
			if (page_ctr >= sa.start_page && page_ctr <= sa.end_page) {
				wf.WriteString(crc)
			}
		}
	} else {
		page_ctr = 1
		for {
			c, _, err := rf.ReadRune()
			if (err == io.EOF) {
				break
			} else if (err != nil) {
				fe = true
				break
			}
			if (c == '\f') {
				page_ctr++
			}
			if (page_ctr >= sa.start_page && page_ctr <= sa.end_page) {
				wf.WriteRune(c)
			}
		}
	}
	if (page_ctr < sa.start_page) {
		fmt.Fprintln(os.Stderr, progname, ": start_page (", sa.start_page,") greater than total pages (", page_ctr,"), no output written")
	} else if (page_ctr < sa.end_page) {
		fmt.Fprintln(os.Stderr, progname, ": end_page (", sa.end_page,") greater than total pages (", page_ctr,"), less output than expected")
	}
	if (fe) {
		fmt.Fprintln(os.Stderr, progname, ": system error occurred on input stream fin")
	} else {
		fin.Close()
		wf.Flush()
		if (len(sa.print_dest) != 0) {
			fout.Close()
		}
		fmt.Fprintln(os.Stderr, progname, ": done")
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "\nUSAGE: ", progname, " -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]")
}
