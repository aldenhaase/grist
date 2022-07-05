FROM aldenhaase/angular-go:base
SHELL ["/bin/bash", "-c"]
RUN mkdir workspace
RUN mkdir workspace/go-app/
COPY src/client workspace/
COPY app_config/app.yaml workspace/go-app/
COPY src/server/src/ workspace/go-app/
COPY /dockerbuild/runAPI.py workspace/
RUN cd workspace && source /root/.bashrc && npm install && ng build --output-path=go-app/dist
CMD export PATH=$PATH:/usr/local/go/bin && python2 workspace/runAPI.py --application=lystr-354722 /workspace/go-app/app.yaml
#run tests