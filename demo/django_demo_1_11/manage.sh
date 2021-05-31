#!/bin/bash

unset DJANGO_SETTINGS_MODULE

./manage.py "$@"
