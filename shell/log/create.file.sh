#!/bin/bash
fileName="assoc"
createNum=10
intervalDay=1
intervalTime=$[$intervalDay * 86400]
current=`date "+%Y-%m-%d %H:%I:%S"`
nowtime=`date -d "${current}" +%s`
i=1
while [ $i -le $createNum ] ; do
	time=$[$nowtime - $intervalTime * $i]
	date=`date -d @${time} "+%Y-%m-%d %H:%I:%S"`
	touch -d "${date}" "${fileName}.${date}.log"
	((i++))
done

