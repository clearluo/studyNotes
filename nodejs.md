## npm使用

* npm安装卡顿是因为国内网络连接npm速度慢，可以使用代理方式解决；

```bash
npm config set registry https://registry.npm.taobao.org
```

* npm全局安装以后，项目中还是找不到对应包

  * 在node安装目录中，新建两个node_modules统计目录 node_cache和node_global

  * 设置变更目录

  ```bash
  npm config set prefix "D:\nodejs\node global"
  npm config set cache "D:\nodejs\node cache"
  ```

  * 查看目录是否生效

  ```bash
  npm config get prefix
  npm condig get cache
  ```

  * 设置系统Path变量为 "D:\nodejs\node global"
  * 再加一个环境变量 NODE_PATH: "D:\nodejs\node global"

  
