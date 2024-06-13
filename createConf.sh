#!/bin/bash
set -e

echo "- api: apps/v1
  name: app
  spec: 
    - container: 
        - name: cont
        - image: app:$1
        - port: 1241" > conf.yaml