FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY src/client/dist/client/server /workspace
COPY src/client/dist/client/browser /workspace/dist
COPY src/server/api /workspace
COPY app_config/api.yaml /workspace
COPY bin/dirty_run.py /workspace
CMD  python2 /workspace/dirty_run.py --application=lystr-354722 /workspace/api.yaml --support_datastore_emulator=yes --port=8081 --host=0.0.0.0 --admin_host=0.0.0.0