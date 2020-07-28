FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
ENTRYPOINT ["ash", "-c", "go build -o /env/qq-bot-ref qq-bot-ref && cd /env && ./qq-bot-ref"]