**Api接口：**

登录：/api/v1/login<br>
    参数：username `string`<br>
    参数：password `string`<br>

注册：/api/v1/register<br>
    参数：username `string`<br>
    参数：password `string`<br>

<br>
登录/注册成功后会返回token，其他接口和websocket授权都需要带上token；api 接口在header添加 Login-Token 字段，websocket在连接后发送一个token认证信息进行认证，15秒之内没有认证成功会强制断开连接.

**websocket长连接：**<br>
连接地址：ws://ip:port/ws（例如`ws://127.0.0.1:8080/ws`）<br>
登录成功后进行连接认证：<br>

用户1：
{
"MsgType": 10000,
"MsgContent": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdCI6MTYxODgyNDk3NSwiaWQiOjZ9.lmLzVsCOHY9sL0jzhL1GmLr9aTYX9jWqRvigCKbMIBE"
}

用户2：
{
"MsgType": 10000,
"MsgContent": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdCI6MTYxODgyNDk3MSwiaWQiOjV9.9gvIqIK1x08IdR07-Ops_HwueaX77YhtAgR7t4eNBlQ"
}


