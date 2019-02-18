#!env python

import argparse
import traceback
import os.path
import queue
import time

from multiprocessing import Process, Queue

from receive import main_loop as otp_receiver_main_loop
import customize_aws_session
import alert_util

def private_server():
  l = []
  q = Queue()

  p = Process(target=otp_worker, args=(q,))
  p.start()
  l.append(p)

  p = Process(target=check_aws_cred_worker, args=(q,))
  p.start()
  l.append(p)

  while True:
    try:
      time.sleep(20)
    except:
      traceback.print_exc()
      break
  
  print('main proc exiting.')
  for _ in l:
    q.put(None)
  q.close()

  for p in l:
    p.join()

def check_aws_cred_worker(q):
  while True:
    check_aws_cred_wrapper()

    try:
      q.get(True, timeout=30)
      print("exiting for aws cred checking worker")
      return
    except queue.Empty:
      pass

def otp_worker(q):
  otp_receiver_main_loop(q, 'delivery')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("task",
            help="task to execute, values [main, help]",
            default="help")
    parser.add_argument("--debug", help="debugging log switch", type=bool, default=False)
    args = parser.parse_args()
    print(args)

    if args.task == 'main':
      private_server()
    else:
        print('unknown task: %s' % args.task)

    exit(0)
