##
##  自动部署redis哨兵集群
##

redis systemd文件
[Unit]
Description=Redis server
After=network.target

[Service]
Type=forking
ExecStart=/opt/redis-6.2.1/src/redis-server /opt/redis-6.2.1/redis.conf
ExecStop=/opt/redis-6.2.1/src/redis-cli -a 123123 shutdown

[Install]
WantedBy=multi-user.target

redis哨兵systemd文件
[Unit]
Description=Redis sentinel server
After=network.target

[Service]
Type=forking
ExecStart=/opt/redis-6.2.1/src/redis-sentinel /opt/redis-6.2.1/sentinel.conf
ExecStop=/opt/redis-6.2.1/src/redis-cli -p 26379 shutdown

[Install]
WantedBy=multi-user.target
