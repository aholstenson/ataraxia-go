package ataraxia

type Network struct {
}

type Config struct{}

type Option func(*Config) error

func New() (*Network, error) {
	return nil, nil
}

func (n *Network) Join() error {
	return nil
}

func (n *Network) Leave() error {
	return nil
}

func (n *Network) NotifyNodeAvailable(c chan *Node) chan *Node {
	return c
}

func (n *Network) NotifyNodeUnavailable(c chan *Node) chan *Node {
	return c
}
