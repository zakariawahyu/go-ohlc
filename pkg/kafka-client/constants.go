package kafka_client

import "time"

const (
	writerReadTimeout  = 1 * time.Second
	writerWriteTimeout = 1 * time.Second
	batchTimeout       = 60 * time.Millisecond
	batchSize          = 100
	writerMaxAttempts  = 10
)
