##### 生成镜像
```
docker build -t go_docker:1.0 .
```
##### 直接用docker启动的方式
```
docker run -d  --name godocker \
--restart=always  \
-v /root/compose/config.yaml:/compose/config.yaml
-p 6001:6001 go_docker:1.0
```
##### 用docker compose启动
```
docker-compose up -d
```

