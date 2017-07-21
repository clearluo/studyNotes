## 安装Consul

* 下载地址：https://www.consul.io/downloads.html
* 解压后只有一个consul可执行文件，放入相应目录下，设置环境变量即可；

## 初体验

* 启动：consul agent -dev
* 查看成员：consul members
* 查看节点：curl 127.0.0.1:8500/v1/catalog/nodes
* 使用DNS协议查看节点信息：dig @127.0.0.1 -p 8600 xx.node.consul

