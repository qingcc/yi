#!/bin/sh

source_path=trans_server.go
image_name=trans_server

echo "===> building container image"

build_result="$(go build -tags 'etcd' $source_path)"

if [[ $build_result =~ ":" ]] ; then
    echo "**** encounter building error, exit"
    echo "$build_result"
    exit
fi

docker rmi -f $image_name
docker build -t $image_name  .

echo '-> ** tagging transfer/'$image_name
docker tag $image_name transfer/$image_name
