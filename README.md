# Go Expert - gRPC

## Running the app

Start the server
```
go run cmd/grpcServer/main.go
```

Run the client
```
go run cmd/grpcClient/main.go
```

## Development

When the proto file is changed, run the command below to generate the code in the `pb` folder through ProtoC.
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
````

When it is the first time running the application, create the tables in the database.
```
sqlite3 db.sqlite
CREATE TABLE categories (id string, name string, description);
CREATE TABLE courses (id string, name string, description, string, category_id string);
```

Testing with Evans
```
evans -r repl
service CategoryService
call CreateCategory
[enter name]
[enter description]
call ListCategories
```