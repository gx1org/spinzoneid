#!/bin/sh

cd ..

echo "building"
go build -o spinzoneid

echo "running server"
pm2 start ./spinzoneid --name "spinzoneid"
pm2 save
