import threading
import sys
import os
sys.path.insert(1,'/usr/lib/google-cloud-sdk/bin')
import dev_appserver
server = threading.Thread(target = dev_appserver.main)
server.daemon = True
server.start()
try:
    os.system("curl -4 --connect-timeout 5 --retry 10 --retry-delay 8 --retry-connrefused http://localhost:8081")
except:
    sys.exit(1)
exitCondition = os.system("newman run https://api.getpostman.com/collections/19636579-214fe140-4cf8-47d5-b914-62456e316c38?apikey="+os.getenv("POSTMAN_API"))
if (exitCondition > 0):
    sys.exit(1)
#os.system is returning 256 on failure which gets dropped to zero