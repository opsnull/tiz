# tiz

构建：
```shell
$ go build -o tiz.app/Contents/MacOS/tiz ./
```

准备：
1. 构建 https://github.com/opsnull/gost，结果命名为 gost_ai:
  + 在官方 gost 的基础上做了修改，起一个本地 OpenAI API 的代理，将请求转发给本地 gost 监听的 HTTP 客户端。
  + 需要将编译后的 gost_ai 放到自己的 PATH 中，如 ~/go/bin。
```shell
git clone git@github.com:opsnull/gost.git
cd gost
go build -o ~/go/bin/gost_ai ./cmd/gost
```

2. 按需修改 gost-bypass.txt 和 tiz.sh, 然后复制到 ~/.ssh 目录下, tiz.app 会执行改脚本来启动 gost_ai 程序.
```shell
cp gost-bypass.txt tiz.sh ~/.ssh
```
 
执行：
1. 双击 tiz.app 执行。
2. 然后就可以在 MacOS 的 systray 菜单栏中看到一个梯子图标, 点击 Enable 即可.

![demo](./demo.gif)
