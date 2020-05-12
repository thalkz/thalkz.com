GOOS=linux GOARCH=amd64 go build
scp -r ./* roland@thalkz.com:/home/roland/blog
