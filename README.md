# goapi
GO api

# command Docker 

docker build -t api:0.0.2 .

docker run --rm --name api api:0.0.2

docker run --rm -d -p 6379:6379  --name redis redis:latest