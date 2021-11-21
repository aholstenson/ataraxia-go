package transport

import (
	"errors"
	"io"

	"github.com/fxamacker/cbor/v2"
)

var ErrUnknownPeerMessage error = errors.New("unknown peer message")

const TAG_PING = 0
const TAG_PONG = 1

const TAG_DATA = 2
const TAG_DATA_ACK = 3
const TAG_DATA_REJECT = 4

const TAG_OK = 10
const TAG_REJECT = 11

const TAG_HELLO = 12
const TAG_SELECT = 13

const TAG_AUTH = 14
const TAG_AUTH_DATA = 15

const TAG_BEGIN = 16
const TAG_BYE = 17

const TAG_NODE_SUMMARY = 18
const TAG_NODE_REQUEST = 19
const TAG_NODE_DETAILS = 20

// EncodeBinaryPeerMessage takes a peer message and turns it into
func EncodeBinaryPeerMessage(data interface{}, target io.Writer) (err error) {
	encoder := cbor.NewEncoder(target)

	switch m := data.(type) {
	case *PingMessage:
		if err = encoder.Encode(TAG_PING); err != nil {
			return
		}
	case *PongMessage:
		if err = encoder.Encode(TAG_PONG); err != nil {
			return
		}
	case *OkMessage:
		if err = encoder.Encode(TAG_OK); err != nil {
			return
		}
	case *RejectMessage:
		if err = encoder.Encode(TAG_REJECT); err != nil {
			return
		}
	case *HelloMessage:
		if err = encoder.Encode(TAG_HELLO); err != nil {
			return
		}
		if err = encoder.Encode(m.Id); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Capabilities)); err != nil {
			return
		}
		for _, cap := range m.Capabilities {
			if err = encoder.Encode(cap); err != nil {
				return
			}
		}
	case *SelectMessage:
		if err = encoder.Encode(TAG_SELECT); err != nil {
			return
		}
		if err = encoder.Encode(m.Id); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Capabilities)); err != nil {
			return
		}
		for _, cap := range m.Capabilities {
			if err = encoder.Encode(cap); err != nil {
				return
			}
		}
	case *AuthMessage:
		if err = encoder.Encode(TAG_AUTH); err != nil {
			return
		}
		if err = encoder.Encode(m.Method); err != nil {
			return
		}
		if err = encoder.Encode(m.Data); err != nil {
			return
		}
	case *AuthDataMessage:
		if err = encoder.Encode(TAG_AUTH_DATA); err != nil {
			return
		}
		if err = encoder.Encode(m.Data); err != nil {
			return
		}
	case *BeginMessage:
		if err = encoder.Encode(TAG_BEGIN); err != nil {
			return
		}
	case *ByeMessage:
		if err = encoder.Encode(TAG_BYE); err != nil {
			return
		}
	case *NodeSummaryMessage:
		if err = encoder.Encode(TAG_NODE_SUMMARY); err != nil {
			return
		}
		if err = encoder.Encode(m.OwnVersion); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Nodes)); err != nil {
			return
		}
		for _, node := range m.Nodes {
			if err = encoder.Encode(node.Id); err != nil {
				return
			}
			if err = encoder.Encode(node.Version); err != nil {
				return
			}
		}
	case *NodeRequestMessage:
		if err = encoder.Encode(TAG_NODE_REQUEST); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Nodes)); err != nil {
			return
		}
		for _, node := range m.Nodes {
			if err = encoder.Encode(node); err != nil {
				return
			}
		}
	case *NodeDetailsMessage:
		if err = encoder.Encode(TAG_NODE_DETAILS); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Nodes)); err != nil {
			return
		}
		for _, node := range m.Nodes {
			if err = encoder.Encode(node.Id); err != nil {
				return
			}
			if err = encoder.Encode(node.Version); err != nil {
				return
			}
			if err = encoder.Encode(len(node.Neighbors)); err != nil {
				return
			}
			for _, neighbor := range node.Neighbors {
				if err = encoder.Encode(neighbor.Id); err != nil {
					return
				}
				if err = encoder.Encode(neighbor.Latency); err != nil {
					return
				}
			}
		}
	case *DataMessage:
		if err = encoder.Encode(TAG_DATA); err != nil {
			return
		}
		if err = encoder.Encode(m.Target); err != nil {
			return
		}
		if err = encoder.Encode(m.Type); err != nil {
			return
		}
		if err = encoder.Encode(m.Data); err != nil {
			return
		}
		if err = encoder.Encode(len(m.Path)); err != nil {
			return
		}
		for _, e := range m.Path {
			if err = encoder.Encode(e.Node); err != nil {
				return
			}
			if err = encoder.Encode(e.Id); err != nil {
				return
			}
		}
	case *DataAckMessage:
		if err = encoder.Encode(TAG_DATA_ACK); err != nil {
			return
		}
		if err = encoder.Encode(m.Id); err != nil {
			return
		}
	case *DataRejectMessage:
		if err = encoder.Encode(TAG_DATA_REJECT); err != nil {
			return
		}
		if err = encoder.Encode(m.Id); err != nil {
			return
		}
	default:
		return ErrUnknownPeerMessage
	}

	return nil
}

func DecodeBinaryPeerMessage(reader io.Reader) (m interface{}, err error) {
	decoder := cbor.NewDecoder(reader)

	var tag int32
	if err = decoder.Decode(&tag); err != nil {
		return
	}

	switch tag {
	case TAG_PING:
		m = &PingMessage{}
	case TAG_PONG:
		m = &PongMessage{}
	case TAG_OK:
		m = &OkMessage{}
	case TAG_REJECT:
		m = &RejectMessage{}
	case TAG_HELLO:
		var id []byte
		if err = decoder.Decode(&id); err != nil {
			return
		}

		caps, err2 := decodePeerStringArray(decoder)
		if err2 != nil {
			return nil, err2
		}

		m = &HelloMessage{
			Id:           id,
			Capabilities: caps,
		}
	case TAG_SELECT:
		var id []byte
		if err = decoder.Decode(&id); err != nil {
			return
		}

		caps, err2 := decodePeerStringArray(decoder)
		if err2 != nil {
			return nil, err2
		}

		m = &SelectMessage{
			Id:           id,
			Capabilities: caps,
		}
	case TAG_AUTH:
		var method string
		if err = decoder.Decode(&method); err != nil {
			return
		}

		var data []byte
		if err = decoder.Decode(&data); err != nil {
			return
		}

		m = &AuthMessage{
			Method: method,
			Data:   data,
		}
	case TAG_AUTH_DATA:
		var data []byte
		if err = decoder.Decode(&data); err != nil {
			return
		}

		m = &AuthDataMessage{
			Data: data,
		}
	case TAG_BEGIN:
		m = &BeginMessage{}
	case TAG_BYE:
		m = &ByeMessage{}
	case TAG_NODE_REQUEST:
		var len int
		if err = decoder.Decode(&len); err != nil {
			return
		}

		ids := make([]NodeId, len)
		for i := 0; i < int(len); i++ {
			if err = decoder.Decode(&ids[i]); err != nil {
				return
			}
		}

		m = &NodeRequestMessage{
			Nodes: ids,
		}
	case TAG_NODE_DETAILS:
		var len int
		if err = decoder.Decode(&len); err != nil {
			return
		}

		nodes := make([]NodeRoutingDetails, len)
		for i := 0; i < int(len); i++ {
			var id NodeId
			if err = decoder.Decode(&id); err != nil {
				return
			}

			var version int32
			if err = decoder.Decode(&version); err != nil {
				return
			}

			var nLen int
			if err = decoder.Decode(&nLen); err != nil {
				return
			}

			neighbors := make([]NodeWithLatency, nLen)

			for j := 0; j < nLen; j++ {
				var nId NodeId
				if err = decoder.Decode(&nId); err != nil {
					return
				}

				var latency int32
				if err = decoder.Decode(&latency); err != nil {
					return
				}

				neighbors[j] = NodeWithLatency{
					Id:      nId,
					Latency: latency,
				}
			}

			nodes[i] = NodeRoutingDetails{
				Id:        id,
				Version:   version,
				Neighbors: neighbors,
			}
		}

		m = &NodeDetailsMessage{
			Nodes: nodes,
		}
	case TAG_DATA:
		var target NodeId
		if err = decoder.Decode(&target); err != nil {
			return
		}

		var messageType string
		if err = decoder.Decode(&messageType); err != nil {
			return
		}

		var data []byte
		if err = decoder.Decode(&data); err != nil {
			return
		}

		var pathLen int
		if err = decoder.Decode(&pathLen); err != nil {
			return
		}

		path := make([]DataMessagePathEntry, pathLen)

		for i := 0; i < pathLen; i++ {
			var pNodeId NodeId
			if err = decoder.Decode(&pNodeId); err != nil {
				return
			}

			var pId int32
			if err = decoder.Decode(&pId); err != nil {
				return
			}

			path[i] = DataMessagePathEntry{
				Node: pNodeId,
				Id:   pId,
			}
		}

		m = &DataMessage{
			Target: target,
			Type:   messageType,
			Path:   path,
			Data:   data,
		}
	case TAG_DATA_ACK:
		var id int32
		if err = decoder.Decode(&id); err != nil {
			return
		}

		m = &DataAckMessage{
			Id: id,
		}
	case TAG_DATA_REJECT:
		var id int32
		if err = decoder.Decode(&id); err != nil {
			return
		}

		m = &DataRejectMessage{
			Id: id,
		}
	default:
		err = ErrUnknownPeerMessage
	}

	return
}

func decodePeerStringArray(d *cbor.Decoder) ([]string, error) {
	var len int
	if err := d.Decode(&len); err != nil {
		return nil, err
	}

	r := make([]string, len)
	for i := 0; i < int(len); i++ {
		if err := d.Decode(&r[i]); err != nil {
			return nil, err
		}
	}

	return r, nil
}
