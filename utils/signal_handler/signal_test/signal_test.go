package signal_test

import (
	"testing"

	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils/signal_handler"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("signal_handler_test:", signal_handler_test)
}

func signal_handler_test(t *testing.T) {
	signal_handler.RegisterStop(func() {
		logs.Debugf("Stopping..")
	})
	signal_handler.Wait()
}
