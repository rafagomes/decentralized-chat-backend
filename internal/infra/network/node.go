package network

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"log"
)

type Node struct {
	Host host.Host
	DHT  *dht.IpfsDHT
}

// NewNode creates a new libp2p node with a DHT
func (n *Node) NewNode(ctx context.Context) (*Node, error) {
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		libp2p.NATPortMap(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	// Set up the DHT
	kadDHT, err := dht.New(ctx, host)
	if err != nil {
		return nil, fmt.Errorf("failed to create DHT: %w", err)
	}

	// Bootstrap the DHT
	if err := kadDHT.Bootstrap(ctx); err != nil {
		return nil, fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	//	TODO: peer discovery, connection manager, etc.

	return &Node{
		Host: host,
		DHT:  kadDHT,
	}, nil
}

// Start starts the node and prints the ID and number of peers in the DHT
func (n *Node) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			peers := n.DHT.RoutingTable().ListPeers()
			log.Printf("Peers in the DHT: %d\n", len(peers))
			log.Printf("Node started with ID %s\n", n.Host.ID())
		}
	}
}

// ConnectionToPeer connects to a peer using a multiaddress
func (n *Node) ConnectionToPeer(addr string) error {
	ma, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		return fmt.Errorf("failed to parse multiaddr: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(ma)
	if err != nil {
		return fmt.Errorf("failed to parse peer address: %w", err)
	}

	if err := n.Host.Connect(context.Background(), *peerInfo); err != nil {
		return fmt.Errorf("failed to connect to peer: %w", err)
	}

	log.Printf("Connected to peer %s\n", peerInfo.ID)
	return nil
}

// SendMessage sends a message to a peer
func (n *Node) SendMessage(peerID peer.ID, msg []byte) error {
	stream, err := n.Host.NewStream(context.Background(), peerID, "/chat/1.0.0")
	if err != nil {
		return fmt.Errorf("failed to create stream: %w", err)
	}

	if _, err := stream.Write(msg); err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	return nil
}

// ReceiveMessage listens for incoming messages on the /chat/1.0.0 stream
func (n *Node) ReceiveMessage() {
	n.Host.SetStreamHandler("/chat/1.0.0", func(stream network.Stream) {
		buf := make([]byte, 1024)
		n, err := stream.Read(buf)
		if err != nil {
			log.Printf("failed to read message: %s\n", err)
			return
		}

		log.Printf("received message: %s\n", string(buf[:n]))
	})
}

// PutValue puts a value into the DHT
func (n *Node) PutValue(ctx context.Context, key string, value []byte) error {
	return n.DHT.PutValue(ctx, key, value)
}

// GetValue retrieves a value from the DHT
func (n *Node) GetValue(ctx context.Context, key string) ([]byte, error) {
	return n.DHT.GetValue(ctx, key)
}

func (n *Node) ConnectToPeers(ctx context.Context, peerID peer.ID) error {
	peerInfo, err := n.DHT.FindPeer(ctx, peerID)
	if err = n.Host.Connect(ctx, peerInfo); err != nil {
		return fmt.Errorf("failed to connect to peer: %w", err)
	}

	return nil
}
