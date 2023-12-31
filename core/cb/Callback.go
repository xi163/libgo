package cb

import (
	"net"
	"net/http"

	"github.com/cwloo/gonet/core/net/conn"
	"github.com/cwloo/gonet/core/net/transmit"
	"github.com/cwloo/gonet/utils/timestamp"
)

type OnVerify func(w http.ResponseWriter, r *http.Request) bool

type OnCondition func(peerAddr net.Addr, peerRegion *conn.Region) bool

type OnProtocol func(proto string) transmit.Channel

type OnNewConnection func(conn any, channel transmit.Channel, protoName string, peerRegion *conn.Region, v ...any)

type OnConnectError func(proto string, err error)

type OnConnected func(peer conn.Session, v ...any)

type OnClosed func(peer conn.Session, reason conn.Reason, v ...any)

type OnMessage func(peer conn.Session, msg any, msgType int, recvTime timestamp.T)

type OnWriteComplete func(peer conn.Session)

type CloseCallback func(peer conn.Session)

type ErrorCallback func(err error)

type ReadCallback func(cmd uint32, msg any, peer conn.Session)

type CustomCallback func(cmd uint32, msg any, peer conn.Session)

type CmdCallback func(msg any, peer conn.Session)

type CmdCallbacks map[uint32]CmdCallback

type TimerCallback func(timerID uint32, dt int32, args ...any) bool

type Processor func(msg any, args ...any) bool
