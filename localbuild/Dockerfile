FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY . /workspace
ARG POSTMAN_API
ENV POSTMAN_API=$POSTMAN_API
RUN  /workspace/cloudbuild/exports.sh && source /root/.bashrc && /workspace/bin/build.sh