package ratecounter

import (
	"strconv"
	"time"
)

type AvgCounter struct {
	accumulator Counter
	occurs      Counter
	interval    time.Duration
}

func NewAvgCounter(intervl time.Duration) *AvgCounter {
	return &AvgCounter{
		interval: intervl,
	}
}

func (r *AvgCounter) Incr(val int64) {
	r.accumulator.Incr(val)
	r.occurs.Incr(1)
	go r.scheduleDecrement(val)
}

func (r *AvgCounter) scheduleDecrement(amount int64) {
	time.Sleep(r.interval)
	r.accumulator.Incr(-1 * amount)
	r.occurs.Incr(-1)
}

func (r *AvgCounter) Avg() int64 {
	denom := r.occurs.Value()
	if denom == 0 {
		return 0
	}
	return r.accumulator.Value() / denom
}

func (r *AvgCounter) String() string {
	return strconv.FormatInt(r.Avg(), 10)
}
