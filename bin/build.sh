#!/bin/bash

#mkdir "$APP_LOCATION"go-app/
#mkdir "$APP_LOCATION"node-app/
#mkdir "$APP_LOCATION"node-app/dist/
#cp /workspace/src/server/api/api.yaml "$APP_LOCATION"go-app/
#cp /workspace/src/client/web.yaml "$APP_LOCATION"node-app/
cd /workspace/src/client
npm install && ng test client --watch=false --code-coverage
#cp -r /workspace/src/client/dist/client/server/* "$APP_LOCATION"node-app/
#cp -r /workspace/src/client/dist/client/browser/* "$APP_LOCATION"node-app/dist/
#cp -r /workspace/src/client/* "$APP_LOCATION"node-app/
#cp -r /workspace/src/server/api/* "$APP_LOCATION"go-app/
#cd "$APP_LOCATION"build/ 
cd /workspace/src/server/go-app/
go test
if [ "$?" -ne "0" ]; then
  echo "Go tests failed"
  exit 1
fi
find -regex "./.*_test.*" | xargs mv -t /tmp
python2 /workspace/bin/runAPI.py --application=lystr-354722 /workspace/src/server/go-app/api.yaml --support_datastore_emulator=yes --port=8081
if [ "$?" -ne "0" ]; then
  echo "API tests failed"
  exit 1
fi
#cd "$APP_LOCATION"node-app/
#python2 "$APP_LOCATION"bin/runClient.py
cd /workspace/src/client
mkdir /node-app/
mv web.yaml /node-app/
npm run build:ssr
cp dist/client/server/* /node-app/
mv dist/client/browser /node-app/dist
mv boot_package.json /node-app/package.json
cd /node-app/
python2 /workspace/bin/runClient.py
if [ "$?" -ne "0" ]; then
  echo "Client tests failed"
  exit 1
fi

