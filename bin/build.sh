#!/bin/bash
mkdir "$APP_LOCATION"go-app/
mkdir "$APP_LOCATION"node-app/
mkdir "$APP_LOCATION"node-app/dist/
cp "$APP_CONFIG_LOCATION"api.yaml "$APP_LOCATION"go-app/
cp "$APP_CONFIG_LOCATION"web.yaml "$APP_LOCATION"node-app/
cp "$APP_CONFIG_LOCATION"boot_client/package.json "$APP_LOCATION"node-app/
cd /workspace/src/client
npm install && ng test client --watch=false --code-coverage && npm run build:ssr
cp -r /workspace/src/client/dist/client/server/* "$APP_LOCATION"node-app/
cp -r /workspace/src/client/dist/client/browser/* "$APP_LOCATION"node-app/dist/
cp -r /workspace/src/server/api/* "$APP_LOCATION"go-app/
#cd "$APP_LOCATION"build/ 
cd "$APP_LOCATION"go-app/
go test
if [ "$?" -ne "0" ]; then
  echo "Go tests failed"
  exit 1
fi
find -regex "./.*_test.*" | xargs mv -t /tmp
python2 "$APP_LOCATION"bin/runAPI.py --application=lystr-354722 "$APP_LOCATION"go-app/api.yaml
if [ "$?" -ne "0" ]; then
  echo "API tests failed"
  exit 1
fi
cd "$APP_LOCATION"node-app/
export PORT="8081"
python2 "$APP_LOCATION"bin/runClient.py
if [ "$?" -ne "0" ]; then
  echo "Client tests failed"
  exit 1
fi