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

[说明文档](doc.md)

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
发送好友消息<br>
撤回好友消息<br>
创建群组<br>
搜索群组<br>
我的群组<br>
申请加群<br>
申请加群请求列表<br>
同意/拒绝加群请求<br>
退出群组<br>
踢出群组<br>


