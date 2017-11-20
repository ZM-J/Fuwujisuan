# 服务计算 Homework3：ex-cloudgo-start

---

## 任务要求
1. 编程 web 服务程序 类似 cloudgo 应用。
* 要求有详细的注释
* 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

## 任务结果
1. 见代码。成功运行后命令行提示如下：
    > PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Fuwujisuan\Homework3> go run main.go
    > [negroni] listening on :8080

2. curl测试命令如下：
    > PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Fuwujisuan\Homework3> curl -v http://localhost:8080/xuming/SharonCheung
    > 详细信息: GET http://localhost:8080/xuming/SharonCheung with 0-byte payload
    > 详细信息: received 67-byte response of content type application/json; charset=UTF-8
    > 
    > 
    > StatusCode        : 200
    > StatusDescription : OK
    > Content           : {
    >                       "PlusOneSecond": "SharonCheung has donated 1s to the elder."
    >                     }
    > 
    > RawContent        : HTTP/1.1 200 OK
    >                     Content-Length: 67
    >                     Content-Type: application/json; charset=UTF-8
    >                     Date: Mon, 20 Nov 2017 09:28:38 GMT
    > 
    >                     {
    >                       "PlusOneSecond": "SharonCheung has donated 1s to the elder."
    >                     }
    > 
    > Forms             : {}
    > Headers           : {[Content-Length, 67], [Content-Type, application/json; charset=UTF-8], [Date, Mon, 20 Nov 2017 09:28:38 GMT]}
    > Images            : {}
    > InputFields       : {}
    > Links             : {}
    > ParsedHtml        : mshtml.HTMLDocumentClass
    > RawContentLength  : 67

    从上面的结果我们能看出，curl能够收到本地服务器的HTTP响应。
3. ab测试命令如下：
    > PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Fuwujisuan\Homework3> ab -n 1000 -c 100 http://localhost:8080/xuming/SunXiaochuan
    > This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
    > Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    > Licensed to The Apache Software Foundation, http://www.apache.org/
    > 
    > Benchmarking localhost (be patient)
    > Completed 100 requests
    > Completed 200 requests
    > Completed 300 requests
    > Completed 400 requests
    > Completed 500 requests
    > Completed 600 requests
    > Completed 700 requests
    > Completed 800 requests
    > Completed 900 requests
    > Completed 1000 requests
    > Finished 1000 requests
    > 
    > 
    > Server Software:
    > Server Hostname:        localhost
    > Server Port:            8080
    > 
    > Document Path:          /xuming/SunXiaochuan
    > Document Length:        67 bytes
    > 
    > Concurrency Level:      100
    > Time taken for tests:   0.198 seconds
    > Complete requests:      1000
    > Failed requests:        0
    > Total transferred:      190000 bytes
    > HTML transferred:       67000 bytes
    > Requests per second:    5047.14 [#/sec] (mean)
    > Time per request:       19.813 [ms] (mean)
    > Time per request:       0.198 [ms] (mean, across all concurrent requests)
    > Transfer rate:          936.48 [Kbytes/sec] received
    > 
    > Connection Times (ms)
    >               min  mean[+/-sd] median   max
    > Connect:        0    0   0.3      0       1
    > Processing:     5   19   7.4     17      51
    > Waiting:        2   16   7.8     15      49
    > Total:          5   19   7.4     17      51
    > 
    > Percentage of the requests served within a certain time (ms)
    >   50%     17
    >   66%     17
    >   75%     18
    >   80%     18
    >   90%     27
    >   95%     40
    >   98%     43
    >   99%     44
    >  100%     51 (longest request)

以上便是ab测试的结果。