#!/bin/bash

echo "Starting Kochava Postback"

docker run --name kochava_redis -d redis
docker run --name kpbi -d --restart always michaelpeterswa/kpbi
docker run --name kpbi -d --restart always michaelpeterswa/kpbd