package facewriter

import (
	"time"
	"sync"
	"testing"
	"fmt"
	"sync/atomic"
)

/*
场景：在一个高并发的web服务器中，要限制IP的频繁访问。
现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。
完成该过程，要求能成功输出 success:100
*/

type Ban struct {
	visitIPs map[string]struct{}
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]struct{})}
}

//判断IP是否存在
func (o *Ban) visit(ip string) bool {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = struct{}{}
	go o.invalidAfter3Min(ip)
	return false
}

// 3分钟后ip失效, 从map中移除. 因此ip再次访问时便可正常访问
func (o *Ban) invalidAfter3Min(ip string) {
	time.Sleep(time.Minute * 3)
	mapMutex.Lock()
	visitIPs := o.visitIPs
	delete(visitIPs, ip)
	o.visitIPs = visitIPs
	mapMutex.Unlock()
}

var mapMutex *sync.Mutex // mutex to avoid concurrent map writes

func TestSync(t *testing.T) {
	mapMutex = new(sync.Mutex)
	var success int64
	ban := NewBan()
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			ipEnd := j
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", ipEnd)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}()
		}
	}
	wg.Wait()

	if success == 100 {
		t.Log("Test case passed")
	} else {
		t.Error("Test case failed, Got ", success, "Expected 100")
	}
}
