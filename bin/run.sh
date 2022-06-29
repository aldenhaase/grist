#!/bin/bash

echo "Building Client"
cd ../client/
ng build --delete-output-path true
cd ../server/
echo "Starting Server"
go run main.go & 
_pid_run=$!
sleep 1
_pid_server=$(ps aux|grep '[g]o-build'|awk '{print $2}'|tr -d '\n')
echo "Server Running..."

echo "Press any key to kill"
while [ true ] ; do
read -s -t 3 -n 1
if [ $? = 0 ] ; then
kill $_pid_run ;
kill $_pid_server ;
exit ;
fi
done