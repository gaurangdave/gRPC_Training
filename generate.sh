# Generating Hello World
sudo protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld/helloworld.proto --plugin=/home/node/go/bin/protoc-gen-go-grpc --plugin=/home/node/go/bin/protoc-gen-go

# Generating Greet
sudo protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto --plugin=/home/node/go/bin/protoc-gen-go-grpc --plugin=/home/node/go/bin/protoc-gen-go

# Generating Calculator
sudo protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calcpb/calc.proto --plugin=/home/node/go/bin/protoc-gen-go-grpc --plugin=/home/node/go/bin/protoc-gen-go


# Generating Donuts
sudo protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative donuts/protos/donuts.proto --plugin=/home/node/go/bin/protoc-gen-go-grpc --plugin=/home/node/go/bin/protoc-gen-go