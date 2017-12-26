# 服务计算 Homework6：reflect

---

## 预备姿势

在`MySQL Workbench`中，加入以下代码来创建数据库：
```sql
CREATE DATABASE Homework6;

USE Homework6;

CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);
```

## 实现细节

1. orm 规则
   
   每个字段用到了`table`与`column`的属性。

2. 实现自动插入数据

   见代码。

   其中`daoUserInstance.Save(user)`这句话便是实现自动插入数据的体现。

3. 实现查询结果自动映射

   见代码。

   其中`daoUserInstance.Find()`这句话便是实现查询结果自动映射的体现。

## 运行结果

```powershell
PS D:\college\Junior\Fuwujisuan\gp\src\github.com\ZM-J\Fuwujisuan\Homework6> go run main.go
INSERT INTO userinfo (uid,username,departname,created) VALUES (?,?,?,?)   [0 SunXiaoChuan huya 2017-12-26 22:28:35]
result: &{1 SunXiaoChuan huya 2017-12-26 00:00:00 +0000 UTC}
```

在`MySQL Workbench`中，查看表格`UserInfo`的全部内容。

```
# uid, username, departname, created
1, SunXiaoChuan, huya, 2017-12-26
2, SunXiaoChuan, huya, 2017-12-26
```

可以看到，数据项被正确地写入到数据库中了。