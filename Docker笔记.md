

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

1. 简单的说，容器是独立运行的一个或一组应用，以及它们的运行态环境。对应的，虚拟机可以理解为模拟运行的一整套操作系统（提供了运行态环境和其他系统环境）和跑在上面的应用。 

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

* Docker 是服务器----客户端架构。命令行运行`docker`命令的时候，需要本机有 Docker 服务。如果这项服务没有启动，可以用下面的命令启动

  ```
  # service 命令的用法
  $ sudo service docker start

  # systemctl 命令的用法
  $ sudo systemctl start docker
  ```

* Docker 需要用户具有 sudo 权限，为了避免每次命令都输入`sudo`，可以把用户加入 Docker 用户组

  ```
  $ sudo usermod -aG docker $USER
  ```

  ​


## 常见命令

1. 命令列表

   ```dockerfile
   docker pull 	获取image
   docker build 	创建image
   docker images 	列出image
   docker run 		运行container
   docker ps 		列出container
   docker rm 		删除container
   docker rmi 		删除image,注意：在删除镜像之前要先用 docker rm 删掉依赖于这个镜像的所有容器
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

8. 复制修改镜像标签

   ```
   docker tag 5db5f8471261 ouruser/sinatra:devel
   ```

9. 保存容器

   ```dockerfile
   docker commit -m 'fun' -a "clearluo" e7c34d924c31 nginx-fun:3.0
   ```

10. 存出镜像

  如果要导出镜像到本地文件，可以使用docker save 命令

  ```
  docker save -o ubuntu_14.04.tar ubuntu:14.04
  ```

11. 载入镜像

   可以使用docker load从导出的本地文件中再导入到本地镜像库

   ```
   docker load --input ubuntu_14.04.tar
   或
   docker load < ubuntu_14.04.tar
   ```

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
    #			注释
    FROM 		base image
    RUN 		镜像创建过程中，执行命令
    ADD 		添加文件,ADD指令还支持通过URL从远程服务器读取资源并复制到镜像中.
    COPY 		拷贝文件(推荐),指令只能从执行docker build所在的主机上读取资源并复制到镜像中
    CMD 		容器启动后，执行的命令,Dockerfile 中可以有多个 CMD 指令，但只有最后一个生效。CMD 可以被 docker run 之后的参数替换
    EXPOSE 	暴露端口
    WORKDIR 	指定路径,为后面的 RUN, CMD, ENTRYPOINT, ADD 或 COPY 指令设置镜像中的当前工作目录,也是容器运行后默认目录
    MAINTAINER 维护者
    ENV 		设定环境变量,环境变量可被后面的指令使用
    ENTRYPOINT 设置容器启动时运行的命令。Dockerfile 中可以有多个 ENTRYPOINT 指令，但只有最后一个生效。CMD 或 docker run 之后的参数会被当做参数传递给 ENTRYPOINT, 不会被忽略，一定会被执行
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
## 多容器管理工具docker-compose

1. Mac/Windows 自带

2. Linux安装

   ```
   Liinux：curl -L https://github.com/docker/compose/releases/download/1.9.0/docker-compose-$(uname -s)-$(uname -m) > /usr/local/bin/docker-compose
   cd /usr/local/bin
   chmod a+x docker-compose
   ```

3. docker-compose.yml常用命令
   ```
   build 		本地创建镜像
   command 	覆盖缺省命令
   depends_on	链接容器
   ports 		暴露端口
   volumes 	卷
   image		pull镜像
   ```

4. docker-compose命令

   ```
   up 		启动服务
   stp		停止服务
   rm		删除服务中的各个容器
   logs	观察各个容器的日志
   ps		列出服务相关的容器
   ```

   ​
## 实际应用

### Mysql使用

```dockerfile
docker run -d -p 3306:3306 -v /var/lib/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
```

### Redis使用

```dockerfile
docker run -d -p 6379:6379 -v /etc/redis/redis.conf:/etc/redis/redis.conf redis:4.0 redis-server /etc/redis/redis.conf --appendonly yes
```

* 可以在redis.conf中配置redis密码，在requirepass后面配置密码；
* 刚启动容器就会自动退出，可在redis.conf中将daemonize yes注释




















