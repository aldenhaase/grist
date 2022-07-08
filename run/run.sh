#!/bin/bash
cd ../src/client/
ng build --delete-output-path --output-path=../../build/dist/
cd .. && cd ..
docker build --no-cache --progress=plain -t tempbuild -f run/Dockerfile .
docker run --name tempcont -d -p 8080:8080 tempbuild
echo "App running -- press any key to kill"
while [ true ] ; do
read -s -t 3 -n 1
if [ $? = 0 ] ; then
docker stop tempcont
docker rm tempcont
docker image rm tempbuild
exit ;
fi
done
