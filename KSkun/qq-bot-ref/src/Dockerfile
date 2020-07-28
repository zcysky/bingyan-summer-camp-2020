FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build qq-bot-ref -o /env/qq-bot-ref
ENTRYPOINT ["/env/qq-bot-ref"]