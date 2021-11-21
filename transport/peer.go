package transport

type Peer struct {
}

func (p *Peer) Id() []byte {
	return nil
}

func (p *Peer) NotifyConnect(c chan *Peer) chan *Peer {
	return c
}
