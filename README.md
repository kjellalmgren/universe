# README

    #
    $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags universe -ldflags '-w'
    $ docker build -f Dockerfile.scratch -t tetracon/universe:0.1 .
    $ docker run --name universe -P tetracon/universe:0.1
