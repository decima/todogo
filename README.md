TODOLIST Golang
==================


Requirements
-------

In order to develop, you'll need Docker. 


Development Environment
--------

First, run, the following command:
```
docker volume create todo_cache
```


To develop, you can run the following command : 
```
docker run --rm -it -p 9010:9000  -v todo_cache:/go/pkg -v $PWD:/app -w /app    golang:1.19-alpine go run .
```

And check on port 9010. 


Production Environment
----------

You can use docker-compose.yml to build the app. 
