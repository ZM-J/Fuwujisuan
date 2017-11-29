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

   > curl [-Uri] <Uri> [-Body <Object> ] [-Certificate <X509Certificate> ] [-CertificateThumbprint <String> ] [-ContentType <String> ] [-Credential <PSCredential> ] [-DisableKeepAlive] 
   > [-Headers <IDictionary> ] [-InFile <String> ] [-MaximumRedirection <Int32> ] [-Method <WebRequestMethod> {Default | Get | Head | Post | Put | Delete | Trace | Options | Merge | Patch} ]
   > [-OutFile <String> ] [-PassThru] [-Proxy <Uri> ] [-ProxyCredential <PSCredential> ] [-ProxyUseDefaultCredentials] [-SessionVariable <String> ] [-TimeoutSec <Int32> ] [-TransferEncoding 
   > <String> {chunked | compress | deflate | gzip | identity} ] [-UseBasicParsing] [-UseDefaultCredentials] [-UserAgent <String> ] [-WebSession <WebRequestSession> ] [ <CommonParameters>]

   我们在这里主要用到`-Uri`参数、`-Method`参数以及`-Body`参数。

   我们先尝试输入下面命令：
   
   > curl "http://localhost:8080/service/userinfo" -Method Post

   返回400 HTTP Bad Request，这是因为我们的表单更新的相应数据还尚未输入。

   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=sunxiaochuan&departname=6324"

   我们能在PowerShell中观察到返回的HTTP相应：

   > StatusCode        : 200
   > StatusDescription : OK
   > Content           : {
   >                       "UID": 1,
   >                       "UserName": "sunxiaochuan",
   >                       "DepartName": "6324",
   >                       "CreateAt": "2017-11-28T20:52:54.2039473+08:00"
   >                     }
   > 
   > RawContent        : HTTP/1.1 200 OK
   >                     Content-Length: 120
   >                     Content-Type: application/json; charset=UTF-8
   >                     Date: Tue, 28 Nov 2017 12:52:54 GMT
   > 
   >                     {
   >                       "UID": 1,
   >                       "UserName": "sunxiaochuan",
   >                       "DepartName": "6324",
   >                       "Creat...
   > Forms             : {}
   > Headers           : {[Content-Length, 120], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 12:52:54 GMT]}
   > Images            : {}
   > InputFields       : {}
   > Links             : {}
   > ParsedHtml        : mshtml.HTMLDocumentClass
   > RawContentLength  : 120
   
   进一步，我们希望能够我们的数据库能支持UTF8文本：

   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=孙亚龙&departname=德云色"

   再次观察相应的HTTP响应，悲剧发生了：

   > StatusCode        : 200
   > StatusDescription : OK
   > Content           : {
   >                       "UID": 2,
   >                       "UserName": "???",
   >                       "DepartName": "???",
   >                       "CreateAt": "2017-11-28T20:55:46.5877896+08:00"
   >                     }
   > 
   > RawContent        : HTTP/1.1 200 OK
   >                     Content-Length: 110
   >                     Content-Type: application/json; charset=UTF-8
   >                     Date: Tue, 28 Nov 2017 12:55:46 GMT
   > 
   >                     {
   >                       "UID": 2,
   >                       "UserName": "???",
   >                       "DepartName": "???",
   >                       "CreateAt": "201...
   > Forms             : {}
   > Headers           : {[Content-Length, 110], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 12:55:46 GMT]}
   > Images            : {}
   > InputFields       : {}
   > Links             : {}
   > ParsedHtml        : mshtml.HTMLDocumentClass
   > RawContentLength  : 110

   我们可以看到，`UserName`和`DepartName`的中文字段均为???，并且在MySQL中，我们也可以看到对应数据确实为???。
   
   后面，搜索到了解决方法：在PowerShell中做好UTF8字段的转码。

   > Add-Type -AssemblyName System.Web
   > $R = [System.Web.HttpUtility]::UrlEncode("username=孙亚龙&departname=德云色")
   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body $R

   然而这上面的命令还是不能达到目的，原因在于，它这个转码把等于号和与符号都转码掉了，go后端识别不出等于号的话，那么`UserName`字段值仍然为空。
  
   最后的解决方法是部分转码，只对数据部分转码即可。

   > Add-Type -AssemblyName System.Web
   > $Username = [System.Web.HttpUtility]::UrlEncode("孙亚龙")
   > $Departname = [System.Web.HttpUtility]::UrlEncode("德云色")
   > curl -Uri "http://localhost:8080/service/userinfo" -Method Post -Body "username=$Username&departname=$Departname"

   收到的HTTP响应如下：
   
   > StatusCode        : 200
   > StatusDescription : OK
   > Content           : {
   >                       "UID": 3,
   >                       "UserName": "孙亚龙",
   >                       "DepartName": "德云色",
   >                       "CreateAt": "2017-11-28T21:21:11.7275144+08:00"
   >                     }
   > 
   > RawContent        : HTTP/1.1 200 OK
   >                     Content-Length: 122
   >                     Content-Type: application/json; charset=UTF-8
   >                     Date: Tue, 28 Nov 2017 13:21:11 GMT
   > 
   >                     {
   >                       "UID": 3,
   >                       "UserName": "孙亚龙",
   >                       "DepartName": "德云色",
   >                       "CreateAt": "201...
   > Forms             : {}
   > Headers           : {[Content-Length, 122], [Content-Type, application/json; charset=UTF-8], [Date, Tue, 28 Nov 2017 13:21:11 GMT]}
   > Images            : {}
   > InputFields       : {}
   > Links             : {}
   > ParsedHtml        : mshtml.HTMLDocumentClass
   > RawContentLength  : 122

   这便解决了UTF8的编码问题。这个时候我们用浏览器进入`http://localhost:8080/service/userinfo`（等价于用GET方法），显示如下

   > [
   >   {
   >     "UID": 1,
   >     "UserName": "sunxiaochuan",
   >     "DepartName": "6324",
   >     "CreateAt": "2017-11-28T00:00:00Z"
   >   },
   >   {
   >     "UID": 2,
   >     "UserName": "???",
   >     "DepartName": "???",
   >     "CreateAt": "2017-11-28T00:00:00Z"
   >   },
   >   {
   >     "UID": 3,
   >     "UserName": "孙亚龙",
   >     "DepartName": "德云色",
   >     "CreateAt": "2017-11-28T00:00:00Z"
   >   }
   > ]

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