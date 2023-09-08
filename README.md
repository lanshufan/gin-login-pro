# gin-login-pro
a simple project about vue3 + GIN login page

## 配置
```
cd server
# edit api.ini
app_mode = debug

# gin info
[server]
bind = 0.0.0.0:8888

# mysql info
[mysql]
ip = 127.0.0.1
port = 3306
...

# import init.sql
use gin;
source init.sql;
```

## 后端
```
cd server

# install dependencies
go mod tidy

# run server
go run main.go

# build
go run build
```

## 前端
```
cd web

# install dependencies
npm install

# start web
npm run serve

# build
npm run build
```

## 访问
```
http://127.0.0.1:8080
# default username and password
admin/123456
```
