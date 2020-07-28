FROM golang:alpine
WORKDIR /app
COPY . /app
ENV GOPROXY https://goproxy.cn
CMD ["bash", "-c", "go build -o /env/qq-bot-ref qq-bot-ref && /env/qq-bot-ref"]