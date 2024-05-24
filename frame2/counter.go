/*
@author: sk
@date: 2023/6/22
*/
package frame

import (
	"fmt"
	"time"
)

type Counter struct {
	count, fps uint64
	lastSec    int64
}

func NewCounter() *Counter {
	return &Counter{lastSec: time.Now().Unix()}
}

func (c *Counter) GetFps() uint64 {
	return c.fps
}

func (c *Counter) Update() {
	c.count++
	sec := time.Now().Unix()
	if c.lastSec == sec {
		return
	}
	c.lastSec = sec
	c.fps = c.count
	c.count = 0
	fmt.Println(c.fps)
}
