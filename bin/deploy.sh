#!/bin/bash
cd ..
gcloud builds submit --region=us-west2 --config cloudbuild.yaml .