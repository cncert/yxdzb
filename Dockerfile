# 构建images(镜像)使用
FROM golang
WORKDIR /go/src/
COPY yxdzb .
COPY script.sh .
EXPOSE 8080
CMD ["/bin/bash", "/go/src/script.sh"]
