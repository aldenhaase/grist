FROM aldenhaase/angular-go:base
RUN mkdir /workspace
COPY . /workspace
ARG POSTMAN_API
ENV POSTMAN_API=$POSTMAN_API
ARG SERVER_SIG
ENV SERVER_SIG=$SERVER_SIG
RUN  /workspace/bin/build.sh