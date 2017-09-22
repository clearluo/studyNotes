

## 基本概念

### 镜像

1. 镜像是只读的，容器在启动的时候创建一层可写层作为最上层。

### 容器

### 仓库

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


## 使用镜像

### 常见命令

1. 从 Docker Hub 仓库下载一个 Ubuntu 12.04 操作系统的镜像 

   ```dockerfile
   docker pull ubuntu:12.04
   ```

   ​

2. 显示本地已有的镜像

   ```dockerfile
   docker images
   ```

3. 创建一个容器，让其中运行bash应用

   ```dockerfile
   docker run -t -i ubuntu:12.04 /bin/bash
   ```

4. ​

























