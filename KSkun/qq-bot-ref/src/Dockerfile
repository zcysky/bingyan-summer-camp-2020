FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
ENTRYPOINT ["go build -o /env/qq-bot-ref qq-bot-ref && /env/qq-bot-ref"]