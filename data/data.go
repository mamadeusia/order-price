package data

import (
	"context"
	"crypto/sha256"
	"encoding"
	"encoding/json"
	"fmt"
	"preh/server/db"
)

type Redisable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	CacheHashKey() string
	CacheHashField() string
}

func setToRedis(ctx context.Context, r Redisable) error {
	fmt.Println("set to redis ", r)
	return db.SetHToRedis(ctx, r.CacheHashKey(), r.CacheHashField(), r)
}

func getFromRedis(ctx context.Context, r Redisable) error {
	bytes, err := db.GetHFromRedis(ctx, r.CacheHashKey(), r.CacheHashField())
	json.Unmarshal(bytes, &r)
	fmt.Println(r)
	return err
}

func updateToRedis(ctx context.Context, r Redisable) error {
	return db.UpdateHRedis(ctx, r.CacheHashKey(), r.CacheHashField(), r)
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
