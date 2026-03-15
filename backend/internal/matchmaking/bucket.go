package matchmaking

import "sync"

type bucketMap struct {
	mu sync.Mutex
	m  map[bucketID]*queue
}

func newBucketMap() *bucketMap {
	return &bucketMap{m: make(map[bucketID]*queue)}
}
