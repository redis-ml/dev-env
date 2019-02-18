#!env python

import boto3
import json

def send_alert(msg, topic):
  sns = boto3.client('sns')
  message = {"message": msg}

  print('to send alert: %s' % msg)
  arn = get_all_topics(sns)[topic]['TopicArn']
  response = sns.publish(
    TargetArn=arn,
    Message=json.dumps({
      'default': json.dumps(message),
      'email': msg,
    }),
    MessageStructure='json',
  )
  print(response)

def get_all_topics(sns):
  topics = sns.list_topics() 
  topic_list = topics['Topics']
  ret = {}
  for t in topic_list:
    topic_name = t['TopicArn'].split(':')[-1]
    ret[topic_name] = t
  return ret
  