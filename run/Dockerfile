FROM aldenhaase/angular-go:base
RUN  mkdir /launcher 
CMD  python2 /scripts/dirty_run.py --application=lystr-354722 /workspace/server/go-app/api.yaml --support_datastore_emulator=yes --port=8081 --host=0.0.0.0 --admin_host=0.0.0.0