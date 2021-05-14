#!/bin/bash
i=1
while [ $i -le 100 ];do
ret=1
for (( j=2;j<$i;j++ ));do
if [ $(($i%$j)) -eq 0 ];then
ret=0
break
fi
done
if [ $ret -eq 1 ];then
echo -n "$i "
fi
i=$(( i+1 ))
done
