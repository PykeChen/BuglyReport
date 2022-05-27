package robot

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRobot(t *testing.T) {
	data := make(chan bool, 1)

	time.AfterFunc(time.Second*1, func() {
		data <- true
	})

	var value bool
	select {
	case value = <- data:
		t.Logf("receive value %v", value)
	case <- time.After(time.Millisecond * 1500):
		t.Logf("[响应csIsInConsulting消息] user_id:%d", 2)
		// 是否要清理异常状态
	}

}

type Pool chan *Object

type Object struct {
	token string
}

func (object Object)Do()  {
	fmt.Printf("value=%d", 23)
}


func New(total int) Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return p
}

func TestSeap2(t *testing.T) {
	p := New(2)

	select {
	case obj := <-p:
		obj.Do( /*...*/ )

	default:
		// No more objects left — retry later or fail
		return
	}

	select {
	case obj := <-p:
		obj.Do( /*...*/ )

		p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}



	select {
	case obj := <-p:
		obj.Do( /*...*/ )

	default:
		// No more objects left — retry later or fail
		return
	}
}
