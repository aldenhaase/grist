import test
import threading
import time
import sys
sys.path.insert(1,'/usr/lib/google-cloud-sdk/bin')
import dev_appserver

server = threading.Thread(target = dev_appserver.main)
server.daemon = True
server.start()
print("server running...")
time.sleep(5)
print("running API test...")
print("shutting down...")