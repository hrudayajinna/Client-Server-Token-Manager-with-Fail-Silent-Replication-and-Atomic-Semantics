@echo off
echo Starting gRPC server...

:: Run the Go server on port 50051
go run ./server/server_grpc.go -port 50051

echo Server started. Press any key to exit...
pause > nul
