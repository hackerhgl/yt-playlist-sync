#!/bin/bash

hgh=3
low=1
time=$((low + RANDOM%(1+hgh-low)))

sleep $time

echo "SCRIPT fork:$1 | item: $2 [SLEEP:$time]"