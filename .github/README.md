# TimeBook Server write in Go Programming Language 

**Surpport by fiber**


加入 `fiber` 依赖
```bash
go get -u github.com/gofiber/fiber
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
