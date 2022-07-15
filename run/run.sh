#!/bin/bash
cd ../src/client/
npm run build:ssr
cd ../../
docker build --no-cache --progress=plain -t tempbuild -f run/Dockerfile .
docker run --name tempcont -d -p 8080:8080 -p 8081:8081 tempbuild
echo "App running -- press any key to kill"

function shutDown(){
    echo "Shutting Down Gracefully"
    docker stop tempcont
    docker rm tempcont
    docker image rm tempbuild
}

trap shutDown EXIT

while [ true ] ; do
docker logs tempcont &> log.txt
read -s -t 3 -n 1
if [ $? = 0 ] ; then
exit ;
fi
done
