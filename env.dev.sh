#!/usr/bin/env bash

# If you have an SSM parameter you can put the name here
PREFIX=${APPNAME}-${ENV}

export SLACK_WEBHOOK_URL=${PREFIX}-slack-webhook-url
