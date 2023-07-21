package cell

import (
	"github.com/xi123/libgo/core/base/run"
	"github.com/xi123/libgo/core/net/conn"
)

// <summary>
// Worker 业务处理单元
// <summary>
type Worker interface {
	OnInit()
	OnTimer(timerID uint32, dt int32, args ...any) bool
}

// <summary>
// WorkerCreator
// <summary>
type WorkerCreator interface {
	Create(proc run.Proc, args ...any) Worker
}

// <summary>
// NetWorker
// <summary>
type NetWorker interface {
	Worker
	OnRead(cmd uint32, msg any, peer conn.Session)
	OnCustom(cmd uint32, msg any, peer conn.Session)
}

// <summary>
// NetWorkerCreator
// <summary>
type NetWorkerCreator interface {
	WorkerCreator
}
