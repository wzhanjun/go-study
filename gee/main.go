package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	gcache "trie-demo/cache"
)

type student struct {
	Name string
	Age  int8
}

// https://juejin.cn/post/7053663185638948895
// https://blog.rexskz.info/use-variant-trie-to-optimize-router-match-with-params.html
// https://segmentfault.com/a/1190000021657573

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	// r := trie.New()
	// r.Use(trie.Logger(), trie.Recovery())
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/hello", func(c *trie.Context) {
	// 		fmt.Println("hello world")
	// 	})
	// 	v2 := v1.Group("/v2")
	// 	v2.Use(trie.OnlyForV2())
	// 	v2.GET("/test", func(c *trie.Context) {
	// 		fmt.Println("hello world")
	// 	})
	// }

	// // r.Static("/assets", "./static")
	// r.SetFuncMap(template.FuncMap{
	// 	"FormatAsDate": FormatAsDate,
	// })
	// r.LoadHTMLGlob("templates/*")
	// stu1 := &student{Name: "Geektutu", Age: 20}
	// stu2 := &student{Name: "Jack", Age: 22}
	// r.GET("/students", func(c *trie.Context) {
	// 	c.HTML(http.StatusOK, "arr.tmpl", trie.H{
	// 		"title":  "gee",
	// 		"stuArr": [2]*student{stu1, stu2},
	// 	})
	// })

	// r.GET("/panic", func(c *trie.Context) {
	// 	names := []string{"geektutu"}
	// 	c.String(http.StatusOK, names[100])
	// })

	// r.Run(":9999")

	// cacheServer()

	mutilPeerServer()

}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func cacheServer() {
	createGroup()

	addr := ":9998"
	peers := gcache.NewHTTPPool(addr)

	log.Println("gcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}

func mutilPeerServer() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Gcache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	// go startAPIServer(apiAddr, gee)
	startCacheServer(addrMap[port], addrs, gee)
}

func createGroup() *gcache.Group {
	return gcache.NewGroup("scores", 2<<10, gcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, gee *gcache.Group) {
	peers := gcache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("gcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, gee *gcache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}
