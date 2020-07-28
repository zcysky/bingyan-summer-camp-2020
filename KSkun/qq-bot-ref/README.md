# QQ 机器人 参考实现

## 配置

在 `/bot` 下新建一个文件 `config.txt`，写入以下内容。

```
login <QQ号> <QQ密码>
```

然后运行 `sudo docker-compose up -d` 即可。

注意：有时 Mirai 会要求输入验证码和手机扫码登录等，此时可以把 qq-bot-ref_mirai 容器的 ENTRYPOINT 改成 `tail -f /dev/null`，后在容器中自行启动 Mirai 进行登录，之后再启动容器 qq-bot-ref_app。
