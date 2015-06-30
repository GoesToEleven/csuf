import json
import sysv_ipc
import time
import twitter

def main():
    api = twitter.Api(consumer_key='Qze48TI7JjUqCIQO2PeJbTCpN',
                      consumer_secret='Kj0auPQIvlqIhQ68sD69SOz52FAv7UA8A5UQbwsUQq8GhdxFbH',
                      access_token_key='3067345412-I81ggjwZaMl65C1blOCJ4KnswQ7spLVAjo5vY65',
                      access_token_secret='aZGySCdiqfoq3J2mRTNGOPldk9crJEKY58btfZEapLCqS')




    mq = sysv_ipc.MessageQueue(1234)
    while True:
        for status in api.GetStreamFilter(track="obama"):
            if "text" in status:
                if len(status["text"]) < 2048:
                    mq.send(status["text"].encode("UTF-8"))

main()
