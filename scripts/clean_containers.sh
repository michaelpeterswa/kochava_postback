#!/bin/bash

echo "Cleaning up Docker containers"
docker stop kpbd && docker rm kpbd
docker stop kpbi && docker rm kpbi
docker stop kochava_redis && docker rm kochava_redis