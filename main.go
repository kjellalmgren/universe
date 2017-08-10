package main

//	expriemntal: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags universe -ldflags '-w'
//	exprimental: GOOS=linux GOARCH=arm64 go build -a --ldflags 'extldflags "-static"' -tags gonet -installsuffix universe .

import (
	"fmt"
)

func main() {
	fmt.Println("Hallo universe!")
}
