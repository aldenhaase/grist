import threading
import time
import sys
import os
sys.path.insert(1,'/usr/lib/google-cloud-sdk/bin')
import dev_appserver
server = threading.Thread(target = dev_appserver.main)
server.daemon = True
server.start()
time.sleep(5)
#os.system("newman run https://api.getpostman.com/collections/THE_CLIENT_TESTS?apikey="+os.getenv("POSTMAN_API"))