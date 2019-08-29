import os

from fabric import task
from invoke import run as local

def get_aws_account_id(c):
  r = local('aws sts get-caller-identity --output text --query Account')
  if r.ok:
    return r.stdout.strip()
  raise "failed to get aws account id: %s" % r.stderr

def get_ecr_repo(account_id, app, version):
  return '%s.dkr.ecr.us-west-2.amazonaws.com/%s:%s' % (account_id, app, version)

def get_ecr_domain(account_id):
  return '%s.dkr.ecr.us-west-2.amazonaws.com' % (account_id)

def get_ecr_login_token(c):
  r = local('aws ecr get-login --no-include-email')
  if r.ok:
    return r.stdout.strip()
  raise "failed to get aws account id: %s" % r.stderr


def pull_docker_repo(c, account_id, repo):
  login_cmd = get_ecr_login_token(c)
  c.sudo(login_cmd)
  c.sudo('docker pull %s' % repo)
  c.sudo('docker logout %s.dkr.ecr.us-west-2.amazonaws.com' % account_id)

def stop_docker_container_by_name(c, name):
    r = c.sudo('docker ps -qaf "name=^/%s"' % name)
    container_id = r.stdout.strip()
    if r.ok and container_id != '':
        print('container_id :%s' % container_id)
        c.sudo('docker stop %s' % container_id)
        c.sudo('docker rm %s' % container_id)

def deploy_docker_app(c, app, docker_opt, version = 'latest', app_env = 'development'):
    account_id = get_aws_account_id(c)
    repo = get_ecr_repo(account_id, app, version)

    # Update local image.
    pull_docker_repo(c, account_id, repo)

    # Stop running apiserver if any.
    stop_docker_container_by_name(c, app)

    c.sudo("""sudo docker run \
            -d \
            -v /etc/ssl/certs:/etc/ssl/certs \
            --restart unless-stopped \
            --log-opt mode=non-blocking \
            --log-opt max-buffer-size=64m \
            --name %s \
            -e APP_ENV=%s \
            %s %s
            """ % (app, app_env, docker_opt, repo))

    cleanup_dangling_docker_image(c)

@task(optional = ['app_env'])
def deploy_payment_integration(c, app_env = 'development'):
    deploy_docker_app(c,
            app = 'payment-integration',
            docker_opt = '-p 50051:50051 -p 50061:50061',
            app_env = app_env)

def cleanup_dangling_docker_image(c):
    c.sudo('bash -c \'docker images -qf "dangling=true" | xargs -I {} docker rmi {}\'')

def upload_nginx_conf(c, tmpdir, website):
    c.put('docker_images/nginx/nginx/conf.d/%s.conf' % website,
        remote = '%s/%s' % (tmpdir, website))
    c.sudo('mv %s/%s /etc/nginx/sites-available/%s' % (tmpdir, website, website))
    c.sudo('ln -sf /etc/nginx/sites-available/%s /etc/nginx/sites-enabled/%s' % (website, website))

@task
def start_global_socks_service(c, name='socks-proxy'):
    local("""bash -c 'docker ps -qf name=%s | xargs -I {} bash -c "docker stop {} && docker rm {}"'
    """ %(name))
    local("""docker run -d \
            --restart unless-stopped \
            --log-opt mode=non-blocking \
            --log-opt max-buffer-size=64m \
            -v /etc/ssl/certs:/etc/ssl/certs \
            -v $HOME/.ssh:/root/.ssh \
            -p 127.0.0.1:8888:8888 \
            --name %s \
            ssh-client \
            ssh -N -D 0.0.0.0:8888 dev.redis.biz
    """ % (name))

@task
def install_nginx(c):
    c.sudo('DEBIAN_FRONTEND=noninteractive apt-get update')
    c.sudo("""DEBIAN_FRONTEND=noninteractive apt-get install -y \
            nginx letsencrypt python3-pip""")
    c.sudo('pip3 install awscli')

    tmpdir = mktemp(c)
    upload_nginx_conf(c, tmpdir, 'default')
    # cleanup
    c.run('rm -rf %s' % tmpdir)

    c.sudo('systemctl restart nginx')

def mktemp(c):
    r = c.run('mktemp -d')
    if not r.ok:
        raise "failed to mktemp dir: %s" % r.stderr
    return r.stdout.strip()

@task
def legacy_deploy_nginx(c):
    tmpdir = mktemp(c)
    website = c.host

    # upload conf files
    upload_nginx_conf(c, tmpdir, website)
    # upload cert files
    account_id = get_aws_account_id(c)
    ciphertext_file = '%s/ciphertext_blob' % tmpdir
    key_file = '%s/key_file' % tmpdir
    encrypted_cert_file = '%s/%s.tar.gz.encrypted' % (tmpdir, website)
    cert_file = '%s/%s.tar.gz' % (tmpdir, website)

    c.put('secure_data/%s.ciphertext_blob' % account_id, remote = ciphertext_file)
    c.put('secure_data/letsencrypt/%s.tar.gz.encrypted' % website, remote = encrypted_cert_file)

    # Decrypt cert files.
    c.run("""
aws kms decrypt \
  --ciphertext-blob "fileb://%s" \
  --query 'Plaintext' \
  --output text \
  | base64 --decode \
  >"%s"
    """ % (ciphertext_file, key_file))
    c.run("""
openssl enc -d -aes256 \
  -kfile "%s" \
  -in "%s" \
  -out "%s"
    """ % (key_file, encrypted_cert_file, cert_file))
    c.sudo('tar zxf "%s" -C /etc' % cert_file)

    # cleanup
    c.run('rm -rf %s' % tmpdir)

    c.sudo('systemctl reload nginx')

@task
def deploy_nginx_in_docker(c, app_env = 'development'):
    deploy_docker_app(c,
            app = 'nginx',
            docker_opt = """-p 443:443 \
            --link apiserver \
            -v /etc/ssl/certs:/etc/ssl/certs \
            -v /root/nginx/letsencrypt:/etc/letsencrypt \
                """,
            app_env = app_env)

@task
def install_docker_node(c):
    c.sudo('DEBIAN_FRONTEND=noninteractive apt-get update')

    c.sudo("""DEBIAN_FRONTEND=noninteractive apt-get install -y \
            vim curl \
            apt-transport-https \
            ca-certificates \
            software-properties-common
            """)
    c.sudo('curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -')
    c.sudo('add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"')
    c.sudo('DEBIAN_FRONTEND=noninteractive apt-get update')
    c.sudo('DEBIAN_FRONTEND=noninteractive apt-get install -y docker-ce')
    # Auto start service
    c.sudo('update-rc.d docker enable')
    c.put('fabric-templates/docker/daemon.json', remote = '/etc/docker/daemon.json')


