package simpleProgress

import (
	"fmt"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	bar := InitProgress(100)
	bar.LText = "LLL"
	bar.RText = "RRR"
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(time.Millisecond*100)
	}
	fmt.Println("")
}
