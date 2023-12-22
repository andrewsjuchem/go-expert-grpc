# Go Expert - gRPC

## Running the app

Starting the server.
```
go run cmd/grpcServer/main.go
```

Running the client.
```
go run cmd/grpcClient/main.go
```

## Testing with Postman

1. Import the proto file into Postman to automatically load the definitions.

2. Call the APIs using the Postman's interface. See payload examples below. The field mask parameter is optional and is for filtering what attributes will be included in the response. If not informed, all fields will be returned.

127.0.0.1:50051 | CourseService/ListCourses
```
{
    "field_mask": {
        "paths": ["id", "name"]
    }
}
```

127.0.0.1:50051 | CategoryService/GetCategory
```
{
    "id": "c24ce97c-baa4-414a-808e-734a1e88bee8",
    "field_mask": {
        "paths": ["id", "name"]
    }
}
```


## Testing with Evans.
```
evans -r repl
service CategoryService
call CreateCategory
[enter name]
[enter description]
call ListCategories
```

## Development

When the proto file is changed, run the command below to generate or re-generate the code in the `pb` folder through ProtoC.
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
````

When it is the first time running the application, create the tables in the database.
```
sqlite3 db.sqlite
CREATE TABLE categories (id string, name string, description);
CREATE TABLE courses (id string, name string, description, string, category_id string);
```