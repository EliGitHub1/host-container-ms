# host-container-ms

Author

name :   Eli Aviv
email:   eliaviv2@gmail.com

version:  

1.0

Description

Restful service in go

How to run

  1. local machine:
    1.1 git clone repository
    1.2 install go 
    1.3 db sqllite file should be in cd...
    1.4 install sqlite?
    
    1.3 buid? mod init? start?
  3. docker:

API endpoints

      /host           GET – Get a list of all hosts, returned as JSON.

      /host/:id       GET – Get host by its ID, returning the host data as JSON.

      /container      GET – Get a list of all containers, returned as JSON.
                      
      queryParam      GET - Get all containers for specific host.
      
      /container/:id  GET – Get container by its ID, returning the container data as JSON
      
      /container      POST - Create a new container in the database.
      

Test
  1. build...
  2. run...

tests:
  1. check db is up
  2. check service is connected to db
  3. test for every endpoint 

