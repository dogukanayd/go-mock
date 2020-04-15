package greeter

import (
	"github.com/dogukanayd/go-mock/example1/greeter/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreetInDefaultMsg(t *testing.T) {
	theDBMock := mocks.DB{} // create the mock
	theDBMock.On("FetchDefaultMessage").Return("default", nil)

	g := Greeter{
		&theDBMock,
		"en",
	}

	assert.Equal(t, "Message is default", g.GreetInDefaultMsg())
	theDBMock.AssertNumberOfCalls(t, "FetchDefaultMessage", 1)
	theDBMock.AssertExpectations(t)
}

func TestGreet(t *testing.T) {
	theDBMock := mocks.DB{}
	theDBMock.On("FetchMessage", "sg").Return("lah", nil) // if FetchMessage("sg") is called, then return "lah"

	g := Greeter{&theDBMock, "sg"}

	assert.Equal(t, "Message is lah", g.Greet())
	theDBMock.AssertExpectations(t)
}

func TestNewGreeter(t *testing.T) {
	theDBMock := mocks.DB{}

	actual := Greeter{&theDBMock, "en"}
	expected := NewGreeter(&theDBMock, "en")

	assert.Equal(t, actual, expected)
}