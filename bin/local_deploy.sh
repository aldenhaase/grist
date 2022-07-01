#!/bin/bash
cd ..
cd ..
rm -rf workspace/
cd lystr
cloud-build-local --config=cloudbuild/local_cloudbuild.yaml --dryrun=false --push=false  --write-workspace=../ .

/usr/bin/python2.7 /usr/lib/google-cloud-sdk/bin/dev_appserver.py --application=lystr-354722 go-app/app.yaml