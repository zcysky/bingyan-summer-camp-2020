FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
RUN go build qq-bot-ref
RUN mv ./qq-bot-ref /env/qq-bot-ref
ENTRYPOINT ["/env/qq-bot-ref"]