#!/bin/bash
cd ..
docker build --no-cache -t tempbuild -f localbuild/Dockerfile .
docker image rm tempbuild