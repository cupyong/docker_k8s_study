docker build -t go_docker:1.0 .
docker run -d  --name godocker \
--restart=always  \
-v /root/compose/config.yaml:/compose/config.yaml
-p 6001:6001 go_docker:1.0