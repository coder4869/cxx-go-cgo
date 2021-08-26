## Prepare
1. 	Install GO, Make GOROOT and GOPATH setting.

	
## Build & Run project
1. 编译工程源码 `$ sh build.sh`

2. 运行 bin 文件 `$ ./src/go_web`
	
	
## Deploy Source Code AS Daemons (recommended)
Here provides method using tool "supervisor" and unit file ".service"

1. Install Supervisor ( http://supervisord.org/ )

```s
# Example for centos:
$ sudo yum install python-setuptools
$ sudo easy_install supervisor
```

2. Review and Create config file for Supervisor

```s
$ sudo echo_supervisord_conf > /etc/supervisord.conf
```
	
3. Make deployment for GoWeb (using compiled binary file WebMain)

```s
# 1. Edit supervisord.conf: 
	$ vi /etc/supervisord.conf

# 2. Append setting at file end:
	command=/opt/GoWeb/src/WebMain  ; absolute path of binary file WebMain
	autostart=true ; autostart with the start of supervisor
	autorestart=true ; autorestart while process down
	startsecs=10
	stdout_logfile=/opt/GoWeb/src/log/WebMain.log  
	stdout_logfile_maxbytes=1MB
	stdout_logfile_backups=10
	stdout_capture_maxbytes=1MB
	stderr_logfile=/opt/GoWeb/src/log/WebMain.log   
	stderr_logfile_maxbytes=1MB
	stderr_logfile_backups=10
	stderr_capture_maxbytes=1MB

# 3. Start supervisor
	# manual start: 
	$ sudo /usr/bin/supervisord -c /etc/supervisord.conf
	# config modified and restart: 
	$ cat /tmp/supervisord.pid | xargs sudo kill -HUP

# 4. supervisor log file: 
	/tmp/supervisord.log
```

5. Make .service deployment for supervisor (autostart with computer's start)

```s
# 1. Create .service: 
$ touch /usr/lib/systemd/system/supervisord.service
# 2. Edit .service: 	
$ vi /usr/lib/systemd/system/supervisord.service

# 3. Setting content as following: 
	# supervisord service for systemd (CentOS 7.0+)
	# by ET-CS (https://github.com/ET-CS)
	
	[Unit]
	Description=Process Monitoring and Control Daemon
	After=rc-local.service
	
	[Service]
	Type=forking
	ExecStart=/usr/bin/supervisord -c /etc/supervisord.conf
	
	[Install]
	WantedBy=multi-user.target

# 4. start .service: 
	$ systemctl start supervisord.service

# 5. view .service status: 
	$ systemctl status supervisord.service
```

# Modify Points
 WEB API Port: in "WebMain.go" file

 DataBase Connection: in “src/conf/DBConf.json” file

 DataBase Connection file Reference: if cannot find file “conf/DBConf.json”, 

	please modify reference path as absolute path
	
	(in "func LoadDBJsonConf()" of “src/conf/conf.go”)