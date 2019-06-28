# **golangSocketPhp**

这是一个简单的基于socket的工具，支持多语言通过socket实现互相通信。实例是php、golang语言，序列化采用的json。

Go socket 采用c/s架构

客户端：net.Dial() Write() Read() Close()

服务器：net.Listen() Accept() Read() Write() Close()

## 测试

1、下载[源代码](https://github.com/guyan0319/golangSocketPhp)至GOPATH目录golangSocketPhp

2、运行服务端，在example目录下server.go

```
go run server.php
```

输出：

Waiting for clients

3、新窗口下运行客户端，在example目录下client.go

```
go run client.go
```

输出：

receive data string[6]:golang

golang这个是从服务端返回的数据。



4、运行php语言客户端，在php目录下的socket_client.php

```
php -f socket_client.php
```

或浏览器访问 http://localhost/xxx/socket_client.php 配置自己的网址

输出结果：

client write success
server return message is:
php



小结：

选json序列化，主要考虑它实现起来简单，很多语言支持。缺点是序列化效率低，序列化后数据相对比较大（这里跟protobuf对比）。