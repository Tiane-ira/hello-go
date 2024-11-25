package extend

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

func DoWork(ctx context.Context, wg *sync.WaitGroup) {
loop:
	for {
		fmt.Printf("开始耗时操作...")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			break loop
		default:
		}

	}
	key := TraceCode("trace")
	trace, ok := ctx.Value(key).(string)
	if ok {
		fmt.Printf("worker done, trace:%s", trace)
	}
	wg.Done()
}
