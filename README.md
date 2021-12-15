# host-container-ms

**Author**

name :   Eli Aviv
email:   eliaviv2@gmail.com

**version** 

1.0

**Description**

Restful service in go

**How to run**

1. git clone repository
2. install go 
3. install gcc compiler
4. change relative path for db file on your machine
5. go mod tidy
6. run

**API endpoints**

      /host           GET – Get a list of all hosts, returned as JSON.

      /host/:id       GET – Get host by its ID, returning the host data as JSON.

      /container      GET – Get a list of all containers, returned as JSON.
                      
      queryParam      GET - Get all containers for specific host.
      
      /container/:id  GET – Get container by its ID, returning the container data as JSON
      
      /container      POST - Create a new container in the database.

**Api examples**

GET http://localhost:8080/host/
GET http://localhost:8080/host/3
GET http://localhost:8080/container/
GET http://localhost:8080/container?host_id=2
GET http://localhost:8080/container/5
POST http://localhost:8080/container/
body:
{
    "host_id": 1,
    "image_name": 112003,
    "name":"name1001232"
}

**Test**
cd tests
run go test (test db connection)
