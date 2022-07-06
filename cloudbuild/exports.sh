#!/bin/bash
echo export APP_LOCATION=\"/workspace/\" >> /root/.bashrc
echo export APP_CONFIG_LOCATION=\"/workspace/app_config/\" >> /root/.bashrc
echo export SERVER_SOURCE_LOCATION=\"/workspace/src/server/src/\" >> /root/.bashrc
echo export CLIENT_SOURCE_LOCATION=\"/workspace/src/client/\" >> /root/.bashrc
echo export PATH=$PATH:/usr/local/go/bin:/root/.nvm/versions/node/v16.15.1/bin >> /root/.bashrc