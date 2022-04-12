package cache

import pb "trie-demo/cache/proto"

type PeerPicker interface {
	PeerPicker(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
