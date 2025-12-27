package p2p

type ShakeHandFunc func(Peer) error

func NOPShakeHandFunc(Peer) error { return nil }
