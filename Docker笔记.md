

## 基本概念

### 镜像

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



























