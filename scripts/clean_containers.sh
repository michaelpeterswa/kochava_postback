#!/bin/bash

echo "Cleaning up Docker containers"
docker stop kp_delivery && docker rm kp_delivery
docker stop kp_ingest && docker rm kp_ingest
docker stop kp_redis && docker rm kp_redis