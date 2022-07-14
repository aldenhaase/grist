#!/bin/bash
mkdir "$APP_LOCATION"go-app/
mkdir "$APP_LOCATION"node-app/
rm -r "$APP_LOCATION"build/
mkdir "$APP_LOCATION"build/
cp "$APP_CONFIG_LOCATION"api.yaml "$APP_LOCATION"go-app/
cp "$APP_CONFIG_LOCATION"web.yaml "$APP_LOCATION"node-app/
cp "$APP_CONFIG_LOCATION"dispatch.yaml "$APP_LOCATION"
cp -r "$WEB_SERVER_SOURCE_LOCATION"* "$APP_LOCATION"node-app/
cp -r "$API_SERVER_SOURCE_LOCATION"* "$APP_LOCATION"go-app/
cp -r "$CLIENT_SOURCE_LOCATION"* "$APP_LOCATION"build/
cd "$APP_LOCATION"build/ 
npm install && ng test lystr --watch=false --code-coverage && ng build --output-path=../node-app/dist/
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
npm install
export PORT="8081"
python2 "$APP_LOCATION"bin/runClient.py
if [ "$?" -ne "0" ]; then
  echo "Client tests failed"
  exit 1
fi