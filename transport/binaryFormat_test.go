package transport

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecodePing(t *testing.T) {
	initial := &PingMessage{}

	var buf bytes.Buffer
	if err := EncodeBinaryPeerMessage(initial, &buf); err != nil {
		t.Fatal("Could not encode:", err)
	}

	decoded, err := DecodeBinaryPeerMessage(&buf)
	if err != nil {
		t.Fatal("Could not decode:", err)
	}

	if !reflect.DeepEqual(decoded, initial) {
		t.Fatal("not equal")
	}
}

func TestEncodeDecodePong(t *testing.T) {
	initial := &PongMessage{}

	var buf bytes.Buffer
	if err := EncodeBinaryPeerMessage(initial, &buf); err != nil {
		t.Fatal("Could not encode:", err)
	}

	decoded, err := DecodeBinaryPeerMessage(&buf)
	if err != nil {
		t.Fatal("Could not decode:", err)
	}

	if !reflect.DeepEqual(decoded, initial) {
		t.Fatal("not equal")
	}
}

func TestEncodeDecodeHello(t *testing.T) {
	initial := &HelloMessage{
		Id:           []byte("test"),
		Capabilities: []string{"a", "b"},
	}

	var buf bytes.Buffer
	if err := EncodeBinaryPeerMessage(initial, &buf); err != nil {
		t.Fatal("Could not encode:", err)
	}

	decoded, err := DecodeBinaryPeerMessage(&buf)
	if err != nil {
		t.Fatal("Could not decode:", err)
	}

	if !reflect.DeepEqual(decoded, initial) {
		t.Fatal("not equal")
	}
}
