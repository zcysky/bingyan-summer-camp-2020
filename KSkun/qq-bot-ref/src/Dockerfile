FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build qq-bot-ref
RUN mv /app/src/qq-bot-ref /app/env/qq-bot-ref
ENTRYPOINT ["/app/env/qq-bot-ref"]