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
exitCondition = os.system("newman run https://api.getpostman.com/collections/19636579-214fe140-4cf8-47d5-b914-62456e316c38?apikey="+os.getenv("POSTMAN_API"))
if (exitCondition > 0):
    sys.exit(1)
#os.system is returning 256 on failure which gets dropped to zero