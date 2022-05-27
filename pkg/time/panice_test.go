package time

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func hardWork(job interface{}) error {
	time.Sleep(time.Second * 5)
	return nil
}


func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error)
	go func() {
		done <- hardWork(job)
		fmt.Printf("job :%v\n", job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestAss(t *testing.T) {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), i)
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("number of goroutines2:\n", runtime.NumGoroutine())
	time.Sleep(time.Second*10)
	fmt.Println("number of goroutines:\n", runtime.NumGoroutine())

}

func TestMain2(t *testing.T) {
	var IsConsultMAP  = make(map[int64]chan bool)


	time.AfterFunc(time.Second * 1, func() {
		IsConsultMAP[44] <- false
	})

	select {
	case value := <-IsConsultMAP[44]:
		t.Logf("收到 value = %v", value)
		delete(IsConsultMAP, 44)
		t.Logf("删除 value = %v, %v", IsConsultMAP[44], IsConsultMAP[55])
	case <- time.After(time.Second * 2):
		t.Logf("timeout....")
	}

}
