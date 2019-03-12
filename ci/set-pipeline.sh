#!/bin/bash
# my-go-examples set-pipeline.sh

fly -t ci set-pipeline -p my-go-examples -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
