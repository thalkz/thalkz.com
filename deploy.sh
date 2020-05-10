GOOS=linux GOARCH=amd64 go build server.go
scp ./server roland@thalkz.com:/srv/server
scp ./public/* roland@thalkz.com:/srv/public
scp ./static/* roland@thalkz.com:/srv/static
