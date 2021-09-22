package mwcoreapi

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(redis *redis.Client) *Cache {
	return &Cache{
		rdb: redis,
	}
}
func (c *Cache) SetKey(key string, value interface{}) error {
	CacheEntry, err := json.Marshal(value)
	if err != nil {
		//logger.Error(err.Error(), "SetKey")
		return err
	}
	err = c.rdb.Set(context.Background(), key, CacheEntry, 0).Err()
	if err != nil {
		//logger.Error(err.Error(), "SetKey")
		return err
	}
	return nil
}
func (c *Cache) GetKey(key string, src interface{}) error {
	val, err := c.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		//logger.Error(err.Error(), "GetKey")
		return err
	}
	return nil
}
