package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var t = time.Now()

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//defer cancel()
	sleepAndTalk(ctx, 5*time.Second, "Hello")
	cancel()
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	fmt.Println("In Sleep and Talk", time.Since(t))
	select {
	case <-time.After(d):
		fmt.Println(msg, time.Since(t))
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
