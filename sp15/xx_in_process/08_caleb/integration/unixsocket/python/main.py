import socket
import json
import os

def main():
    print("starting server")
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.bind("/tmp/example.sock")
    sock.listen(1)
    while True:
        conn, addr = sock.accept()
        try:
            f = conn.makefile('rw')
            for ln in f:
                obj = json.loads(ln)
                print("recv %s(%s)" % (obj["method"], obj["params"][0]))
                if obj["method"] == "add":
                    print("send %s" % obj["params"][0])
                    conn.send(json.dumps({
                        "result": obj["params"][0]["X"] + obj["params"][0]["Y"],
                        "id": obj["id"],
                        "error": None,
                    }))
        except ValueError, e:
            print("ValueError", e)
        finally:
            conn.close()

try:
    main()
finally:
    os.unlink("/tmp/example.sock")
