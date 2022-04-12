package cache

import (
	"fmt"
	"log"
	"sync"
	pb "trie-demo/cache/proto"
	"trie-demo/cache/singleflight"
)

type Group struct {
	name   string
	getter Getter
	cache  cache
	peers  PeerPicker
	loader *singleflight.Group
}

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	lock   sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	lock.Lock()
	defer lock.Unlock()

	group := &Group{
		name:   name,
		getter: getter,
		cache:  cache{cacheBytes: cacheBytes},
		loader: &singleflight.Group{},
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	lock.RLock()
	group := groups[name]
	lock.RUnlock()
	return group
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.cache.get(key); ok {
		log.Println("[cache] hit")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (val ByteView, err error) {
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PeerPicker(key); ok {
				if val, err := g.getFromPeer(peer, key); err == nil {
					log.Printf("[gcache] Failed to get from peer:%#v, err:%#v \n", peer, err)
					return val, nil
				}
			}
		}
		return g.getLocal(key)
	})
	if err == nil {
		return viewi.(ByteView), nil
	}
	return
}

func (g *Group) getLocal(key string) (val ByteView, err error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	val = ByteView{b: cloneByte(bytes)}
	g.populateCache(key, val)
	return val, nil
}

func (g *Group) populateCache(key string, val ByteView) {
	g.cache.add(key, val)
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("register peerpick called more than once")
	}
	g.peers = peers
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &pb.Request{
		Group: g.name,
		Key:   key,
	}
	res := &pb.Response{}
	err := peer.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Val}, nil
}
