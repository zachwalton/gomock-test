package gmt_test

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	gmt "github.com/mattnolf/gomock-test"
	mock "github.com/mattnolf/gomock-test/mock"
	"testing"
)

func TestStart(t *testing.T) {
	t.Run("test the consumer runs", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		errCouldNotConsume := errors.New("could not consume")

		mockHandler := func(id string) error {
			return errCouldNotConsume
		}

		mockConsumer := mock.NewMockConsumer(ctrl)
		mockConsumer.EXPECT().Consume("some-name", newHandlerMatcher(mockHandler))

		gmt.Start(mockConsumer, gmt.WithSomeName(mockHandler))
	})
}

type handlerMatcher struct {
	handler gmt.ConsumerHandler
}

var _ gomock.Matcher = &handlerMatcher{}

// Matches implements gomock.Matcher
func (fm *handlerMatcher) Matches(x interface{}) bool {
	handler, ok := x.(gmt.ConsumerHandler)
	if !ok {
		return false
	}

	// compare output of expected handler and handler invoked by gmt.Start()
	return fm.handler("") == handler("")
}

// String implements gomock.Matcher
func (fm handlerMatcher) String() string {
	return fmt.Sprintf("handler: %+v", fm.handler)
}

func newHandlerMatcher(handler gmt.ConsumerHandler) *handlerMatcher {
	return &handlerMatcher{handler: handler}
}
