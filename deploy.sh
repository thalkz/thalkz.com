GOOS=linux GOARCH=amd64 go build server.go
scp -r ./* roland@thalkz.com:/home/roland/thalkz.com
