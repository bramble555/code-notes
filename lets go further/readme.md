## JSON Encoding

if you want to omit the field when the field is empty, you can

use this way.

```
createAt time.Time `json:"-"`
or 
createAt time.Time `json:"omitenmpy"`
```

## Advanced JSON Customization

 The json.Marshaler interface can cater to your needs when you customize json.

```
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```

notes：

1. You must use a '' envelop the []byte. You can choose use strconv.Quote(str).

   like this

   ```
   return []byte(strconv.Quote(str)),nil
   ```

   Or you cant get the answer you want

2. You had better use a value receiver for MarshalJSON().Because 

   the value rceveiver can work values and points. If you deliberately use  a point receiver , it only can work points.

## JSON Decoding

While `json.Decoder` is more effective for handling large data due to its streaming approach, `json.Unmarshal` is the standard method and works well for simpler use cases



## Restricting Inputs

If you want to check for unknown keys in the request body, you can use `json.NewDecoder(r.Body).DisallowUnknownFields()`. However, `json.Marshal` doesn’t have a similar feature for this purpose. But you can code them.

## Validation Rules

https://github.com/DataDavD/greenlight/blob/main/internal/validator/validator.go

## Rate Limiting

### Global Rate Limiting

### IP-based Rate Limiting

## Graceful Shutdown of Background Tasks

Using `sync.WaitGroup` in Go can help ensure that tasks, such as sending emails, complete before the program shuts down, even when you close or interrupt the process. This way, you can achieve a "graceful shutdown" and prevent tasks from being abruptly cut off.

to do i cant implement it.

## Cross Origin Requests

"What is 'origin'? The same protocol, host, and port define a shared origin. The same-origin policy is often enforced in such cases."



[![source.png](https://i.postimg.cc/Lsg0vQWf/source.png)](https://postimg.cc/KKFN8fxv)

What is a cross-origin request?

If you have an API hosted at `api.example.com` and a trusted JavaScript front-end application running on `www.example.com`, you'll likely want to allow cross-origin requests from the trusted `www.example.com` domain to your API.

But how do cross-origin requests work?

To begin with, the simplest way to enable this is by including the following header in all your API responses:

```
Access-Control-Allow-Origin: *
```

What does this look like in Go Gin web code?

```Go
func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 Access-Control-Allow-Origin 响应头
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		
		// 继续处理请求
		c.Next()
	}
}


r.Use(enableCORS())
```

Notes:

- If you have a rate limiter `HandlerFunc`, ensure that `enableCORS` is used **before** the rate limiter in your middleware chain.
- Of course, if you're targeting only a specific website, you can simply replace the website's origin with `*`.
- However, if you want to allow multiple specific websites but avoid using `*`, things get more complicated. Unfortunately, web browsers do not support the use of multiple space-separated values in the `Access-Control-Allow-Origin` header.
- But don't worry! We can handle this using a command-line flag. By leveraging the `strings.Fields()` function, you can split the origin values into a `[]string` slice for easy processing.
- Here's an example of how to achieve this:

```Go
config := Config{
		TrustedOrigins: []string{
			"http://example.com",
			"http://trusted.com",
		},
	}
	
	
// enableCORS 中间件
func enableCORS(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 添加 Vary: Origin 响应头
		c.Writer.Header().Add("Vary", "Origin")

		// 获取请求头中的 Origin 值
		origin := c.Request.Header.Get("Origin")

		// 如果 Origin 存在，且有配置受信任的来源
		if origin != "" && len(config.TrustedOrigins) > 0 {
			// 遍历受信任的来源列表，检查是否匹配
			for _, trustedOrigin := range config.TrustedOrigins {
				if origin == trustedOrigin {
					// 如果匹配，设置 Access-Control-Allow-Origin 响应头
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		// 继续处理后续的中间件或请求处理函数
		c.Next()
	}
}

r.Use(enableCORS(config))
```

"OK, if you don't modify the headers, it's probably a simple request.

A simple request in CORS meets the following conditions:

- The request method is GET, HEAD, or POST.
- The request headers are all either forbidden headers or one of the four CORS-safe headers: Accept , Accept-Language ,Content-Language , Content-Type
- The value of `Content-Type` in the request header  is either `text/plain`, `multipart/form-data`, or `application/x-www-form-urlencoded`.
- It doesn't contain custom request headers, such as `Authorization` or `X-Custom-Header`."

When a cross-origin request doesn’t meet these conditions, then the web browser will trigger an initial ‘preflight’ request before the real request.

There are three headers here which are relevant to CORS.

- Origin
- Access-Control-Request-Method 
- Access-Control-Request-Headers

Notes :

​	It’s important to note that Access-Control-Request-Headers won’t list all the headers that the real request will use. Only headers that are not CORS-safe or forbidden will be listed

To respond to a preflight request, the first thing we need to do is confirm that it is a preflight request, rather than a regular (possibly cross-origin) OPTIONS request.
We can identify a preflight request by the following three criteria:

- The request method must be OPTIONS.

- The request header must include Origin.

- The request header must include Access-Control-Request-Method.

If any of these three components is missing, we can determine that the request is not a preflight request.

Let update the above example .

```Go
func (app *application) enableCORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 设置Vary头
        c.Header("Vary", "Origin, Access-Control-Request-Method")

        // 获取请求的Origin头
        origin := c.GetHeader("Origin")
        
        // 检查Origin是否存在并且受信任的来源配置不为空
        if origin != "" && len(app.config.cors.trustedOrigins) != 0 {
            for _, trustedOrigin := range app.config.cors.trustedOrigins {
                if origin == trustedOrigin {
                    // 设置Access-Control-Allow-Origin头
                    c.Header("Access-Control-Allow-Origin", origin)
                    
                    // 设置Access-Control-Allow-Credentials头
                    c.Header("Access-Control-Allow-Credentials", "true")

                    // 检查请求方法是否为OPTIONS，且请求头中包含Access-Control-Request-Method，表明是预检请求
                    if c.Request.Method == http.MethodOptions && c.GetHeader("Access-Control-Request-Method") != "" {
                        // 设置预检请求的相关头
                        c.Header("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
                        c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

                        // 返回200 OK并结束处理
                        c.JSON(http.StatusOK, gin.H{})
                        return
                    }
                    break
                }
            }
        }

        // 调用下一层的处理器
        c.Next()
    }
}

```

Notes :

- **`c.Header` replaces the old value with the new one.  **c.Writer.HeaderAdd`  appends an additional value.
- When Access-Control-Allow-Credentials taking Cookie or JWT  is true, Access-Control-Allow-Origin is not *.

# Metrics

## stack and  heap 

OS thread stack is fixed,a default of 8MB.

goroutine stacks start with a small amount of memory (currently 2KB).

```
$ go test -bench . -benchmem
BenchmarkStackIt-8  680439016  1.52 ns/op  0 B/op  0 allocs/op

0 allocs/op 说明堆上没有
```

Usually,a general rule we can infer from this is that sharing pointers up the stack results in allocations, whereas sharing points down the stack doesn’t.

```
func stackIt2() *int {
   y := 2
   res := y * 2
   return &res
}
&res  will allocate heap


func stackIt3(y *int) int {
   res := *y * 2
   return res
}
*y will allocate stack
```

However, this is not guaranteed, so you’ll still need to verify with benchmarks to be sure. 

## `expvar`package

`expvar`package can list some parameter.

```
TotalAlloc — Cumulative bytes allocated on the heap (will not decrease).

HeapAlloc — Current number of bytes on the heap.

HeapObjects — Current number of objects on the heap.

Sys — Total bytes of memory obtained from the OS (i.e. total memory reserved by the Go

runtime for the heap, stacks, and other internal data structures).

NumGC — Number of completed garbage collector cycles.

NextGC — The target heap size of the next garbage collector cycle (Go aims to keep
HeapAlloc ≤ NextGC).
```

for example

```
package router

import (
	"expvar"
	"net/http"
	"runtime"
	"time"

	"github.com/bramble555/blog/global"
	"github.com/gin-gonic/gin"
)

func InitMetricsRoutes(r *gin.RouterGroup) gin.IRoutes {
	// Publish the number of active goroutines.
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))

	// Publish the database connection pool statistics (for GORM).
	expvar.Publish("database", expvar.Func(func() interface{} {
		sqlDB, err := global.DB.DB() // 注意这里是 global.DB，而不是 global.db
		if err != nil {
			return err.Error()
		}
		return sqlDB.Stats() // 返回 *sql.DB 的连接池统计
	}))

	// Publish the current time.
	expvar.Publish("timeNow", expvar.Func(func() interface{} {
		return time.Now().Format("2006-01-02 15:04:05")
	}))

	// Publish Go memory stats
	expvar.Publish("memoryStats", expvar.Func(func() interface{} {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		return map[string]uint64{
			"Alloc":        memStats.Alloc,         // 当前已分配的字节数
			"HeapIdle":     memStats.HeapIdle,      // 空闲堆内存
			"HeapInuse":    memStats.HeapInuse,     // 使用中的堆内存
			"HeapReleased": memStats.HeapReleased,  // 已释放的堆内存
			"HeapSys":      memStats.HeapSys,       // 堆的总内存
			"NumGC":        uint64(memStats.NumGC), // 完成的 GC 循环数
			"LastGC":       memStats.LastGC,        // 最近一次 GC 的时间
			"Lookups":      memStats.Lookups,       // 全局符号表查找次数
			"GCSys":        memStats.GCSys,         // 垃圾回收使用的内存
		}
	}))

	// 直接在 Gin 中处理 /api/debug/vars 路径的请求
	r.GET("/debug/vars", func(c *gin.Context) {
		// 获取 expvar.Func 并调用它，获得实际值
		goroutines := expvar.Get("goroutines").(expvar.Func)()
		database := expvar.Get("database").(expvar.Func)()
		timeNow := expvar.Get("timeNow").(expvar.Func)()
		memoryStats := expvar.Get("memoryStats").(expvar.Func)()

		// 返回 JSON 格式的响应
		c.JSON(http.StatusOK, map[string]interface{}{
			"goroutines":  goroutines,
			"database":    database,
			"timeNow":     timeNow,
			"memoryStats": memoryStats,
		})
	})

	return r
}

```

