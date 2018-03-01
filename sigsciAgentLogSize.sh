#!/bin/sh

file=/var/log/sigsci.log
maxSize=180385
agentLogSize=$(wc -c < $file)
echo "logsize=$agentLogSize"
echo "maxSize=$maxSize"

while [ ! -e /tmp/stop-sigsci ]
do

	agentLogSize=$(wc -c < $file)
	if [ $agentLogSize -ge $maxSize ]; then
		echo "Clearing $file"
		echo "" > $file
	fi
	sleep 5

done
