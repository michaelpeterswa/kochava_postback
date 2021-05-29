## Setup

```
    docker pull redis
    docker pull michaelpeterswa/kochava_postback_ingest
    docker pull michaelpeterswa/kochava_postback_delivery
```

## Container Spin Up

```
    docker run -p 6379:6379 -d --name kp_redis redis
    docker run -d --name kp_ingest -p 80:80 --restart always michaelpeterswa/kochava_postback_ingest
    docker run -d --name kp_delivery --restart always michaelpeterswa/kochava_postback_delivery
```
