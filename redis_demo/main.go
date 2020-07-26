package main


import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100,  // 连接池大小
	})


	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func hgetDemo(){
	v, err := rdb.HGetAll("user").Result()
	if err != nil {
		// redis.Nil
		// 其他错误
		fmt.Printf("hgetall failed, err:%v\n", err)
		return
	}
	fmt.Println(v)

	v2 := rdb.HMGet("user", "name", "age").Val()
	fmt.Println(v2)

	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)

}

func redisExample2() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}


func watchDemo(){
	// 监视watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 业务逻辑
			time.Sleep(time.Second*5)
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err:%v\n", err)
		return
	}
	fmt.Println("tx exec success")
}

// nilDemo 使用pipeline查询多个key时某个key有空值
// 使用pipeline多次执行HGetAll 不会因为某个key没有值而出现err
func nilDemo(){
	rdb.HMSet("nilDemo:k1", map[string]interface{}{"name":"q1mi", "score":18})
	rdb.HMSet("nilDemo:k3", map[string]interface{}{"name":"q1mi", "score":18})

	pipeline := rdb.Pipeline()
	pipeline.HGetAll("nilDemo:k1")
	pipeline.HGetAll("nilDemo:k2")
	pipeline.HGetAll("nilDemo:k3")
	cmders, err := pipeline.Exec()
	if err != nil {
		fmt.Printf("nilDemo pipeline.Exec() failed, err:%v\n", err)
		return
	}
	for _, cmder := range cmders{
		fmt.Println(cmder == nil)
		v, ok := cmder.(*redis.StringStringMapCmd)
		if !ok{
			fmt.Println("cmder.(*redis.StringStringMapCmd) failed")
			continue
		}
		// 打印值
		fmt.Println(v.Val())
	}
}



// nilDemo2 使用pipeline查询多个key时某个key有空值
// 使用pipeline多次执行 Get 会因为某个key没有值而出现redis.Nil的err
func nilDemo2(){
	rdb.Set("nilDemo2:k1", "v1", 0)
	rdb.Set("nilDemo2:k3", "v3", 0)

	pipeline := rdb.Pipeline()
	pipeline.Get("nilDemo2:k1")
	pipeline.Get("nilDemo2:k2")
	pipeline.Get("nilDemo2:k3")
	cmders, err := pipeline.Exec()
	if err != nil {
		fmt.Printf("nilDemo2 pipeline.Exec() failed, err:%v\n", err)
		return
	}
	for _, cmder := range cmders{
		fmt.Println(cmder == nil)
		v, ok := cmder.(*redis.StringCmd)
		if !ok{
			fmt.Println("cmder.(*redis.StringCmd) failed")
			continue
		}
		// 打印值
		fmt.Println(v.Val())
	}
}

// nilDemo3 使用pipeline查询多个key时某个key有空值
// 使用pipeline多次执行 HGet 会因为某个key没有值而出现redis.Nil的err
func nilDemo3(){
	rdb.HMSet("nilDemo3:k1", map[string]interface{}{"name":"q1mi", "score":18})
	rdb.HMSet("nilDemo3:k3", map[string]interface{}{"name":"q1mi", "score":18})

	pipeline := rdb.Pipeline()
	pipeline.HGet("nilDemo3:k1", "score")
	pipeline.HGet("nilDemo3:k2", "score")
	pipeline.HGet("nilDemo3:k3", "score")
	cmders, err := pipeline.Exec()
	if err != nil {
		fmt.Printf("nilDemo3 pipeline.Exec() failed, err:%v\n", err)
		return
	}
	for _, cmder := range cmders{
		fmt.Println(cmder == nil)
		v, ok := cmder.(*redis.StringCmd)
		if !ok{
			fmt.Println("cmder.(*redis.StringCmd) failed")
			continue
		}
		// 打印值
		fmt.Println(v.Val())
	}
}

// nilDemo4 使用pipeline查询多个key中某个没有值的field
// 使用pipeline多次执行 HGet 会因为某个field没有值而出现redis.Nil的err
func nilDemo4(){
	rdb.HMSet("nilDemo4:k1", map[string]interface{}{"name":"q1mi", "score":18})
	rdb.HMSet("nilDemo4:k2", map[string]interface{}{"name":"q1mi"})
	rdb.HMSet("nilDemo4:k3", map[string]interface{}{"name":"q1mi", "score":18})

	pipeline := rdb.Pipeline()
	pipeline.HGet("nilDemo4:k1", "score")
	pipeline.HGet("nilDemo4:k2", "score")
	pipeline.HGet("nilDemo4:k3", "score")
	cmders, err := pipeline.Exec()
	if err != nil && err != redis.Nil {
		fmt.Printf("nilDemo4 pipeline.Exec() failed, err:%v\n", err)
		return
	}
	for _, cmder := range cmders{
		fmt.Println(cmder == nil)
		v, ok := cmder.(*redis.StringCmd)
		if !ok{
			fmt.Println("cmder.(*redis.StringCmd) failed")
			continue
		}
		// 打印值
		fmt.Println(v.Val())
	}
}

func main() {
	if err := initClient();err!=nil{
		fmt.Printf("init redis client failed, err:%v\n", err)
		return
	}
	fmt.Println("connect redis success...")
	// 程序退出时释放相关资源
	defer rdb.Close()


	//redisExample()
	//hgetDemo()
	//redisExample2()
	//watchDemo()
	nilDemo()
	nilDemo2()
	nilDemo3()
	nilDemo4()
}
