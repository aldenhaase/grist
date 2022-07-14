FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY build /workspace
COPY src/server/src /workspace
#COPY src/server/web /workspace
#COPY src/server/api /workspace
COPY bin/dirty_run.sh /workspace
CMD  /workspace/dirty_run.sh