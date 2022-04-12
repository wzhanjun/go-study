package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"trie-demo/cache/consistent"
	pb "trie-demo/cache/proto"

	"github.com/golang/protobuf/proto"
)

const (
	defaultBasePath = "/_gcache/"
	defaultReplicas = 10
)

type HTTPPool struct {
	addr       string
	basePath   string
	lock       sync.Mutex
	peers      *consistent.Map
	httpGetter map[string]*httpGetter
}

func NewHTTPPool(addr string) *HTTPPool {
	return &HTTPPool{
		addr:     addr,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[server %s] %s", p.addr, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexcepted path:" + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)

	// format /<basepath>/<groupname>/<key>
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupname, key := parts[0], parts[1]

	group := GetGroup(groupname)
	if group == nil {
		http.Error(w, "no such group: "+groupname, http.StatusNotFound)
		return
	}
	val, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := proto.Marshal(&pb.Response{Val: val.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/octet-stream")
	w.Write(body)
}

type httpGetter struct {
	baseUrl string
}

func (p *HTTPPool) Set(peers ...string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.peers = consistent.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetter = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetter[peer] = &httpGetter{baseUrl: peer + p.basePath}
	}
}

func (p *HTTPPool) PeerPicker(key string) (PeerGetter, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if peer := p.peers.Get(key); peer != "" && peer != p.addr {
		p.Log("Pick peer %s", peer)
		return p.httpGetter[peer], true
	}
	return nil, false
}

func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf("%v%v%v", h.baseUrl, url.QueryEscape(in.GetGroup()), url.QueryEscape(in.GetKey()))
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returnd：%d", res.StatusCode)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}
	return nil
}

var _ PeerGetter = (*httpGetter)(nil)
var _ PeerPicker = (*HTTPPool)(nil)
