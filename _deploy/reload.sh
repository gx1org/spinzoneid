#!/bin/sh

cd ..

git diff
git checkout .
git pull

echo "building"
go build -o spinzoneid
echo "running server"
pm2 reload spinzoneid --update-env
