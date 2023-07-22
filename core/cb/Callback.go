package cb

import (
	"net"
	"net/http"

	"github.com/xi163/libgo/core/net/conn"
	"github.com/xi163/libgo/core/net/transmit"
	"github.com/xi163/libgo/utils/timestamp"
)

type OnHandshake func(w http.ResponseWriter, r *http.Request) bool

type OnCondition func(addr net.Addr) bool

type OnProtocol func(proto string) transmit.Channel

type OnNewConnection func(conn any, channel transmit.Channel, protoName string, v ...any)

type OnConnected func(peer conn.Session, v ...any)

type OnClosed func(peer conn.Session, reason conn.Reason)

type OnMessage func(peer conn.Session, msg any, recvTime timestamp.T)

type OnWriteComplete func(peer conn.Session)

type CloseCallback func(peer conn.Session)

type ErrorCallback func(err error)

type ReadCallback func(cmd uint32, msg any, peer conn.Session)

type CustomCallback func(cmd uint32, msg any, peer conn.Session)

type CmdCallback func(msg any, peer conn.Session)

type CmdCallbacks map[uint32]CmdCallback

type TimerCallback func(timerID uint32, dt int32, args ...any) bool

type Processor func(msg any, args ...any) bool
