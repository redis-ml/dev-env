#!env python

import boto3
import json
import time
import os.path
import traceback

from queue import Empty
from multiprocessing import Queue

import otp_cred

print('Loading function')
clients = {}
enabled_services = ['sqs']

for svc in enabled_services:
    print('Initializing boto resource for %s' % svc)
    clients[svc] = boto3.client(svc)

def main_loop(q, queue_name):
    sqs = clients['sqs']
    queue = boto3.resource('sqs').get_queue_by_name(QueueName=queue_name)

    while True:
        try:
            q.get_nowait()
            print("exiting for aws otp receiver worker")
            return
        except Empty:
            pass

        try:
            one_cycle(sqs, queue.url)
        except (KeyboardInterrupt, SystemExit):
            raise
        except:
            print("failed for one cycle")
            traceback.print_exc()
            time.sleep(20)

def one_cycle(sqs, queue_url):
        print("working loop %s" % (time.strftime('%Y-%m-%d %H:%M:%S', time.localtime())))
        resp = sqs.receive_message(
            QueueUrl=queue_url,
            MaxNumberOfMessages=1,
            WaitTimeSeconds=20,
        )
        # print(resp)
        """
        {
            'Messages': [
                {
                    'MessageId': 'c77e43d8-3a38-4429-b78d-190b62b3ff95',
                    'ReceiptHandle': 'AQEBeY0NUt84GL/AzeQa2VTWzCdOWn4tVr6WYq32qJlqBKnkIfULpQBy8n1pjEVZqSlzCHLPiL3ukpO7EwwjXWPZ5vX9/rp5Qs+Mp+ejHfsRlFg2JGWV96TlELEiFzcnwAuDML99lIY25gt5UKYv+XxA6HjJ6Pl0vEe0HuXv9GU1qapfHagdcG8vTfscFihbe03d9zuNSz/PGcm7vVHXZ8HCt6AJqyBhGUZVv0Mornbtj5X0YByD1jdacoV/KRRoFhZ0nahiIIdpq1I26GmlnOiLMnDJM0/Ml+TCBcr7esLVB0zsyLT+CMS7KwAeLxmHsBWgzRSQ8p1oUgVPTM4T9sEvvWOeD/fSS4HEyKvzXv8xYzFTtcf83BDU67t9OrnG5OBx',
                    'MD5OfBody': '5875b87592b057181b53affed32e753c',
                    'Body': '{"meta": {"FirstEnqueueAt": 1550043781, "Version": 0}, "payload": {"otp": "760603", "type": "otp", "ts_mills": 1550043781071, "version": 1}}'
                }
            ],
            'ResponseMetadata': {
                'RequestId': '4464fe94-78ab-5fbc-bb4c-aa49b416dca2',
                'HTTPStatusCode': 200,
                'HTTPHeaders': {
                    'x-amzn-requestid': '4464fe94-78ab-5fbc-bb4c-aa49b416dca2',
                    'date': 'Wed, 13 Feb 2019 07:43:01 GMT',
                    'content-type': 'text/xml',
                    'content-length': '1067'
                },
                'RetryAttempts': 0
            }
        }
        """

        if 'Messages' in resp:
            for message in resp['Messages']:
                ret = process_message(message, resp)
                if ret:
                    # message.delete()
                    sqs.delete_message(
                        QueueUrl=queue_url,
                        ReceiptHandle=message['ReceiptHandle'],
                    )

def process_message(message, resp):
    """
    {
        'meta': {
            'FirstEnqueueAt': 1550043781,
            'Version': 0
        },
        'payload': {
            'otp': '760603',
            'type': 'otp',
            'ts_mills': 1550043781071,
            'version': 1
        }
    }
    """
    try:
        body = json.loads(message['Body'])
        payload = body['payload']
        if payload['type'] == 'otp':
            created_ts = payload['ts_mills'] / 1000
            now_ts = int(time.time())
            if now_ts > created_ts + 50:
                print("skipping expired otp: created at: %d : %s" % (created_ts,
                    time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(created_ts))))
                return True
            
            # print(body)
            host_home = '/host-home'
            ctx = {
                'request_id': str(payload['ts_mills']) + time.strftime(' %Y-%m-%d %H:%M:%S', time.gmtime(created_ts)),
            }
        print('unknown payload type %s: %s' % (payload['type'], payload))
        return True
    except (KeyboardInterrupt, SystemExit):
        raise
    except:
        traceback.print_exc()
        print("failed to process message: %s" % (message['MessageId']))
        return False

if __name__ == '__main__':
    q = Queue()
    main_loop(q, 'delivery')
