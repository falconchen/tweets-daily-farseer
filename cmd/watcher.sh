#!/bin/bash
sCount=`ps aux | grep "farseer_linux_amd64" |wc -l`

echo $sCount
if [ $sCount -lt '2' ];then
	cd /web/farseer && /web/farseer/cmd/farseer_linux_amd64 > /web/farseer/cmd/log.txt
fi
