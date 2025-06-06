#!/bin/sh

export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/root/.local/share/pnpm

cd ..

git diff
git checkout .
git pull

echo "building"
go build -o spinzoneid
echo "running server"
pm2 reload spinzoneid --update-env
