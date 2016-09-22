#!/bin/bash
. ~/lib-shell/lib.sh

git_repo_clone git://git.apache.org/zookeeper.git release-3.4.8
git_repo_clone https://github.com/facebook/folly.git
git_repo_clone https://github.com/google/googletest.git
git_repo_clone https://github.com/facebook/proxygen.git
git_repo_clone https://github.com/facebook/fbthrift.git
git_repo_clone https://github.com/facebook/wangle.git
git_repo_clone https://github.com/facebook/folly.git
git_repo_clone https://github.com/facebook/rocksdb.git
git_repo_clone https://github.com/apache/thrift.git

# For Fun
git_repo_clone https://github.com/spring-projects/eclipse-integration-gradle.git
git_repo_clone https://github.com/MailCore/mailcore2.git
