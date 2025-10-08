module user-service

go 1.21

go mod init user-service     # create go.mod (if not created yet)
go mod tidy                  # fetch dependencies if any
go build -o main .           # build binary
