package loadbalance

import (
	"math/rand"
	"time"
)

// 随机
func Rand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

// 总是选取第一个
func First() int {
	return 0
}

// 哈希ip，对长度取模
func Hash() int {
	return 0
}

// 一致性哈希
