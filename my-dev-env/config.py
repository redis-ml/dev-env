from os.path import expanduser

def local_shared_docker_data_dir():
    return "%s/shared" % (local_user_dir())

def local_user_dir():
    return expanduser("~")

