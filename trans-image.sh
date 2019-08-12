#!/bin/sh

# sh trans-image.sh pull root@10.20.1.72 postgres:9.6.2 harbor.waret.net/library/postgres:9.6.2

if [[ $# == 0 ]]; then
    echo trans-image.sh push \$image \$himage
    echo trans-image.sh pull \$host \$image \$himage
    exit 0;
fi

if [[ "$1" == "push" ]]; then
    image=$2
    himage=$3
    docker pull $image
    docker tag $image $himage
    docker push $himage
    docker rmi $himage
elif [[ "$1" == "pull" ]]; then
    host=$2
    image=$3
    himage=$4
    ssh $host docker pull $himage
    ssh $host docker tag $himage $image
    ssh $host docker rmi $himage
fi
