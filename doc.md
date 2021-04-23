**Api接口：**

登录：/api/v1/login<br>
    参数：username `string`<br>
    参数：password `string`<br>

注册：/api/v1/register<br>
    参数：username `string`<br>
    参数：password `string`<br>

<br>
登录/注册成功后会返回token，其他接口和websocket授权都需要带上token；api 接口在header添加 Login-Token 字段，websocket在连接后发送一个token认证信息进行认证，15秒之内没有认证成功会强制断开连接。<br>

**websocket长连接：**<br>
连接地址：ws://ip:port/ws（例如`ws://127.0.0.1:8080/ws`）<br>
登录成功后进行连接认证：<br>

**通信消息JSON格式：**<br>
`{"Cmd": "指令","FromId": "来源id","ToId": "接收id","Ope": "消息通道","Type": "消息类型","Body": "消息内容"}`<br>
字段说明：<br>
Cmd：指令<br>
FromId：消息发送方id<br>
ToId：消息接收方id，ope=0时为用户id，ope=1是为群组id
Ope：消息通道，0好友消息，1群消息<br>
Type：消息类型，消息类型：0 文本消息，1 图片，2 语音，3 视频，4 地理位置信息，6 文件，10 提示消息<br>
Body：消息内容，登录认证时填用户Token<br>

**用户连接websocket后发送登录认证消息：**<br>
`{"Cmd":3,"FromId":0,"ToId":0,"Ope":0,"Type":0,"Body":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdCI6MTYxOTE0Mzg0OSwiaWQiOjV9.KQ7dOv6bE_fP5NpMehziesFMsZXDAdVrbYBHyZROw40"}
`<br>
登录成功会收到消息：<br>
`{"Cmd":4,"FromId":0,"ToId":0,"Ope":0,"Type":0,"Body":"认证成功"}
`


