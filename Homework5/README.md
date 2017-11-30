# 服务计算 Homework5：ex-cloudgo-data

---

## 预备姿势
1. 配置MySQL服务器

   首先我们找到MySQL的安装包`mysql-5.7.17.msi`并下载安装。MySQL提供了非常方便的Workbench来让我们通过GUI来操作。
   
   之后我们在创建实例`Homework5`的时候，杯具发生了：**Can't get hostname for your address**

   这个网上有部分说法是说在配置文件的`[mysqld]`块中加上一句`skip-name-resolve`。然而实践证明，这并没有什么卵用。后面我找到了[比较科学的解决办法](http://blog.csdn.net/AdolphKevin/article/details/70800880)，按照这解决办法来操作，问题便迎刃而解。

   之后的代码便跟老师的创建表格代码类似：

   ```SQL
   create database Homework5;
   USE Homework5;

   CREATE TABLE IF NOT EXISTS`userinfo` (
       `uid` INT(10) NOT NULL AUTO_INCREMENT,
       `username` VARCHAR(64) NULL DEFAULT NULL,
       `departname` VARCHAR(64) NULL DEFAULT NULL,
       `created` DATE NULL DEFAULT NULL,
       PRIMARY KEY (`uid`)
   );

   CREATE TABLE IF NOT EXISTS `userdetail` (
       `uid` INT(10) NOT NULL DEFAULT '0',
       `intro` TEXT NULL,
       `profile` TEXT NULL,
       PRIMARY KEY (`uid`)
   );
   ```
   在GUI中，我们也能通过鼠标一指禅来看到当前创建好了的表格的状态。


## 任务要求

1. 添加数据

   使用curl来对服务器进行POST请求，进一步来添加数据。

   在PowerShell中，curl命令语法如下：

   ```powershell
   curl [-Uri] <Uri> [-Body <Object> ] [-Certificate <X509Certificate> ] [-CertificateThumbprint <String> ] [-ContentType <String> ] [-Credential <PSCredential> ] [-DisableKeepAlive] [-Headers <IDictionary> ] [-InFile <String> ] [-MaximumRedirection <Int32> ] [-Method <WebRequestMethod> {Default | Get | Head | Post | Put | Delete | Trace | Options | Merge | Patch} ] [-OutFile <String> ] [-PassThru] [-Proxy <Uri> ] [-ProxyCredential <PSCredential> ] [-ProxyUseDefaultCredentials] [-SessionVariable <String> ] [-TimeoutSec <Int32> ] [-TransferEncoding <String> {chunked | compress | deflate | gzip | identity} ] [-UseBasicParsing] [-UseDefaultCredentials] [-UserAgent <String> ] [-WebSession <WebRequestSession> ] [ <CommonParameters>]
   ```
   
   我们在这里主要用到`-Uri`参数、`-Method`参数以及`-Body`参数。

   我们先尝试输入下面命令：
   
   ```powershell
   curl "http://localhost:8080/service/userinfo" -Method Post
   ```

   返回400 HTTP Bad Request，这是因为我们的表单更新的相应数据还尚未输入。

   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=sunxiaochuan&departname=6324"

   我们能在PowerShell中观察到返回的HTTP相应：

   ```powershell
   StatusCode        : 200
   StatusDescription : OK
   Content           : {
                         "UID": 1,
                         "UserName": "sunxiaochuan",
                         "DepartName": "6324",
                         "CreateAt": "2017-11-28T20:52:54.2039473+08:00"
                       }
   
   RawContent        : HTTP/1.1 200 OK
                       Content-Length: 120
                       Content-Type: application/json; charset=UTF-8
                       Date: Tue, 28 Nov 2017 12:52:54 GMT
   
                       {
                         "UID": 1,
                         "UserName": "sunxiaochuan",
                         "DepartName": "6324",
                         "Creat...
   Forms             : {}
   Headers           : {[Content-Length, 120], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 12:52:54 GMT]}
   Images            : {}
   InputFields       : {}
   Links             : {}
   ParsedHtml        : mshtml.HTMLDocumentClass
   RawContentLength  : 120
   ```

   进一步，我们希望能够我们的数据库能支持UTF8文本：

   ```powershell
   curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=孙亚龙&departname=德云色"
   ```

   再次观察相应的HTTP响应，悲剧发生了：

   ```powershell
   StatusCode        : 200
   StatusDescription : OK
   Content           : {
                         "UID": 2,
                         "UserName": "???",
                         "DepartName": "???",
                         "CreateAt": "2017-11-28T20:55:46.5877896+08:00"
                       }
   
   RawContent        : HTTP/1.1 200 OK
                       Content-Length: 110
                       Content-Type: application/json; charset=UTF-8
                       Date: Tue, 28 Nov 2017 12:55:46 GMT
   
                       {
                         "UID": 2,
                         "UserName": "???",
                         "DepartName": "???",
                         "CreateAt": "201...
   Forms             : {}
   Headers           : {[Content-Length, 110], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 12:55:46 GMT]}
   Images            : {}
   InputFields       : {}
   Links             : {}
   ParsedHtml        : mshtml.HTMLDocumentClass
   RawContentLength  : 110
   ```

   我们可以看到，`UserName`和`DepartName`的中文字段均为???，并且在MySQL中，我们也可以看到对应数据确实为???。
   
   后面，搜索到了解决方法：在PowerShell中做好UTF8字段的转码。

   ```powershell
   Add-Type -AssemblyName System.Web
   $R = [System.Web.HttpUtility]::UrlEncode("username=孙亚龙&departname=德云色")
   curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body $R
   ```

   然而这上面的命令还是不能达到目的，原因在于，它这个转码把等于号和与符号都转码掉了，go后端识别不出等于号的话，那么`UserName`字段值仍然为空。
  
   最后的解决方法是部分转码，只对数据部分转码即可。

   ```powershell
   Add-Type -AssemblyName System.Web
   > $Username = [System.Web.HttpUtility]::UrlEncode("孙亚龙")
   > $Departname = [System.Web.HttpUtility]::UrlEncode("德云色")
   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=$Username&departname=$Departname"
   ```

   收到的HTTP响应如下：
   
   ```powershell
   StatusCode        : 200
   StatusDescription : OK
   Content           : {
                         "UID": 3,
                         "UserName": "孙亚龙",
                         "DepartName": "德云色",
                         "CreateAt": "2017-11-28T21:21:11.7275144+08:00"
                       }
   
   RawContent        : HTTP/1.1 200 OK
                       Content-Length: 122
                       Content-Type: application/json; charset=UTF-8
                       Date: Tue, 28 Nov 2017 13:21:11 GMT
   
                       {
                         "UID": 3,
                         "UserName": "孙亚龙",
                         "DepartName": "德云色",
                         "CreateAt": "201...
   Forms             : {}
   Headers           : {[Content-Length, 122], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 13:21:11 GMT]}
   Images            : {}
   InputFields       : {}
   Links             : {}
   ParsedHtml        : mshtml.HTMLDocumentClass
   RawContentLength  : 122
   ```

   这便解决了UTF8的编码问题。这个时候我们用浏览器进入`http://localhost:8080/service/userinfo`（等价于用GET方法），显示如下

   ```json
   [
     {
       "UID": 1,
       "UserName": "sunxiaochuan",
       "DepartName": "6324",
       "CreateAt": "2017-11-28T00:00:00Z"
     },
     {
       "UID": 2,
       "UserName": "???",
       "DepartName": "???",
       "CreateAt": "2017-11-28T00:00:00Z"
     },
     {
       "UID": 3,
       "UserName": "孙亚龙",
       "DepartName": "德云色",
       "CreateAt": "2017-11-28T00:00:00Z"
     }
   ]
   ```

   在MySQL的GUI中也可以看到对应的数据。

2. 数据查询
   
   浏览器输入`http://localhost:8080/service/userinfo?userid=3`，即可看到negroni响应：

   > [negroni] 2017-11-28T23:47:45+08:00 | 400 |      2.001ms | localhost:8080 | GET /service/userinfo

3. 数据服务测试
   
   未完待续

4. 用`xorm`而不是`database/sql`来构建数据服务

   见`xorm`文件夹下的代码。值得注意的是，实体entity

   用浏览器进入`http://localhost:8080/service/userinfo`以及`http://localhost:8080/service/usercount`，也能看到跟之前一样的结果。

5. 比较`xorm`和`database/sql`

   相对于原生库，GO的`xorm`库将数据库的基本操作（增删查改）也封装成了一个`engine`对象。 这样的话，操作数据就不用编写DAO服务了，全藉由`engine`对象所提供的方法来完成就可以了。

   当然，这样做也有弊端。由于`xorm`相当是对SQL的进一步封装，事实上其效率是不如SQL的。

   1. orm是否实现了dao的自动化？
   
   是。

   2. 使用ab测试性能。
   
   测试语句为`ab -n 1000 -c 100 "http://localhost:8080/service/userinfo"`。测试结果如下：

   对于`database/sql`实现版本（用到dao而不是xorm）：

   ```
   This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
   Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
   Licensed to The Apache Software Foundation, http://www.apache.org/
   
   Benchmarking localhost (be patient)
   Completed 100 requests
   Completed 200 requests
   Completed 300 requests
   Completed 400 requests
   Completed 500 requests
   Completed 600 requests
   Completed 700 requests
   Completed 800 requests
   Completed 900 requests
   Completed 1000 requests
   Finished 1000 requests
   
   
   Server Software:
   Server Hostname:        localhost
   Server Port:            8080
   
   Document Path:          /service/userinfo
   Document Length:        355 bytes
   
   Concurrency Level:      100
   Time taken for tests:   1.031 seconds
   Complete requests:      1000
   Failed requests:        0
   Total transferred:      479000 bytes
   HTML transferred:       355000 bytes
   Requests per second:    969.59 [#/sec] (mean)
   Time per request:       103.136 [ms] (mean)
   Time per request:       1.031 [ms] (mean, across all concurrent requests)
   Transfer rate:          453.55 [Kbytes/sec] received
   
   Connection Times (ms)
                 min  mean[+/-sd] median   max
   Connect:        0    0   0.4      0       7
   Processing:    15   97  51.2    119     180
   Waiting:       15   97  51.2    119     178
   Total:         15   97  51.2    119     180
   
   Percentage of the requests served within a certain time (ms)
     50%    119
     66%    134
     75%    139
     80%    145
     90%    159
     95%    164
     98%    168
     99%    170
    100%    180 (longest request)
   ```

   对于`xorm`实现版本：

   ```
   Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
   Licensed to The Apache Software Foundation, http://www.apache.org/
   
   Benchmarking localhost (be patient)
   Completed 100 requests
   Completed 200 requests
   Completed 300 requests
   Completed 400 requests
   Completed 500 requests
   Completed 600 requests
   Completed 700 requests
   Completed 800 requests
   Completed 900 requests
   Completed 1000 requests
   Finished 1000 requests
   
   
   Server Software:
   Server Hostname:        localhost
   Server Port:            8080
   
   Document Path:          /service/userinfo
   Document Length:        367 bytes
   
   Concurrency Level:      100
   Time taken for tests:   1.000 seconds
   Complete requests:      1000
   Failed requests:        0
   Total transferred:      491000 bytes
   HTML transferred:       367000 bytes
   Requests per second:    1000.32 [#/sec] (mean)
   Time per request:       99.968 [ms] (mean)
   Time per request:       1.000 [ms] (mean, across all concurrent requests)
   Transfer rate:          479.64 [Kbytes/sec] received

   Connection Times (ms)
                 min  mean[+/-sd] median   max
   Connect:        0    0   0.2      0       1
   Processing:    25   95  47.5    120     161
   Waiting:       25   94  47.6    120     161
   Total:         25   95  47.5    120     161
   
   Percentage of the requests served within a certain time (ms)
     50%    120
     66%    133
     75%    136
     80%    139
     90%    143
     95%    146
     98%    149
     99%    152
    100%    161 (longest request)
   ```

