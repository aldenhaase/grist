import threading
import sys
import os

sys.path.insert(1,'/usr/lib/google-cloud-sdk/bin')
import dev_appserver


def apiStart():
    server = threading.Thread(target = dev_appserver.main)
    server.daemon = True
    server.start()
apiStart()
os.system("node /workspace/main.js")