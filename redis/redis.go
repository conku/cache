package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// var (
// 	ctx = context.Background()
// )

// Redis provides a cache backed by a Redis server.
type Redis struct {
	Config *redis.Options
	Client *redis.Client
}

// New returns an initialized Redis cache object.
func New(config *redis.Options) *Redis {
	client := redis.NewClient(config)
	return &Redis{Config: config, Client: client}
}

// Get returns the value saved under a given key.
func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

// GetByte returns the value saved under a given key.
func (r *Redis) GetByte(key string) ([]byte, error) {
	return r.Client.Get(key).Bytes()
}

// IncrBy returns the value saved under a given key.
func (r *Redis) IncrBy(key string, value int64) (int64, error) {
	return r.Client.IncrBy(key, value).Result()
	//return r.Client.Get(key).Result()
}

// DecrBy returns the value saved under a given key.
func (r *Redis) DecrBy(key string, value int64) (int64, error) {
	return r.Client.DecrBy(key, value).Result()
	//return r.Client.Get(key).Result()
}

// Unmarshal retrieves a value from the Redis server and unmarshals
// it into the passed object.
func (r *Redis) Unmarshal(key string, object interface{}) error {
	value, err := r.Get(key)
	if err == nil {
		err = json.Unmarshal([]byte(value), object)
	}
	return err
}

// Set saves an arbitrary value under a specific key.
func (r *Redis) Set(key string, value interface{}) error {
	return r.Client.Set(key, convertToBytes(value), 0).Err()
}

func convertToBytes(value interface{}) []byte {
	switch result := value.(type) {
	case string:
		return []byte(result)
	case []byte:
		return result
	default:
		bytes, _ := json.Marshal(value)
		return bytes
	}
}

// Fetch returns the value for the key if it exists or sets and returns the value via the passed function.
func (r *Redis) Fetch(key string, fc func() interface{}) (string, error) {
	if str, err := r.Get(key); err == nil {
		return str, nil
	}
	results := convertToBytes(fc())
	return string(results), r.Set(key, results)
}

// Delete removes a specific key and its value from the Redis server.
func (r *Redis) Delete(key string) error {
	return r.Client.Del(key).Err()
}

// RPush 在名称为key的list尾添加一个值为value的元素
func (r *Redis) RPush(key string, value interface{}) error {
	return r.Client.RPush(key, value).Err()
}

// LPush 在名称为key的list头添加一个值为value的 元素
func (r *Redis) LPush(key string, value interface{}) error {
	return r.Client.LPush(key, value).Err()
}

// LLen 返回名称为key的list的长度
func (r *Redis) LLen(key string) (int64, error) {
	return r.Client.LLen(key).Result()
}

// LSet 给名称为key的list中index位置的元素赋值
func (r *Redis) LSet(key string, index int64, value interface{}) (string, error) {
	return r.Client.LSet(key, index, value).Result()
}

// LIndex 返回名称为key的list中index位置的元素
func (r *Redis) LIndex(key string, index int64) (string, error) {
	return r.Client.LIndex(key, index).Result()
}

// HSet 向名称为key的hash中添加元素field
func (r *Redis) HSet(key string, field string, value interface{}) error {
	return r.Client.HSet(key, field, value).Err()
}

// HMSet 向名称为map的hash中添加元素field
func (r *Redis) HMSet(key string, fields map[string]interface{}) error {
	return r.Client.HMSet(key, fields).Err()
}

// HGet 返回名称为key的hash中field对应的value
func (r *Redis) HGet(key string, field string) (string, error) {
	return r.Client.HGet(key, field).Result()
}

// HLen 返回名称为key的list的长度
func (r *Redis) HLen(key string) (int64, error) {
	return r.Client.HLen(key).Result()
}

// HGetall 返回名称为key的hash中所有的键（field）及其对应的value
func (r *Redis) HGetall(key string) (map[string]string, error) {
	return r.Client.HGetAll(key).Result()
}

// HDel 返回名称为key的hash中field对应的value
func (r *Redis) HDel(key string, field string) error {
	return r.Client.HDel(key, field).Err()
}

// HExists 返回名称为key的hash中field对应的value
func (r *Redis) HExists(key string, field string) error {
	return r.Client.HExists(key, field).Err()
}

func (r *Redis) SAdd(key string, field string) error {
	return r.Client.SAdd(key, field).Err()
}

func (r *Redis) SRandMember(key string) (string, error) {
	return r.Client.SRandMember(key).Result()
}

// Expire 设置过期时间
func (r *Redis) Expire(key string, expiration time.Duration) (bool, error) {
	return r.Client.Expire(key, expiration).Result()
}

//Do 设置过期时间 EXPIRE/PEXPIREAT
/*
EXPIRE aa 60 接口定义：EXPIRE key "seconds"
接口描述：设置一个key在当前时间"seconds"(秒)之后过期。返回1代表设置成功，返回0代表key不存在或者无法设置过期时间。

PEXPIRE 接口定义：PEXPIRE key "milliseconds"
接口描述：设置一个key在当前时间"milliseconds"(毫秒)之后过期。返回1代表设置成功，返回0代表key不存在或者无法设置过期时间。

EXPIREAT aa 1586941008 接口定义：EXPIREAT key "timestamp"
接口描述：设置一个key在"timestamp"(时间戳(秒))之后过期。返回1代表设置成功，返回0代表key不存在或者无法设置过期时间。

PEXPIREAT aa 1586941008000 接口定义：PEXPIREAT key "milliseconds-timestamp"
接口描述：设置一个key在"milliseconds-timestamp"(时间戳(毫秒))之后过期。返回1代表设置成功，返回0代表key不存在或者无法设置过期时间

TTL 接口定义：TTL key
　　　　接口描述：获取key的过期时间。如果key存在过期时间，返回剩余生存时间(秒)；如果key是永久的，返回-1；如果key不存在或者已过期，返回-2。

PTTL 接口定义：PTTL key
　　　　接口描述：获取key的过期时间。如果key存在过期时间，返回剩余生存时间(毫秒)；如果key是永久的，返回-1；如果key不存在或者已过期，返回-2。

PERSIST 接口定义：PERSIST key
　　　　接口描述：移除key的过期时间，将其转换为永久状态。如果返回1，代表转换成功。如果返回0，代表key不存在或者之前就已经是永久状态。

SETEX 接口定义：SETEX key "seconds" "value"
　　接口描述：SETEX在逻辑上等价于SET和EXPIRE合并的操作，区别之处在于SETEX是一条命令，而命令的执行是原子性的，所以不会出现并发问题。
*/
func (r *Redis) Do(cmd, key, seconds string) error {
	return r.Client.Do(cmd, key, seconds).Err()
}
