import threading
import time
import sys
import os

def npmStart():
    os.system("npm start")

server = threading.Thread(target = npmStart)
server.daemon = True
server.start()
time.sleep(1)
exitCondition = os.system("newman run https://api.getpostman.com/collections/19636579-606c55f6-e459-4dfa-80fe-37aeb4c29f89?apikey="+os.getenv("POSTMAN_API"))
if (exitCondition > 0):
    sys.exit(1)
#os.system is returning 256 on failure which gets dropped to zero