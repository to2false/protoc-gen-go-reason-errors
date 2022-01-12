gen:
	protoc --go_out=. --go_opt=paths=source_relative --proto_path=. ./errors/errors.proto
