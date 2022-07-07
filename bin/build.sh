#!/bin/bash
mkdir "$APP_LOCATION"go-app/
mkdir "$APP_LOCATION"build/
cp "$APP_CONFIG_LOCATION"app.yaml "$APP_LOCATION"go-app/
cp -r "$SERVER_SOURCE_LOCATION"* "$APP_LOCATION"go-app/
cp -r "$CLIENT_SOURCE_LOCATION"* "$APP_LOCATION"build/
cd "$APP_LOCATION"build/ 
npm install && ng test lystr --watch=false --code-coverage && ng build --output-path=../go-app/dist/
cd "$APP_LOCATION"go-app/
go test
find -regex "./.*_test.*" | xargs mv -t /tmp
python2 "$APP_LOCATION"bin/runAPI.py --application=lystr-354722 "$APP_LOCATION"go-app/app.yaml