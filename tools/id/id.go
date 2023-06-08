//Package id
/*
@Title: id.go
@Description
@Author: kkw 2023/5/24 15:27
*/
package id

import (
	"sync"
	"time"
)

const (
	epoch      = 1600000000000 // 自定义的起始时间戳，单位为毫秒
	workerBits = 5             // 工作节点ID的位数
	seqBits    = 8             // 序列号的位数
	maxWorker  = -1 ^ (-1 << workerBits)
	maxSeq     = -1 ^ (-1 << seqBits)
)

var (
	mutex    sync.Mutex
	lastTs   int64
	workerID uint
	sequence uint
)

func init() {
	workerID = 1
}

// GenerateID 生成有序的分布式ID
func GenerateID() int64 {
	mutex.Lock()
	defer mutex.Unlock()

	ts := time.Now().UnixNano() / 1e6 // 当前时间戳，单位为毫秒

	if ts == lastTs {
		sequence = (sequence + 1) & maxSeq
		if sequence == 0 {
			// 当前毫秒内的序列号用完，等待下一毫秒
			for ts <= lastTs {
				ts = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sequence = 0
	}

	lastTs = ts

	// 将各部分组合成一个64位的整数
	id := (ts-epoch)<<workerBits | int64(workerID)<<seqBits | int64(sequence)

	return id
}
