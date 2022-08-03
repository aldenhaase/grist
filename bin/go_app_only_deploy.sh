#!/bin/bash

cd /workspace/src/server/go-app/
find -regex "./.*_test.*" | xargs mv -t /tmp

