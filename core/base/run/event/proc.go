package event

import (
	"time"

	"github.com/xi163/libgo/core/cb"
	"github.com/xi163/libgo/core/net/conn"
)

// <summary>
// Proc 事件处理单元
// <summary>
type Proc interface {
	Post(data *Data)
	PostRead(cmd uint32, msg any, peer conn.Session)
	PostReadWith(handler cb.ReadCallback, cmd uint32, msg any, peer conn.Session)
	PostCustom(cmd uint32, msg any, peer conn.Session)
	PostCustomWith(handler cb.CustomCallback, cmd uint32, msg any, peer conn.Session)
	PostClosing(d time.Duration, peer conn.Session)
}
