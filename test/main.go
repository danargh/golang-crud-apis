package main

import (
	"context"
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	ctx := context.Background()
	fmt.Println(ctx)
}
