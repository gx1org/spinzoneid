#!/bin/sh

echo "running go mod tidy"
cd ..
go mod tidy
if [ ! -f .env ]; then
  cp .env.example .env
fi
cd ..

echo "TODO:"
echo "- Edit .env"
echo "- Exec start.sh"
