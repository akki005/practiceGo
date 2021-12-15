package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	rootCtx := context.Background()
	ctx, cancelFunc := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancelFunc()

	// taskDoneCh := make(chan int)

	//task will exit as we
	go task(ctx)

	<-ctx.Done()
}

func task(ctx context.Context) {
	time.Sleep(1 * time.Second)
	fmt.Println("finished task")
}
