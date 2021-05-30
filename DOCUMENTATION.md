# Useful Commands

## Delivery Agent
Run for development
```
go run .
```
Build a binary
```
go build .
```

## Build Containers (if necessary)

```
docker build -t michaelpeterswa/kochava_postback_ingest ingest/
docker build -t michaelpeterswa/kochava_postback_delivery delivery/
```

## Pull Containers (from DockerHub)

```
docker pull redis
docker pull michaelpeterswa/kochava_postback_ingest
docker pull michaelpeterswa/kochava_postback_delivery
```

## Start Containers

```
docker run -p 6379:6379 -d --name kp_redis redis
docker run -d --name kp_ingest -p 80:80 --restart always michaelpeterswa/kochava_postback_ingest
docker run -d --name kp_delivery --restart always michaelpeterswa/kochava_postback_delivery
```

## Stop Containers & Remove Containers

```
docker stop kp_delivery && docker rm kp_delivery
docker stop kp_ingest && docker rm kp_ingest
docker stop kp_redis && docker rm kp_redis
```