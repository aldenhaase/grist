#!/bin/bash
docker build --no-cache -t temp_build -f ../dockerbuild/Dockerfile ../
docker image rm temp_build