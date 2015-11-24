package ratecounter

import (
	"testing"
	"time"
)

func TestAvgCounter(t *testing.T) {
	interval := 1 * time.Second
	r := NewAvgCounter(interval)

	check := func(expected int64) {
		val := r.Avg()
		if val != expected {
			t.Error("Expected ", val, " to equal ", expected)
		}
	}

	check(0)
	r.Incr(350)
	check(350)
	r.Incr(450)
	check(400)
	time.Sleep(2 * interval)
	check(0)

	r.Incr(500)
	time.Sleep(500 * time.Millisecond)
	r.Incr(1000)
	r.Incr(1000)
	check(833)
	time.Sleep(500 * time.Millisecond)
	check(1000)
}
