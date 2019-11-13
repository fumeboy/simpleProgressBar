package simpleProgress

import (
	"fmt"
	"time"
)

type progress struct {
	cap          int64
	len          int64
	throttle     time.Duration
	throttleLock bool
	LText        string
	RText        string
}

func (p *progress) Write(b []byte) (n int, err error) {
	n = len(b)
	p.Add(n)
	return
}

func (p *progress) Add(l int) {
	p.len += int64(l)
	if !p.throttleLock {
		p.throttleLock = true
		go p.Print()
	}
}

func drawProgressBar(l int, c int) string {
	var ret = make([]byte, 30)
	var l_ = (l * 30) / c
	if l_ > 30{
		l_ = 30
	}
	for i := 0; i < l_; i++ {
		ret[i] = '$'
	}
	if l_ < 30{
		ret[l_] = '>'
		for i := l_+1; i < 30; i++ {
			ret[i] = ' '
		}
	}
	return string(ret)
}

func (p *progress) Print() {
	fmt.Printf("\r%s |%s| %d/%d %s", p.LText, drawProgressBar(int(p.len), int(p.cap)), int(p.len), int(p.cap), p.RText)
	time.Sleep(p.throttle)
	p.throttleLock = false
}

func InitProgress(cap int64) (p *progress) {
	p = &progress{
		cap:      cap,
		len:      0,
		throttle: 10 * time.Millisecond,
	}
	return
}
