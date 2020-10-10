# QQ 机器人 参考实现

## 配置

修改 `docker-compose.yml` 中的环境变量 `QQ_NUMBER` 和 `QQ_PASSWORD`，然后运行 `sudo docker-compose up -d` 即可。

注意：有时 Mirai 会要求输入验证码和手机扫码登录等，此时可以把 qq-bot-ref_mirai 容器的 ENTRYPOINT 改成 `tail -f /dev/null`，后在容器中自行启动 Mirai 进行登录，之后再启动容器 qq-bot-ref_app。

## 实现存在的问题

`/bot/Dockerfile` 里使用 MiraiOK 项目在容器内自动配置环境是不推荐的做法，更好的做法是把已经配置好的环境打包成镜像，再直接使用该镜像。
