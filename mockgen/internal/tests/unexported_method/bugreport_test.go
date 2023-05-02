package bugreport

import (
	"testing"

	"github.com/neutrinocorp/mock/gomock"
)

func TestCallExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := NewMockExample(ctrl)
	e.EXPECT().someMethod(gomock.Any()).Return("it works!")

	CallExample(e)
}
