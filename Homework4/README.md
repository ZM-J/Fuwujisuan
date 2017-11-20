# 服务计算 Homework4：ex-cloudgo-inout

---

## 任务要求
1. 支持静态文件服务
2. 支持简单 js 访问
3. 提交表单，并输出一个表格
4. 对 `/unknown` 给出开发中的提示，返回码 `5xx`

## 任务结果
1. 见代码。成功运行后命令行提示如下：
    > PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Fuwujisuan\Homework4> go run main.go
    > [negroni] listening on :8080

    与老师上课的时候代码类似。用浏览器进入`localhost:8080/static/`能看到到`assets/index.html`网页，用浏览器进入`localhost:8080/static/images/`会显示文件列表，用浏览器进入`localhost:8080/static/images/赴戍登程口占示家人二首.txt`也能看到林则徐的著名诗篇Θ..Θ

2. 见代码。成功运行后命令行提示和上一要求相同。javascript代码通过获取到Go返回的JSON对象来实现DOM对应元素的替换。

    ```HTML
    <div>
        <p class="item-name">物品名：</p>
        <p class="item-price">价格：</p>
    </div>
    ```

    被替换成
    
    ```HTML
    <div>
        <p class="item-name">物品名：Clearlove7</p>
        <p class="item-price">价格：4396</p>
    </div>
    ```

3. 依照老师所给的步骤，在`localhost:8080/static/`设置一个包含物品名和价格的一个表单，通过这个表单来实现表单的POST提交，并跳转到`localhost:8080/record`来实现POST提交的完成。具体逻辑如下：
    在record中显示一个表格形式的模板页面，这个页面位于项目文件夹`/templates/record.tmpl`。
    * 如果进入`localhost:8080/record`的方法为GET，那么这数据提交是不合法的，需重定向到`localhost:8080/static/`。
    * 如果进入`localhost:8080/record`的方法为POST，但是价格不是一个整数或者一个一位小数或者一个两位小数的话（这工作可以藉由正则表达式匹配来完成），那么数据提交也不合法，重定向到`localhost:8080/static/`。
    * 如果进入`localhost:8080/record`的方法为POST，且价格合法的话，那么就将表单数据应用于模板页面并显示即可。

4. 我们首先找到`NotFound`和`NotFoundHandler`的[定义](https://go-zh.org/pkg/net/http/)与[实现](https://go-zh.org/src/net/http/server.go)，并模仿其中写出的`NotImplemented`和`NotImplementedHandler`的实现。在服务器中部署相应路由的时候加上
    ```
    mx.HandleFunc("/api/unknown", NotImplemented)
    ```
    即可。
    访问`localhost:8080/api/unknown`的时候，会显示页面501 Not Implemented，并且终端会显示
    > [negroni] 2017-11-20T23:38:57+08:00 | 501 |      0s | localhost:8080 | GET /api/unknown
