package transport

type NodeId []byte

// ByeMessage indicates that a disconnect should occur
type ByeMessage struct{}

// PingMessage is sent to peers to do a liveness check
type PingMessage struct{}

// PongMessage is the reply sent for PingMessage
type PongMessage struct{}

// OkMessage is sent for certain operations during the handshake
type OkMessage struct{}

// RejectMessage is sent for certain operations during the handshake
type RejectMessage struct{}

// BeginMessage is sent when a connection is negotiated and should be
// established
type BeginMessage struct{}

// HelloMessage is the first message sent when trying to establish a connection
// with another peer
type HelloMessage struct {
	// Identifier of this peer
	Id NodeId

	// Capabilities picked, used to negotiate protocol versions and features
	Capabilities []string
}

// SelectMesage is sent as a reply to a Hello indicating the id of the remote
// peer and the capabilities it picked
type SelectMessage struct {
	// Identifier of this peer
	Id NodeId

	// Capabilities picked from the set sent in Hello
	Capabilities []string
}

// AuthMessage is sent to pick how to authenticate with a remote peer
type AuthMessage struct {
	// Method of authentication picked
	Method string

	// Data associated with the
	Data []byte
}

// AuthDataMessage contains additional data as needed for a specific
// authentication method
type AuthDataMessage struct {
	Data []byte
}

// NodeSummaryMessage contains information about the nodes a peer can see
type NodeSummaryMessage struct {
	// OwnVersion represents the version of this routing. This will increase
	// if detects changes in peers around it
	OwnVersion int32

	// Nodes contains a summary of routing for all the nodes the peer can see
	Nodes []NodeRoutingsummary
}

// NodeRoutingSummary is a summary of a Node id and routing version
type NodeRoutingsummary struct {
	// Id is the identifier of the Node
	Id NodeId

	// Version is the version of the routing, this is used to request more
	// information if the local version differs
	Version int32
}

// NodeRequestMessage is used to request information about nodes from a peer
type NodeRequestMessage struct {
	// Nodes to fetch information for
	Nodes []NodeId
}

// NodeDetailsMessage is the reply to a NodeRequestMessage and contains detailed
// information about the nodes
type NodeDetailsMessage struct {
	// Nodes contains the requested routing details
	Nodes []NodeRoutingDetails
}

// NodeRoutingDetails carries detailed routing information for a specific node
type NodeRoutingDetails struct {
	// Id of this node
	Id NodeId

	// Version of this node
	Version int32

	// Neighbors that this node can see
	Neighbors []NodeWithLatency
}

// NodeWithLatency describes the latency to a neighbor as seen by a certain
// node
type NodeWithLatency struct {
	// Id of the node
	Id NodeId
	// Latency in milliseconds
	Latency int32
}

// DataMessage contains data intended for a certain node
type DataMessage struct {
	// Path describes the current path for the data. Nodes append their
	// id here which is used to make sure send loops are avoided
	Path []DataMessagePathEntry

	// Target is the node that this message should eventually reach
	Target NodeId

	// Type of data
	Type string

	// Data is the binary representation of the data
	Data []byte
}

// DataMessagePathEntry holds information about the a node and id pair
type DataMessagePathEntry struct {
	// Node is the id of the node this passed through
	Node NodeId

	// Id is a local identifier used to keep track of the message
	Id int32
}

// DataAckMessage is used when a node has handled a certain data message,
// either by consuming it or by forwarding it and receiving a reply
type DataAckMessage struct {
	// Id of data acknowledged
	Id int32
}

// DataRejectMessage is used when a node can not handle a certain data message
type DataRejectMessage struct {
	// Id of data rejected
	Id int32
}
