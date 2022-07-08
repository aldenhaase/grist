FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY build /workspace
COPY src/server/src /workspace
COPY bin/dirty_run.sh /workspace
CMD  /workspace/dirty_run.sh