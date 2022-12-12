```
docker run -it --rm --name mongodb_container \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-v mongodata:/data/db -d \
-p 27017:27017 mongo

docker exec -it mongodb_container /bin/bash

mongosh -u admin -p admin --authenticationDatabase admin

db.createUser({ user: 'user', pwd: 'password', roles: [{role: 'readWrite', db: 'microservices'}]});
```


```
protoc -I=./messages --go_out=. ./messages/*.proto

protoc -I=./messages --go-grpc_out=. ./messages/*.proto
```