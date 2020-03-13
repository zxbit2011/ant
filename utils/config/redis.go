package config

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisConnPool struct {
	redisPool   *redis.Pool
	projectName string
}

func InitRedis(red Redis, projectName string) (*RedisConnPool, error) {
	var Cache = &RedisConnPool{}
	Cache.projectName = projectName
	Cache.redisPool = newPool(red.Addr, red.Password, red.Database, red.MaxOpenConns, red.MaxIdleConns)

	if Cache.redisPool == nil {
		return Cache, errors.New("InitRedis redisPool=nil failed！")
	}
	return Cache, nil
}

func newPool(server, password string, database, maxOpenConns, maxIdleConns int) *redis.Pool {
	return &redis.Pool{
		MaxActive:   maxOpenConns, // max number of connections
		MaxIdle:     maxIdleConns,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("select", database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// 关闭连接池
func (p *RedisConnPool) Close() error {
	err := p.redisPool.Close()
	return err
}

// 当前某一个数据库，执行命令
func (p *RedisConnPool) Do(command string, args ...interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

//// String（字符串）
func (p *RedisConnPool) SetString(key string, value interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("SET", p.getKey(key), value)
}

//// String（字符串）
func (p *RedisConnPool) SetAndExpire(key string, value interface{}, expire int64) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("SET", p.getKey(key), value)
	if err == nil {
		v, err = conn.Do("EXPIRE", p.getKey(key), expire)
	}
	return v, err
}

func (p *RedisConnPool) GetString(key string) (string, error) {
	// 从连接池里面获得一个连接
	conn := p.redisPool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer conn.Close()
	return redis.String(conn.Do("GET", p.getKey(key)))
}

func (p *RedisConnPool) GetBytes(key string) ([]byte, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", p.getKey(key)))
}

func (p *RedisConnPool) GetInt(key string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", p.getKey(key)))
}

func (p *RedisConnPool) GetInt64(key string) (int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("GET", p.getKey(key)))
}

//// Key（键）
func (p *RedisConnPool) DelKey(key string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("DEL", p.getKey(key))
}

func (p *RedisConnPool) ExpireKey(key string, seconds int64) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("EXPIRE", p.getKey(key), seconds)
}

func (p *RedisConnPool) Keys(pattern string) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

func (p *RedisConnPool) KeysByteSlices(pattern string) ([][]byte, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.ByteSlices(conn.Do("KEYS", pattern))
}

//// Hash（哈希表）
func (p *RedisConnPool) SetHashMap(key string, fieldValue map[string]interface{}) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("HMSET", redis.Args{}.Add(p.getKey(key)).AddFlat(fieldValue)...)
}

func (p *RedisConnPool) GetHashMapString(key string) (map[string]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", p.getKey(key)))
}

func (p *RedisConnPool) GetHashMapInt(key string) (map[string]int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.IntMap(conn.Do("HGETALL", p.getKey(key)))
}

func (p *RedisConnPool) GetHashMapInt64(key string) (map[string]int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int64Map(conn.Do("HGETALL", p.getKey(key)))
}

func (p *RedisConnPool) SetList(key string, list []string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("LPUSH", redis.Args{}.Add(p.getKey(key)).AddFlat(list)...)
}

func (p *RedisConnPool) SetStringList(key string, s string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("LPUSHX", redis.Args{}.Add(p.getKey(key)).AddFlat(s)...)
}

func (p *RedisConnPool) GetFristDel(key string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("BLPOP", p.getKey(key), 3)
}

func (p *RedisConnPool) GetLastDel(key string) (interface{}, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return conn.Do("BRPOP", p.getKey(key), 3)
}

func (p *RedisConnPool) GetListString(key string) ([]string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HGETALL", p.getKey(key), 0, -1))
}

func (p *RedisConnPool) SetSETString(key string, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("SADD", redis.Args{}.Add(p.getKey(key)).AddFlat(value)...))
}

func (p *RedisConnPool) CheckSETString(key, value string) (int, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("SISMEMBER", redis.Args{}.Add(p.getKey(key)).AddFlat(value)...))
}

func (p *RedisConnPool) GetSETRandStringRm(key string) (string, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("SPOP", p.getKey(key))
	return redis.String(v, err)
}

func (p *RedisConnPool) GetSETCount(key string) (int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("SCARD", p.getKey(key))
	return redis.Int64(v, err)
}

func (p *RedisConnPool) DelSETKeyValue(key string, value ...string) (int64, error) {
	conn := p.redisPool.Get()
	defer conn.Close()
	v, err := conn.Do("SREM", redis.Args{}.Add(p.getKey(key)).AddFlat(value)...)
	return redis.Int64(v, err)
}

func (p *RedisConnPool) getKey(key string) string {
	return fmt.Sprintf("%s_%s", p.projectName, key)
}
