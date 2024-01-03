const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

// Load the protobuf definition
const packageDefinition = protoLoader.loadSync('../proto/course_category.proto', {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

// Load the gRPC package definition
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

// Create a new gRPC client
const client = new protoDescriptor.pb.CategoryService('127.0.0.1:50051', grpc.credentials.createInsecure());

// Call the ListCourse endpoint
client.ListCategories({}, (error, response) => {
  if (!error) {
    console.log('ListCategories response:', response);
  } else {
    console.error('Error calling ListCategories:', error);
  }
});