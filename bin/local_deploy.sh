#!/bin/bash
cd ..
cloud-build-local --config=cloudbuild.yaml --dryrun=false --push=false  --write-workspace=../ .
