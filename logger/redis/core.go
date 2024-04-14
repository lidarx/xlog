package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xlog/encoder"
)

func newRedisWriter(key string, cli *redis.Client) *redisWriter {
	return &redisWriter{
		cli: cli, listKey: key,
	}
}

// RedisWriter 为 logger 提供写入 redis 队列的 io 接口
type redisWriter struct {
	cli     *redis.Client
	listKey string
}

func (w *redisWriter) Write(p []byte) (int, error) {
	n, err := w.cli.RPush(w.listKey, p).Result()
	return int(n), err
}

func NewCore(level zapcore.Level, redisClient *redis.Client, key string, encoderConfig zapcore.EncoderConfig, extra ...zapcore.Field) (core zapcore.Core) {
	redisWriter := newRedisWriter(key, redisClient)
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	return zapcore.NewCore(jsonEncoder, zapcore.AddSync(redisWriter), zap.NewAtomicLevelAt(level)).With(extra)
}

func NewDefaultCore(level zapcore.Level, redisClient *redis.Client, key string, extra ...zapcore.Field) (core zapcore.Core) {
	return NewCore(level, redisClient, key, encoder.DefaultRedisEncoderConfig(), extra...)
}
