1.下载安装
wget https://archive.apache.org/dist/kafka/2.4.0/kafka_2.12-2.4.0.tgz
tar -xzvf kafka_2.12-2.4.0.tgz

2.启动zookeeper
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties

三、kafka配置

kafka的配置文件在config/server.properties文件中，主要修改参数如下，更具体的参数说明以后再整理下。
broker.id是kafka broker的编号，集群里每个broker的id需不同。我是从0开始。

listeners是监听地址，需要提供外网服务的话，要设置本地的IP地址

log.dirs是日志目录，需要设置

设置Zookeeper集群地址，我是在同一个服务器上搭建了kafka和Zookeeper，所以填的本地地址

启动kafka
bin/kafka-server-start.sh -daemon config/server.properties

停止
bin/kafka-server-stop.sh config/server.properties


kafka监控工具
https://github.com/yahoo/CMAK
https://github.com/yahoo/CMAK/archive/refs/tags/3.0.0.5.tar.gz

