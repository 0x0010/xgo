# xgo
下载本仓库代码至本地测试的方法
1. 将代码clone到本地
````shell
cd $GOPATH/src
mkdir github.com
cd github.com
mkdir 0x0010
cd 0x0010
git clone https://github.com/0x0010/xgo.git
````
2. 编译打包，以hello为例
````shell
cd $GOPATH/src/github.com/0x0010/xgo/hello
go install
````
安装完之后，在命令行执行hello，将会看到如下内容：
````shell
➜  xgo git:(master) ✗ hello
Hello, world.
````
