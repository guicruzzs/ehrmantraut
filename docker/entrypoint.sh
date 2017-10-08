#!/bin/bash

service mosquitto start
# TODO: Remove this sleep :(
sleep 1
mosquitto_sub -t test | /ehrmantraut/ehrmantraut
