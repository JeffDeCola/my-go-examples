#!/bin/bash
# my-go-examples update_concourse.sh

fly -t ci set-pipeline -p my-go-examples -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
