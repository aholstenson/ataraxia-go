package ataraxia

type Node struct {
}

// Messages returns a channel that can be used to listen to all incoming
// data from this specific node
func (n *Node) Messages() chan Message {
	return nil
}

// Send a message of the given type to this node
func (n *Node) Send(messageType string, data interface{}) error {
	return nil
}

// NotifyUnavailable takes a channel that will be notified when this node is
// no longer available in the network. When this occurs this node should be
// discarded - it will not be reused
func (n *Node) NotifyUnavailable(c chan *Node) chan *Node {
	// TODO
	return c
}
