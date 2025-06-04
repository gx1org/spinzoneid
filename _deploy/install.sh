#!/bin/sh

echo "running go mod tidy"
cd ..
go mod tidy
cp .env.example .env
cd ..

echo "TODO:"
echo "- Edit .env"
echo "- Exec start.sh"
