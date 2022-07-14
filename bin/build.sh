#!/bin/bash
mkdir "$APP_LOCATION"go-app/
#mkdir node-app
rm -r "$APP_LOCATION"build/
mkdir "$APP_LOCATION"build/
cp "$APP_CONFIG_LOCATION"app.yaml "$APP_LOCATION"go-app/
#copy web.yaml node-app/
#copy api.yaml go-app/
cp -r "$SERVER_SOURCE_LOCATION"* "$APP_LOCATION"go-app/
#copy api_source_location to go-app/
#copy server_source to node-app/
cp -r "$CLIENT_SOURCE_LOCATION"* "$APP_LOCATION"build/
#good
cd "$APP_LOCATION"build/ 
npm install && ng test lystr --watch=false --code-coverage && ng build --output-path=../go-app/dist/
#--output path = node-app/dist/
#CD NODE-APP AND NPM INSTALL FOR EXPRESS DEPENDENCIES
cd "$APP_LOCATION"go-app/
go test
find -regex "./.*_test.*" | xargs mv -t /tmp
python2 "$APP_LOCATION"bin/runAPI.py --application=lystr-354722 "$APP_LOCATION"go-app/app.yaml
#python2 --applicationbin/runAPI.py --service "$APP_LOCATIONgo-app/api.yaml"
#python2 --applicationrunClient.py --service "$APP_LOCATIONnode-app/web.yaml"