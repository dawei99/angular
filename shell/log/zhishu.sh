#!/bin/bash
i=0
while [ $i -le 10 ] ; do
  isZhishu=1
  for ((j=2;j<i;j++)) ; do
    if [ $((i%j)) -gt 0 ] ; then
      isZhishu=0
    fi
  done
  if [ $isZhishu -eq 1 ] ; then
    echo $i
  fi
  ((i++))
done

