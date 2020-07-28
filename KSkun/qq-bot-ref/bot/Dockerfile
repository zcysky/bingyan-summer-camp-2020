FROM ubuntu:focal
WORKDIR /app
RUN apt update
RUN apt install -y wget
RUN wget -P /app http://t.imlxy.net:64724/mirai/MiraiOK/miraiOK_linux_amd64
RUN wget -P /app/plugins https://github.com/project-mirai/mirai-api-http/releases/download/v1.7.3/mirai-api-http-v1.7.3.jar
COPY config.txt /app/
COPY setting.yml /app/plugins/MiraiAPIHTTP/
RUN chmod -R 755 /app
ENTRYPOINT ["./miraiOK_linux_amd64"]