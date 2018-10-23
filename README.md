# GO Swagger

Simple program that creates a Swagger server, access the endpoint containg the JSON, parses data and retrieves only necessary information, and prints to a provided endpoint 

To start server run (on Linux)
./data-parser-server --port=8082

To start on Windows:
.\data-parser-server --port=8082

If there are errors regarding Dependencies try deleting Gopkg.lock and Gopkg.toml and run:
dep init
