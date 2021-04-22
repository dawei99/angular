#!/bin/bash
# ------------------------------
# | 全省 water、 huayan库 pull |
# ------------------------------
#

source ./config.sh
logFile=/home/rjh/www/kunming_water/shell/logs/yunnan_pull_water_huayan/

ips=(140 141 142 144 145 146 147 148 149 150 151 152 153 154)
#name=(kunming lijiang qujing xishuangbanna chuxiong shaotong yuxi wenshan baoshan lincang puer dali dehong mengzi)
# 是否pull water库
pullWater=

i=0
for ip in ${ips[*]}
do
  logPath=${logFile}${ip}.log # 日志文件

  # 清空日志
  echo "" > $logPath

  # 监测连通性
  ping 192.168.14.${ip} -w 3 &> /dev/null

  if [[ $? > 0 ]] ; then
    echo "无法连通" > $logPath
    continue
  fi

  pullCommand="cd /home/water/huayan && git pull"
  if test $pullWater ; then pullCommand="${pullCommand} && cd /home/water && git pull"; fi
  runCommand="cd /home/water/huayan && php ahrun.php lims"

  if [ $ip -eq 140 ] ; then
    pullCommand="${pullCommand} && cd /home/kunming_water/huayan && git pull"
    if [ $pullWater ] ; then pullCommand="${pullCommand} && cd /home/kunming_water && git pull" ; fi
    runCommand="${runCommand} && cd /home/kunming_water/huayan && php ahrun.php lims"
  fi

  echo "#########${ip}########"
  expect &> $logPath <<EOF
    spawn ssh ranjianghua@121.42.140.28
    expect "*#" {send "ssh ${nodeLoginUser}@192.168.14.${ip}\n"}
    expect  "*password: " { send "${nodePasswd}\n" }
    expect "*\$ " { send "su zhouguangli\n" }
    expect "*：" { send "${zhouguangliPasswd}\n" }

        expect "*\$ " { send "${pullCommand}\n" }
        expect "*\$ " { send "${runCommand}\n" }

    expect "*\$ " { send "exit\n" }
    expect eof
EOF

i=$((i+1))
done
