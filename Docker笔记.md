

## 基本概念

### 术语

```
host 		宿主机
image 		镜像
container 	容器
registry 	仓库
daemon		守护进程
client 		客户端
```

### 镜像

1. 镜像是只读的，容器在启动的时候创建一层可写层作为最上层。

### 容器

### 仓库

1. 常见仓库

   daocloud

   时速云

   aliyun

2. 使用方式

   ```
   docker search whalesay
   docker pull whalesay
   docker push myname/whalesay
   ```

3. 复制镜像并上传

   ```
   docker tag docker/whalesay clearluo/whalesay
   docker login
   docker push clearluo/whalesay
   ```

   ​

## 安装Docker

### Ubuntu16.04安装

* 阿里脚本自动安装

  ```bash
  curl -sSL http://acs-public-mirror.oss-cn-hangzhou.aliyuncs.com/docker-engine/internet | sh -
  ```

* 阿里镜像加速

  ```
  https://cr.console.aliyun.com/#/accelerator
  ```

  对于使用 systemd 的系统，用  systemctl enable docker  启用服务后，编辑/etc/systemd/system/multi-user.target.wants/docker.service  文件，找到ExecStart=  这一行，在这行最后添加加速器地址  --registry-mirror=<加速器地址>  ，如：

  ```
  ExecStart=/usr/bin/dockerd --registry-mirror=https://jxus37ad.mi
  rror.aliyuncs.com
  ```

  注：对于 1.12 以前的版本， dockerd  换成  docker daemon  。
  重新加载配置并且重新启动。

  ```
  $ sudo systemctl daemon-reload
  $ sudo systemctl restart docker
  ```


## 常见命令

1. 命令列表

   ```dockerfile
   docker pull 	获取image
   docker build 	创建image
   docker images 	列出image
   docker run 		运行container
   docker ps 		列出container
   docker rm 		删除container
   docker rmi 		删除image
   docker cp 		在host和container之间拷贝文件
   docker commit 	保存改动为新的image
   ```

2. 从 Docker Hub 仓库下载一个 Ubuntu 12.04 操作系统的镜像 

   ```dockerfile
   docker pull ubuntu:12.04
   ```

3. 显示本地已有的镜像

   ```dockerfile
   docker images
   ```

4. 交互式进入运行容器

   ```
   docker exec -it nginx /bin/bash
   ```

5. 创建一个容器，让其中运行bash应用

   ```dockerfile
   docker run -t -i ubuntu:12.04 /bin/bash
   ```

6. 查看当前运行docker的control

   ```dockerfile
   docker ps  		列出运行中的容器
   docker ps -a 	列出所有的容器
   ```

7. 运行nginx

   ```dockerfile
   docker run -p 8080:80 -d nginx
   // 其中8080为host端口，80为容器端口
   ```


7. 复制主机文件到容器

   ```dockerfile
   docker cp index.html 17add7bbc58c://usr/share/nginx/html
   ```

8. 保存容器
   ```dockerfile
   docker commit -m 'fun' e7c34d924c31 nginx-fun:3.0
   ```
   ​
## Dockerfile

1. 通过Dockerfile创建镜像

   DockerFile内容

   ```
   FROM alpine:latest
   MAINTAINER clearluo
   CMD echo "Hello Docker"
   ```

   生成镜像命令

   ```
   docker build -t hello_docker .
   ```

   运行镜像

   ```
   docker run hello_docker
   ```

2. Docker语法

   Dockerfile中的每一行都产生一个新层

   ```
    FROM 	base image
    RUN 		执行命令
    ADD 		添加文件
    COPY 		拷贝文件
    CMD 		执行命令
    EXPOSE 	暴露端口
    WORKDIR 	指定路径
    MAINTAINER 维护者
    ENV 		设定环境变量
    ENTRYPOINT 容器入口
    USER 		指定用户
    VOLUME 	mount point
   ```


## Volume

1. 提供独立于容器之外的持久化存储层

2. 指定宿主目录映射到容器目录

   ```
   docker run -d -p 8080:80 -v $PWD/html:/usr/share/nginx/html nginx
   ```

3. 基于数据容器映射容器存储层

   ```
   docker create -v $PWD/data:/var/mydata --name data_container ubuntu
   docker run -it --volumes-from data_container ubuntu /bin/bash​
   ```
## 多容器app

1. Mac/Windows 自带

2. Linux安装

   ```
   Liinux：curl -L https://github.com/docker/compose/releases/download/1.9.0/docker-compose-$(uname -s)-$(uname -m) > /usr/local/bin/docker-compose
   cd /usr/local/bin
   chmod a+x docker-compose
   ```

   ​

























