#ppIm

分布式社交服务<br>

自动注册集群、发现集群、分布式websocket长连接服务<br>

语言：Golang

环境要求：<br>
Mysql5.6+<br>
Redis3.2+<br>
ElasticSearch7.0+<br>
Etcd3+<br>

运行：<br>
1、部署elasticsearch、redis、mysql、etcd<br>
2、修改配置文件config.yml<br>
3、go run main.go


接口安全：JWT<br>
接口列表：<br>
1、用户登录<br>
2、用户注册<br>
3、设置昵称<br>
4、设置头像<br>
5、实名认证<br>
6、上传用户位置(经纬度)<br>
7、附近的人<br>
8、好友列表<br>
9、添加好友<br>
10、添加好友请求<br>
11、同意/拒绝好友请求<br>
12、删除好友<br>
13、发送私聊消息<br>
14、撤回私聊消息<br>
15、创建群组<br>
16、群组列表（开发中）<br>
17、请求加入群组（开发中）<br>
18、加入群组请求处理（开发中）<br>
19、离开群组（开发中）<br>
20、设置群成员（开发中）




