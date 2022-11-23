import threading
import sys
import os

sys.path.insert(1,'/usr/lib/google-cloud-sdk/bin')
import dev_appserver

os.system("cp /workspace/boot/launcher.json /launcher/package.json")

os.system("chown root.root -R /launcher")
def apiStart():
    server = threading.Thread(target = dev_appserver.main)
    server.daemon = True
    server.start()
apiStart()

os.system("cd /workspace/client && npm install && cd /launcher && npm run dev")