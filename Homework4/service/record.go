package service

import (
	"net/http"
	//"fmt"
	"github.com/unrolled/render"
	"regexp"
)

type Item struct {
	Name  string
	Price string
}

func recordHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" { // 如果是GET方法，则退回到之前的网页
			http.Redirect(w, r, "/static/", http.StatusFound)
		} else { // POST方法
			r.ParseForm();
			item := Item{
				Name: r.Form["name"][0],
				Price: r.Form["price"][0],
			}
			if m, _ := regexp.MatchString(`^\d+(\.\d{1,2})?$`, item.Price); !m { // 价格不是整数或小数
				http.Redirect(w, r, "/static/", http.StatusFound)
			}
			formatter.HTML(w,http.StatusOK,"record", item)
		}
	}
}