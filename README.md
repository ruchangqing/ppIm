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
2、导入 sql/database.sql 到mysql
3、修改配置文件config.yml<br>
4、go run main.go<br>

**分布式部署：**<br>
1、启动多个节点服务<br>
2、nginx负载均衡

 [说明文档](doc.md)

**API接口列表：**<br>
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
16、群组列表<br>
17、请求加入群组<br>
18、加入群组请求处理<br>
19、离开群组<br>
20、设置群成员<br>
21、发送私聊消息<br>
22、撤回私聊消息<br>


