FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build -o /env/qq-bot-ref qq-bot-ref
ENTRYPOINT ["/env/qq-bot-ref"]