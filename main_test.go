package gmt_test

import (
	"github.com/golang/mock/gomock"
	gmt "github.com/mattnolf/gomock-test"
	mock "github.com/mattnolf/gomock-test/mock"
	"testing"
)

func TestStart(t *testing.T) {
	t.Run("test the consumer runs", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockHandler := func(id string) error {
			return nil
		}

		mockConsumer := mock.NewMockConsumer(ctrl)
		mockConsumer.EXPECT().Consume("some-name", mockHandler)

		gmt.Start(mockConsumer, gmt.WithSomeName(mockHandler))
	})
}
