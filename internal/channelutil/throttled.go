package channelutil

import "time"

type ThrottledChannel[T any] struct {
	in  <-chan T
	out chan T

	lastSentAt   time.Time
	sendInterval time.Duration
}

func NewThrottle[T any](in <-chan T, sendInterval time.Duration) *ThrottledChannel[T] {
	ch := &ThrottledChannel[T]{
		in:           in,
		out:          make(chan T),
		lastSentAt:   time.Now(),
		sendInterval: sendInterval,
	}
	ch.run()

	return ch
}

func (tc *ThrottledChannel[T]) GetReader() <-chan T {
	return tc.out
}

func (tc *ThrottledChannel[T]) run() {
	go func() {
		for v := range tc.in {
			if tc.lastSentAt.Add(tc.sendInterval).Before(time.Now()) {
				tc.out <- v
				tc.lastSentAt = time.Now()
			}
		}

		close(tc.out)
	}()
}
