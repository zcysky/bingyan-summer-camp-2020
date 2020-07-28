FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build qq-bot-ref
ENTRYPOINT ["./qq-bot-ref"]