#!/bin/bash

echo "Cleaning up Docker containers"
docker stop kochava_postback && docker rm kochava_postback
docker stop kochava_redis && docker rm kochava_redis