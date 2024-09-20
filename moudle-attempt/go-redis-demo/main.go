package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	_, err := rdb.Ping().Result()
	return err
}
func main() {
	err := Init()
	if err != nil {
		fmt.Println("redis init error", err.Error())
		return
	}
	fmt.Println("redis init succeed")
	defer rdb.Close()
	// 键操作
	err = rdb.Set("name", "ck", 0).Err()
	if err != nil {
		fmt.Println("redis set name err", err.Error())
	}
	val, err := rdb.Get("name").Result()
	if err != nil {
		fmt.Println("redis get name err", err.Error())
	}
	fmt.Println("name is ", val)
	val, err = rdb.Get("age").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("age is not exist", err.Error())
		} else {
			fmt.Println("redis get age err", err.Error())
		}

	} else {
		fmt.Println("age is ", val)
	}
	// 哈希操作
	err = rdb.HSet("students", "name", "ckk").Err()
	if err != nil {
		fmt.Println("redis hset name err", err.Error())
	}
	err = rdb.HSet("students", "age", "19").Err()
	if err != nil {
		fmt.Println("redis hset age err", err.Error())
	}
	name, err := rdb.HGet("students", "name").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("name is not exist", err.Error())
		} else {
			fmt.Println("redis get name err", err.Error())
		}
	}
	fmt.Println("name is ", name)
	stds, err := rdb.HGetAll("students").Result()
	if err != nil {
		fmt.Println("redis hset studentserr", err.Error())
	}
	fmt.Println(stds)
	res := rdb.HMGet("students", "name", "age").Val()
	fmt.Println(res...)
	// 有序集合操作
	err = rdb.ZAdd("scores",
		redis.Z{Member: "alice", Score: 100},
		redis.Z{Member: "berry", Score: 200},
	).Err()
	if err != nil {
		fmt.Println("redis zset scores serr", err.Error())
	}
	scores, err := rdb.ZRangeWithScores("scores", 0, -1).Result()
	if err != nil {
		fmt.Println("获取有序集合元素失败:", err)
		return
	}
	for _, z := range scores {
		fmt.Printf("%s: %.1f\n", z.Member, z.Score)
	}
	// 给berry + 10分
	err = rdb.ZIncrBy("scores", 10, "berry").Err()
	if err != nil {
		fmt.Println("redis zset scores berry increase 10 serr", err.Error())
	}
	scores, _ = rdb.ZRangeWithScores("scores", 0, -1).Result()
	for _, z := range scores {
		fmt.Printf("%s: %.1f\n", z.Member, z.Score)
	}
	// 取200分以上
	op := redis.ZRangeBy{
		Min: "200",
		Max: "+inf",
	}
	scores, _ = rdb.ZRangeByScoreWithScores("scores", op).Result()
	for _, z := range scores {
		fmt.Printf("%s: %.1f\n", z.Member, z.Score)
	}
	// pipeline 进行网络性能优化(同一时间有很多命令同时发出)
	osScore := rdb.Set("db", 100, 0).Val()
	fmt.Println(osScore)
	// 初始化 pipeline
	pipe := rdb.Pipeline()
	osScoreInt64 := pipe.IncrBy("db", 10).Val()
	pipe.Expire("db", time.Hour)
	// 俩条命令一起执行
	_, err = pipe.Exec()
	fmt.Println(osScoreInt64, err)
	// rdb.TxPipeline() 这个会开启事务，会整体执行

	// watch监控，如何原来key是某个值，我再去执行操作，负责不执行(抢购)
	key := "watch_count"
	err = rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 业务逻辑
		pipe = tx.Pipeline()
		pipe.Set(key, n+1, 0)
		_, err = pipe.Exec()
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		fmt.Println("Error executing transaction:", err)
		return
	}

	// 在事务成功执行后获取最新的键值
	re, err := rdb.Get(key).Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return
	}

	fmt.Println("Current value:", re)
	fmt.Println("Transaction completed successfully.")
}
