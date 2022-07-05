#!/bin/bash
docker build --no-cache -t temp_build -f ../dockerbuild/Dockerfile ../
docker run --rm -it temp_build
docker image rm temp_build