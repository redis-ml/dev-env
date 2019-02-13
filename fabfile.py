#!env python

import base64, os, subprocess
import boto3

from fabric import task

@task
def deploy_fawkes_docker(c, app_version = 'latest', app_name = 'fawkes'):
    cmd, endpoint = push_fawkes_docker(c, app_version = app_version)
    c.sudo(cmd)
    docker_registry = endpoint_to_docker_registry(endpoint)
    c.sudo('docker pull %s/%s:%s' % (docker_registry, app_name, app_version))
    c.sudo("bash -c 'docker images -qf dangling=true | xargs -I {} docker rmi {}'")

def endpoint_to_docker_registry(endpoint):
    protocal = 'https://'
    if endpoint.startswith(protocal):
        return endpoint[len(protocal):]
    return endpoint

@task
def push_fawkes_docker(c, app_version = 'latest'):
    if check_local_env(c):
        raise "Aborting... please check the above message."

    cmd, endpoint = aws_ecr_login(c)
    push_docker(c, app_name = 'fawkes', app_version = app_version, endpoint = endpoint)
    return [cmd, endpoint]

@task
def push_docker(c, app_name, endpoint, app_version = 'latest'):
    docker_registry = endpoint_to_docker_registry(endpoint)
    cmd = 'docker tag %s:%s %s/%s:%s' % (app_name, app_version, docker_registry, app_name, app_version)
    local_run(cmd)
    cmd = 'docker push %s/%s:%s' % (docker_registry, app_name, app_version)
    local_run(cmd)

@task
def aws_ecr_login(c):
    cmd, endpoint = get_ecr_login_cmd(c)
    local_run(cmd)
    return [cmd, endpoint]

@task
def get_ecr_login_cmd(c):
    client = boto3.client('ecr')
    r = client.get_authorization_token()
    # print(r)
    info = r['authorizationData'][0]
    cred = base64.b64decode(info['authorizationToken']).decode('utf-8').split(':', 1)
    endpoint = info['proxyEndpoint']
    return ['docker login -u %s -p %s %s' % (cred[0], cred[1], endpoint), endpoint]

@task
def start_docker_build_env(c, linked_svc=''):
    container_name = 'build-env-private'
    container_image = 'build-base'
    start_docker_dev_env(c,
            container_name = container_name,
            container_image = container_image,
            linked_svc = linked_svc,
    )

@task
def start_docker_dev_env(c,
        container_name,
        container_image,
        linked_svc = '',
        extra_params = ''):
    homedir = c.run('echo ~').stdout.strip()
    cwd = c.run('pwd').stdout.strip()

    linked_svc_param = ''
    linked_svc_list = list(filter(lambda x: len(x) > 0, linked_svc.split(',')))

    if len(linked_svc_list) > 0:
        linked_svc_param = '--link %s' % (' --link '.join(linked_svc_list))

    cleanup_docker_container(c, container_name)
    c.run("""
docker run \
  -d --restart unless-stopped \
  {linked_svc_param} \
  --name {name} \
  -v {HOME}/shared/linux/opt:/opt \
  -v {HOME}/shared/:/shared/ \
  -v {HOME}/:/host-home/ \
  -v {HOME}/github/:/root/github/ \
  -v {HOME}/secure/aws.redis/aws:/root/.aws \
  -v {HOME}/secure/private/ssh:/root/.ssh \
  -v {HOME}/secure/private/root:/root \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /root \
  {extra_params} \
  {container_image} \
  sleep 300d \
    """.format(
        linked_svc_param = linked_svc_param,
        name = container_name,
        HOME = homedir,
        extra_params = extra_params,
        container_image = container_image,
        ))

@task
def cleanup_docker_container(c, container_name):
    result = c.run("docker ps -qa --format '{{.ID}} {{.Names}}'")
    if not result.ok:
        raise Exception('failed to query docker containers')
    for line in result.stdout.strip().split("\n"):
        l = line.split(' ', 1)
        if l[1] != container_name:
            continue
        id = l[0]
        c.run('docker stop %s' % id)
        c.run('docker rm %s' % id)

@task
def check_local_env(c):
    print("checking local environment settings")
    if os.environ['AWS_DEFAULT_REGION'] is None:
        """
        This check is necessary since we some tasks relies on local CLI, where the env is usually
        different from within Python runtime.
        """
        print("  [WARNING] AWS_DEFAULT_REGION is not set.")
        return True
    return False

def get_aws_ecr_registry(account_id = None, region = None):
    if account_id is None:
        account_id = get_aws_account_id()
    if region is None:
        region = get_aws_default_region()
    return '%s.dkr.ecr.%s.amazonaws.com' % (account_id, region)

def local_run(cmd):
    print(cmd)
    resp = subprocess.run(cmd, shell = True, check = True)
    print(resp)

def get_aws_default_region():
    sess = boto3.session.Session()
    return sess.region_name

def get_aws_account_id():
    client = boto3.client('sts')
    r = client.get_caller_identity()
    return r['Account']

@task
def deploy_aws_meta(c, host):
    print("deploy aws meta service")

@task
def demo(c):
    """
    For learning purpose only.
    """
    result = c.run('hehe')

    print(c.run('pwd').stdout)
    if c.run("""abc """, warn = True).failed:
        print("failed")
    c.run("""echo abc """)
