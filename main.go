//go:generate mockgen -package=gmt_mock -destination=mock/consumer_mock.go -source=main.go Consume

package gmt

type ConsumerHandler func(id string) error

type ConsumerOption func(consumer Consumer) error

// WithSomeName is a ConsumerOption
func WithSomeName(handler ConsumerHandler) ConsumerOption {
	return func(consumer Consumer) error {
		return consumer.Consume("some-name", handler)
	}
}

type Consumer interface {
	Consume(name string, handler ConsumerHandler) error
}

// an implementation of Consumer
type mockConsumer struct {}
func (mc mockConsumer) Consume(name string, handler ConsumerHandler) error {
	return nil
}


func Start(c Consumer, opt ConsumerOption) {
		opt(c)
}

func main() {
	mockConsumer := mockConsumer{}
	mockHandler := func(id string) error {
		return nil
	}

	Start(mockConsumer, WithSomeName(mockHandler))
}
