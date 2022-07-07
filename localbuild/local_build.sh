#!/bin/bash
cd ..
docker build --no-cache --progress=plain -t tempbuild --build-arg POSTMAN_API=${POSTMAN_API} -f localbuild/Dockerfile .
docker image rm tempbuild
