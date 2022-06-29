#!/bin/bash

echo "Building Client"
cd ../client/
ng build --delete-output-path true
cd ../server/
echo "Starting Server"
go run main.go & 
_pid=$!
echo "Server Running..."

echo "Press any key to kill"
while [ true ] ; do
read -t 3 -n 1
if [ $? = 0 ] ; then
kill $_pid;
exit ;
fi
done