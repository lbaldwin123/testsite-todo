#!/bin/bash -e 
echo "query todo 3"
curl http://localhost:8080/todos/3

echo "add content"
curl -H "Content-Type: application/json" -d '{"name":"New Todoaaaaaaaaaaaaaaaaaaaaaaaaa"}' http://localhost:8080/todos

echo "query content"
curl http://localhost:8080/todos/3

echo "delete content"
curl -H "Accept: application/json" -X DELETE http://localhost:8080/todos/3

echo "query content"
curl http://localhost:8080/todos/3
