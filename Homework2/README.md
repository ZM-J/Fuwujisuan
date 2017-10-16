### CLI 命令行实用程序开发基础

## markdown在线编辑&查看
https://www.zybuluo.com/mdeditor

## 作业任务
设计一个命令行程序selpg，使得能完成选择页面的功能。

## 分析
[教程](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)中给出了实现好了的C程序，我们只需将其翻译成对应的Go程序即可。

大体上Go程序和C程序的语法都是很相似的。对于有出入的地方，只需百度+Google+抱大腿+混即可。

## 测试
将cmd定位至Go程序的所在文件夹下编译Go程序：`go build selpg.go`，会生成相应可执行程序。

1. `selpg.exe -s1 -e1 testfile.txt`

测试效果如下：

**cmd**

![top](https://imgsa.baidu.com/forum/w%3D580/sign=dfcb00b9366d55fbc5c6762e5d234f40/a62afdc9a786c91786623e5ec23d70cf3ac75751.jpg)

![bot](https://imgsa.baidu.com/forum/w%3D580/sign=32d356e4042442a7ae0efdade142ad95/a43d8d45ad3459821e3589f507f431adcaef847f.jpg)

2. `selpg.exe -s1 -e1 < testfile.txt`

测试效果如下：

**cmd**

![top](https://imgsa.baidu.com/forum/w%3D580/sign=6515387069d0f703e6b295d438fa5148/b16b8a1fa8d3fd1fae32cc5d3b4e251f94ca5ffb.jpg)

![bot](https://imgsa.baidu.com/forum/w%3D580/sign=3d846eff55df8db1bc2e7c6c3922dddb/399b2e097bf40ad1033da8845c2c11dfabecced9.jpg)

3. `dir | selpg.exe -s1 -e2`

测试效果如下：

**cmd**

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=0e1fb593bc4543a9f51bfac42e178a7b/0531032fb9389b50f2c48b358e35e5dde6116efb.jpg)

4. `selpg.exe -s10 -e20 testfile.txt >out.txt`

测试效果如下：

**cmd**

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=fb96f7bb8113632715edc23ba18ea056/82d51ffb828ba61e429aa3484a34970a314e5946.jpg)

**out.txt**

![out](https://imgsa.baidu.com/forum/w%3D580/sign=d421259c0146f21fc9345e5bc6256b31/74f09513c8fcc3cec016a2549945d688d63f20e7.jpg)

5. `selpg.exe -s10 -e20 testfile.txt 2>err.txt`

测试效果如下：

**cmd**

![top](https://imgsa.baidu.com/forum/w%3D580/sign=5a15e6c3afc27d1ea5263bcc2bd5adaf/30db77b4c9ea15ceb8f9b46ebd003af33b87b2a7.jpg)

![bot](https://imgsa.baidu.com/forum/w%3D580/sign=d9c6d3469813b07ebdbd50003cd69113/ead3e494d143ad4bf32712a789025aafa50f06a8.jpg)

**err.txt**

![err](https://imgsa.baidu.com/forum/w%3D580/sign=7946b057b1a1cd1105b672288913c8b0/588020328744ebf89bd63058d2f9d72a6159a717.jpg)

6. `selpg.exe -s10 -e20 testfile.txt >out.txt 2>err.txt`

测试效果如下：

**cmd**

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=4a14645f23381f309e198da199004c67/6d3fa1b0cb134954dd18be035d4e9258d0094ab0.jpg)

**out.txt（注意行数是正确的）**

![out](https://imgsa.baidu.com/forum/w%3D580/sign=39665a562a3fb80e0cd161df06d02ffb/2df7a683b9014a90ce8c10daa2773912b21bee0c.jpg)

**err.txt**

![err](https://imgsa.baidu.com/forum/w%3D580/sign=c54d934831292df597c3ac1d8c305ce2/2cd0c525b899a9011dab0d5116950a7b0308f518.jpg)

7. `selpg.exe -s10 -e20 testfile.txt >out.txt 2>nul`

测试效果如下：

**cmd**

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=48aee8470ae9390156028d364bed54f9/b1c6e3fd1e178a828d762fc7fd03738da877e874.jpg)

**out.txt**

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=056e8e93bc4543a9f51bfac42e168a7b/0531032fb9389b50f9b5b0358e35e5dde6116e74.jpg)

8. `selpg.exe -s10 -e20 testfile.txt >nul`

测试效果如下：

![cmd](https://imgsa.baidu.com/forum/w%3D580/sign=6745fe75790e0cf3a0f74ef33a46f23d/e2e096a0cd11728bb007df45c3fcc3cec2fd2c94.jpg)

9. `selpg.exe -s10 -e20 testfile.txt | find "6"`

测试效果如下：

**cmd**

![top](https://imgsa.baidu.com/forum/w%3D580/sign=2f40b37d2334349b74066e8df9eb1521/ee247ddb81cb39dbbc903531db160924aa183061.jpg)

![bot](https://imgsa.baidu.com/forum/w%3D580/sign=87000587fc03918fd7d13dc2613d264b/95fe84d2fd1f4134732d57192e1f95cad0c85ea4.jpg)

10. `selpg.exe -s10 -e20 testfile.txt 2>err.txt | find "7"`

测试效果如下：

**cmd**

![top](https://imgsa.baidu.com/forum/w%3D580/sign=0b9fe5e6c01349547e1ee86c664f92dd/481426d062d9f2d3026c7988a2ec8a136227ccbd.jpg)

![bot](https://imgsa.baidu.com/forum/w%3D580/sign=5877c79472310a55c424defc87444387/12df940f7bec54e799904779b2389b504ec26abe.jpg)

**err.txt**

![err](https://imgsa.baidu.com/forum/w%3D580/sign=3942a0abc8cec3fd8b3ea77de689d4b6/ea19fb2b6059252d547a40633f9b033b5ab5b94d.jpg)

## 总结
从上面的测试中可以看到，我们所编写的selpg能够完成所要求的任务。