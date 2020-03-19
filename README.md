# 项目用作后端接口

## 1. 开发环境

### 1.1 准备环境

```
sudo docker-compose up -d
```

### 1.2 创建数据库表

```
go run main.go orm syncdb -db="default" -force=false -v=true

Usage of orm command: syncdb:
  -db="default": DataBase alias name
  -force=true: drop tables before create
  -v=true: verbose info
```

### 1.3 启动项目

```
bee run
```

## 2. 构建镜像

```
sudo docker build -f DockerFile -t beego-web
```
