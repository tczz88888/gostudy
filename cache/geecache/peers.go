package geecache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool) //根据传入的key找到相应的节点
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error) //根据节点和key返回缓存的值
}
