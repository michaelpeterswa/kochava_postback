#!/bin/bash

echo "building ingest..."

docker build -t michaelpeterswa/kochava_postback_ingest ingest/

echo "building delivery..."

docker build -t michaelpeterswa/kochava_postback_delivery delivery/

echo "both builds completed successfully"