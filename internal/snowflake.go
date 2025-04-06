package internal

import (
	"sync"
	"time"
)

var (
	Epoch    int64 = 1288834974657
	NodeBits uint8 = 10
	StepBits uint8 = 12
)

type ID int64

type Node struct {
	mu        sync.Mutex
	epoch     time.Time
	time      int64
	node      int64
	step      int64
	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

func NewSnowflakeNode(node int64) *Node {
	n := Node{}
	n.node = node
	n.nodeMax = -1 ^ (-1 << NodeBits)
	n.nodeMask = n.nodeMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeShift = NodeBits + StepBits
	n.nodeShift = StepBits

	var curTime = time.Now()
	// adicionar time.Duration a cur.Time para ter certeza que monotonic clock está disponível
	n.epoch = curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime))

	return &n
}

func (n *Node) GenerateID() ID {

	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.epoch).Milliseconds()

	if now == n.time {
		n.step = (n.step + 1) & n.stepMask

		if n.step == 0 {
			for now <= n.time {
				now = time.Since(n.epoch).Milliseconds()
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	r := ID((now)<<n.timeShift |
		(n.node << n.nodeShift) |
		(n.step),
	)
	return r
}
