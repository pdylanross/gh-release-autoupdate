package channelutil

type Ringbuffer[T any] struct {
	in  <-chan T
	out chan T
}

func NewRingbuffer[T any](in <-chan T, size int) *Ringbuffer[T] {
	rb := &Ringbuffer[T]{
		in:  in,
		out: make(chan T, size),
	}
	go rb.run()

	return rb
}

func (rb *Ringbuffer[T]) GetReader() <-chan T {
	return rb.out
}

func (rb *Ringbuffer[T]) run() {
	for v := range rb.in {
		select {
		case rb.out <- v:
			continue
		default:
			<-rb.out
			rb.out <- v
		}
	}

	close(rb.out)
}
