package cache

type CacheStoreInterface interface {
	Get(key string) (string, error)
	Unmarshal(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
}

type RedisStoreInterface interface {
	GetByte(key string) ([]byte, error)
	Get(key string) (string, error)
	IncrBy(key string, value int64) (int64, error)
	DecrBy(key string, value int64) (int64, error)
	Unmarshal(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
	RPush(key string, value interface{}) error
	LPush(key string, value interface{}) error
	LLen(key string) (int64, error)
	LIndex(key string, index int64) (string, error)
	HSet(key string, field string, value interface{}) error
	HMSet(key string, fields map[string]interface{}) error
	HGet(key string, field string) (string, error)
	HLen(key string) (int64, error)
	HDel(key string, field string) error
	HExists(key string, field string) error
	HGetall(key string) (map[string]string, error)
	SAdd(key string, value string) error
	Do(cmd string, key string, seconds string) error
	SRandMember(key string) (string, error)
	Expire(key string) (bool, error)
}
