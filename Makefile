gen_volvo:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/volvo.proto

gen_skoda:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/skoda.proto

all: gen_volvo gen_skoda 
	@printf "⏳ Building all the proto files...\n"
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/car.proto
	@echo "✅ Done"


clean:
	rm -r pb/*.pb.go