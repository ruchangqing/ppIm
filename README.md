**分布式社交服务：ppIm**<br>

**语言：Golang**

**特性：**<br>
1、集群服务注册发现<br>
2、分布式连接管理<br>
3、websocket协议支持各种平台<br>

**环境要求：**<br>
1、Etcd<br>
2、Mysql5.6+<br>
3、Redis<br>
4、ElasticSearch7.0+<br>

**运行：**<br>
1、部署elasticsearch、redis、mysql、etcd<br>
2、导入 sql/database.sql 到mysql<br>
3、修改配置文件config.yml<br>
4、go run main.go<br>

**分布式部署：**<br>
1、启动多个节点服务<br>
2、nginx负载均衡<br>

**pprof性能监控访问：**<br>
1、配置项：http.pprof<br>
2、web访问：http://ip:port/debug/pprof<br>

**websocket并发测试：**<br>
执行并发测试：go run ./test/ws.client.go；<br>
建议linux下进行并发测试，默认并发10000个客户端连接websocket并发送认证消息，可手动修改并发客户端数量。<br>


**API接口列表：**<br>
用户登录<br>
用户注册<br>
设置昵称<br>
设置头像<br>
实名认证<br>
上传用户位置(经纬度)<br>
附近的人<br>
搜索用户<br>
好友列表<br>
添加好友<br>
添加好友请求列表<br>
同意/拒绝好友请求<br>
删除好友<br>
**发送好友消息**<br>
**撤回好友消息**<br>
创建群组<br>
搜索群组<br>
我的群组<br>
申请加群<br>
申请加群请求列表<br>
同意/拒绝加群请求<br>
退出群组<br>
踢出群组<br>
**发送群组消息**<br>
**撤回群组消息**

<br>

**文档说明:**<br><br>
**一、Api接口：**

登录：/api/v1/login<br>
参数：username `string`<br>
参数：password `string`<br>

注册：/api/v1/register<br>
参数：username `string`<br>
参数：password `string`<br>

<br>
登录/注册成功后会返回token，其他接口和websocket授权都需要带上token；api 接口在header添加 Login-Token 字段，websocket在连接后发送一个token认证信息进行认证，15秒之内没有认证成功会强制断开连接。<br>

**二、websocket长连接：**<br>
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


