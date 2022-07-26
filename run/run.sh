#!/bin/bash
docker build --no-cache --progress=plain -t tempbuild -f Dockerfile .

cd ..
docker run -v $PWD/src:/workspace -v $PWD/bin:/scripts --name tempcont -d -p 8081:8081 -p 8000:8000 -p 4200:4200 tempbuild
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

