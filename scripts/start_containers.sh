#!/bin/bash

echo "Starting Kochava Postback"

docker run --name kochava_redis -d redis
docker run --name kochava_postback --link kochava_redis:redis michaelpeterswa/kochava_postback