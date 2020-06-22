# TimeBook Server writen in Go Programming Language 

**Powered by gin**

## TODOs

Item | Status
:-:|:-:
需求分析|100%
项目代码结构|100%
数据库设计|100%
gin服务端框架搭建|100%
auth model|100%
auth api|100%
note model|100%
note api|50%
WebClient 项目代码结构|0%
WebClient 首页|0%
WebClient 登录注册|0%
WebClient NoteList页面|0%
WebClient 个人信息页面|0%
FlutterApp  项目代码结构|0%
FlutterApp  UI规划|0%
FlutterApp  网络基础|0%
FlutterApp  数据缓存和图片处理基础|0%
FlutterApp  登录注册UI和功能|0%
FlutterApp  个人信息UI和功能|0%
FlutterApp  NodeListUI和功能|0%  
用户信息|0%
评论|0%
推送|0%
笔记本|0%






## DevLogs

加入 `gin` 依赖
```bash
go get -u github.com/gin-gonic/gin
```

加入 `viper` 依赖管理配置信息
```bash
go get -u github.com/spf13/viper
```

加入 `gorm` 依赖 管理数据库信息
```bash
go get -u github.com/jinzhu/gorm
```

加入 `mysql` 驱动 配合 `gorm` 完成数据库操作
```bash
go get -u github.com/go-sql-driver/mysql
```

加入 beego库的 `validation` 依赖，处理认证校验的问题
```bash
go get -u github.com/astaxie/beego/validation 
```

加入 jwt 依赖，处理**认证token** 生成的逻辑
```bash
go get -u github.com/dgrijalva/jwt-go
```

服务端配置信息
```yaml
# debug or release
RUN_MODE: debug

# [app]
PAGE_SIZE: 10
JWT_SECRET: 23347$040412

# [server]
HTTP_PORT: 3000
READ_TIMEOUT: 60
WRITE_TIMEOUT: 60

#[database]
TYPE: mysql
USER: xxx
PASSWORD: 123456
#127.0.0.1:3306
HOST: 127.0.0.1:3306
NAME: tbook
TABLE_PREFIX: tbook_
```

安装 docker 后 拉取最近的 mysql 镜像
```bash
docker pull mysql
```

开启一个 mysql docker 容器
```bash
docker run --name mysql \
-v $HOME/docker-mysql/conf:/etc/mysql/conf.d \
-v $HOME/docker-mysql/logs:/logs \
-v $HOME/docker-mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d \
-i \
-p 3306:3306 mysql 
```

进入 mysql 容器 登录数据库 新建一个数据库 tbook
```bash
docker start mysql
docker exec -it mysql bash
$ mysql -u root -p
123456
> show databases;
> create database tbook;
> show databases;
> quit;
$ exit
```
