#!/bin/bash
export NKEYS_PATH=/datastore/.nsc/nkeys
export XDG_DATA_HOME=/datastore/.nsc
export XDG_CONFIG_HOME=/datastore/.nsc/.config

nsc update

nsc init -d /datastore/.nsc/nats -n jetOperator
nsc edit operator --service-url nats://localhost:4222
nsc generate config --nats-resolver > /datastore/nats-res.cfg

