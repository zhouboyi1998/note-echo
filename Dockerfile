FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-echo
COPY ./main /go/note-echo
COPY ./application-docker.yaml /go/note-echo

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18081
ENTRYPOINT ["./main"]
