yarn global add grpc-tools  
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:. --grpc_out=grpc_js:. hello_service.proto  