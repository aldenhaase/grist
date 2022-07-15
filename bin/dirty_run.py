import threading
import os

def npmStart():
    os.system("cd /workspace && export PORT=8081 && node main.js")

server = threading.Thread(target = npmStart)
server.daemon = True
server.start()

def apiStart():
    os.system("cd /workspace && go run main.go")

api = threading.Thread(target = apiStart)
api.daemon = True
api.start()
while True:
    next
#os.system is returning 256 on failure which gets dropped to zero