package ch

import "github.com/xi163/libgo/core/base/mq"

// <summary>
// Queue chan消息队列
// <summary>
type Queue interface {
	mq.Queue
	Read() <-chan any
	Signal() <-chan bool
	Full() bool
	Length() int
	AssertEmpty()
	Reset()
}
