package ataraxia

// Group within the network
type Group interface {
	// Name of the group, can be used for debugging
	Name() string

	// Join this group
	Join() error

	// Leave this group
	Leave() error

	// Broadcast a message to all nodes in this group
	Broadcast(messageType string, data interface{}) error

	Messages() chan *Message

	NotifyNodeAvailable(c chan *Node) chan *Node

	NotifyNodeUnavailable(c chan *Node) chan *Node
}
