goini	
========

##描述

使用goini简单读取配置文件。

##安装方法

	go get github.com/xuyingquan/goini

##使用方法

>配置文件格式样列

	[DEFAULT]
	host = 172.30.40.10
	port = 8080
	
	[mysql]
	username = root
	password = 123456
	
	[rabbitmq]
	username = guest
	password = password


>使用案例

	conf := goini.NewConfig("/etc/test.conf")
	host, _ := conf.GetString("DEFAULT", "host")
	port, _ := conf.GetInt("DEFALUT", "port")
	fmt.Printf("host = %s, port=%d\n", host, port)
