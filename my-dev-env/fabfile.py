#!env python

from fabric import task

import config

@task
def build_platform(c):
    c.run(""" echo %s""" % ('hehe'))
    c.run(
    """docker run \
            --rm -t \
            -v %s/root:/root \
            build-base \
            bash -c '
            set -ex
export PYENV_ROOT="${HOME}/.pyenv"
export RBENV_ROOT="${HOME}/.rbenv"
export PATH="${PYENV_ROOT}/bin:${RBENV_ROOT}/bin:${RBENV_ROOT}/shims:${PATH}"

# Install pyenv
curl -L https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
eval "$(pyenv init -)"
pyenv install -s 3.6.6

# Install rbenv
# This is a hack for Ubuntu 18.04
apt-get install -y libssl1.0-dev
curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash
eval "$(rbenv init -)"
rbenv install -s 2.3.1
            '
    """ % (config.local_shared_docker_data_dir()))

