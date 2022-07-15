FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY src/client/dist/client/server /workspace
COPY src/client/dist/client/browser /workspace/dist
COPY src/server/api /workspace
COPY bin/dirty_run.py /workspace
CMD  python2 /workspace/dirty_run.py