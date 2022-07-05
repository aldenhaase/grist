#!/bin/bash

docker build -t aldenhaase/jic/temp_build:latest -f ../dockerbuild/Dockerfile ../
docker image rm aldenhaase/jic/temp_build:latest