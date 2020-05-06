package decorator

import (
	"log"
	"os"
	"testing"
)

func TestDecorator(t *testing.T) {
	piFunc := LogWrapper(Pi, log.New(os.Stdout, "Test-", 1))
	piFunc(10000000)
}
