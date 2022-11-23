#!/bin/bash

cd /workspace/src/client
npm install
cd /workspace/src/server/go-app/
find -regex "./.*_test.*" | xargs mv -t /tmp
go run /workspace/bin/construct_yaml.go api.yaml
cd /workspace/src/client
mkdir /node-app/
mv web.yaml /node-app/
npm run build
cp dist/client/server/* /node-app/
mv dist/client/browser /node-app/dist
mv boot_package.json /node-app/package.json
cd /node-app/

