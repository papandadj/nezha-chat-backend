# nezha-chat-backend
聊天系统后端 对应的前端代码 https://github.com/papandadj/nezha-chat-frontend

效果展示 http://182.61.3.243:3000/demo.mp4

![](./api/doc/media/backendFunc.png)

如图， 后端主要有这么多的服务， 其中以web结尾的对前端提供http接口， web和srv之间通过grpc进行通讯， 服务注册采用的是etcd。

http服务采用的gin框架， grpc采用的是go-micro。 

项目文件夹分为下面几个部分

![](./api/doc/media/alldirs.png)

其中， 每一个文件夹代表一个服务(除了deploy， 跟common， script)， 它们是可以单独开启的， 通过etcd注册自己。 deploy文件夹里面放的是docker 部署的配置。 common里面放的是通用的一些结构体或者方法。 pkg里面是自己写的函数或者封装的库。

#### 怎样本地部署

因为mysql没有封装成镜像， 你需要首先创建配置好mysql， 并且将deploy里面的文件夹的每一个config.toml中， 将mysql改成自己的账号密码。 还要确保数据库都建好了， 建库语句在script里面的script.sql里面， 初始化数据在data.sql里面。 


##### 下载代码

`git clone git@github.com:papandadj/nezha-chat-backend.git`

##### 编译二进制

`cd deploy/ && make buildBinary`

##### 编译docker镜像

`make buildImage`

##### 开启服务

`docker-compose up`

后端运行 `docker-compose up -d` 

然后等一下， 因为`chat-srv`要确保`rabbitmq`启动后才会启动， 而docker-compose好像没有探针， 所以我只能写了出错时重启。

##### 查看是否启动

`docker container ls `看下需要的服务是否都启动了， 之后， 可以登录 `http://localhost:16686` 查看jaeger服务是否启动。

这里面主要是mysql， 其他都是docker部署， 一般不会出问题。

后端开启后会开启5个端口
9500： user-web
9501： friend-web
9502： common-web
9503： chat-web
15674: rabbitmq对外暴露的socket接口