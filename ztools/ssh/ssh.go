package ssh

import (
	"errors"
	"io"
	"net"
)

var errShortPacket = errors.New("SSH packet too short")
var errLongPacket = errors.New("SSH packet too long")
var errInvalidPadding = errors.New("Invalid SSH padding")
var errUnexpectedMessage = errors.New("Unexpected SSH message type")
var errShortBuffer = errors.New("Buffer too short")
var errInvalidPlaintextLength = errors.New("Plaintext not a multiple of block size")

// Client wraps a network connection with an SSH client connection
func Client(c net.Conn, config *Config) *Conn {
	return &Conn{
		conn:   c,
		config: config,
	}
}

// SSH message types. These are usually the first byte of the payload
const (
	SSH_MSG_KEXINIT byte = 0x14
)

type Config struct {
	KexAlgorithms             NameList
	HostKeyAlgorithms         NameList
	EncryptionClientToServer  NameList
	EncryptionServerToClient  NameList
	MACClientToServer         NameList
	MACServerToclient         NameList
	CompressionClientToServer NameList
	CompressionServerToClient NameList
	Random                    io.Reader
}

func (c *Config) getKexAlgorithms() NameList {
	if c.KexAlgorithms != nil {
		return c.KexAlgorithms
	}
	return KnownKexAlgorithmNames
}

func (c *Config) getHostKeyAlgorithms() NameList {
	if c.HostKeyAlgorithms != nil {
		return c.HostKeyAlgorithms
	}
	return KnownHostKeyAlgorithmNames
}

func (c *Config) getClientEncryption() NameList {
	if c.EncryptionClientToServer != nil {
		return c.EncryptionClientToServer
	}
	return KnownEncryptionAlgorithmNames
}

func (c *Config) getServerEncryption() NameList {
	if c.EncryptionServerToClient != nil {
		return c.EncryptionServerToClient
	}
	return c.getClientEncryption()
}

func (c *Config) getClientMAC() NameList {
	if c.MACClientToServer != nil {
		return c.MACClientToServer
	}
	return KnownMACAlgorithmNames
}

func (c *Config) getServerMAC() NameList {
	if c.MACServerToclient != nil {
		return c.MACServerToclient
	}
	return c.getClientMAC()
}

func (c *Config) getClientCompression() NameList {
	if c.CompressionClientToServer != nil {
		return c.CompressionClientToServer
	}
	return KnownCompressionAlgorithmNames
}

func (c *Config) getServerCompression() NameList {
	if c.CompressionServerToClient != nil {
		return c.CompressionServerToClient
	}
	return c.getClientCompression()
}
