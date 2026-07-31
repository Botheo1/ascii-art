#!/usr/bin/env bash
set -e

IMAGE_NAME="myapp"
TAG="latest"
CONTAINER_NAME="myapp-container"
PORT=8080

case "$1" in
  build)
    docker build -t "$IMAGE_NAME:$TAG" .
    ;;

  run)
    docker run -d --name "$CONTAINER_NAME" -p "$PORT:$PORT" "$IMAGE_NAME:$TAG"
    ;;

  stop)
    docker stop "$CONTAINER_NAME"
    docker rm "$CONTAINER_NAME"
    ;;

  clean)
    docker rmi "$IMAGE_NAME:$TAG"
    ;;

  *)
    echo "Usage: $0 {build|run|stop|clean}"
    exit 1
    ;;
esac