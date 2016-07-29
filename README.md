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

	conf := SetConfig("/etc/test.conf")
	host := conf.GetValue("DEFAULT", "host")
	fmt.Printf("host = %s\n", host)
