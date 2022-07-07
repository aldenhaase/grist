FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY . /workspace
RUN  /workspace/cloudbuild/exports.sh && source /root/.bashrc && /workspace/bin/build.sh