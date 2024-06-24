@echo off
echo Starting gRPC client commands...

:: First set of commands
echo Running commands for ID 1231...
go run ./client/client_grpc.go -create id 1231
go run ./client/client_grpc.go -write -id 1231 -name abc -low 0 -mid 10 -high 100 localhost 50051
go run ./client/client_grpc.go -read -id 1231 localhost 50051
go run ./client/client_grpc.go -drop 1231 localhost 50051

:: Second set of commands
echo Running commands for ID 1232 and 1233...
go run ./client/client_grpc.go -create id 1232
go run ./client/client_grpc.go -create id 1233
go run ./client/client_grpc.go -write -id 1232 -name def -low 0 -mid 10 -high 100 localhost 50051
go run ./client/client_grpc.go -write -id 1233 -name ghi -low 0 -mid 10 -high 100 localhost 50051
go run ./client/client_grpc.go -read -id 1232 localhost 50051
go run ./client/client_grpc.go -read -id 1233 localhost 50051
go run ./client/client_grpc.go -drop 1232 localhost 50051
go run ./client/client_grpc.go -drop 1233 localhost 50051

:: Third set of commands
echo Running commands for ID 1234 and 1235...
go run ./client/client_grpc.go -create id 1234
go run ./client/client_grpc.go -create id 1235
go run ./client/client_grpc.go -write -id 1234 -name jkl -low 0 -mid 10 -high 100 localhost 50051
go run ./client/client_grpc.go -write -id 1235 -name mno -low 0 -mid 10 -high 100 localhost 50051
go run ./client/client_grpc.go -read -id 1234 localhost 50051
go run ./client/client_grpc.go -read -id 1235 localhost 50051
go run ./client/client_grpc.go -drop 1234 localhost 50051
go run ./client/client_grpc.go -drop 1235 localhost 50051

echo All commands executed.
pause
