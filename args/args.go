package args

import (
	"errors"
	"flag"
	"fmt"
)

type (
	arguments struct {
		FileName    string
		PacketSize  int
		PacketizeAt string
		Scheme      string
	}
)

const (
	prePacketization  = "pre"
	postPacketization = "post"

	// UnknownScheme represents unknown coding scheme
	UnknownScheme = 0
	// SchemeGCContent represents GC content coding scheme
	SchemeGCContent = 1
)

var (
	args arguments
)

func init() {
	flag.StringVar(&args.FileName, "f", "", "specify the name of file")
	flag.IntVar(&args.PacketSize, "size", 32, "specify the payload size of packet")
	flag.StringVar(&args.PacketizeAt, "p", "pre", "specify when packetizing")
	flag.StringVar(&args.Scheme, "s", "", "specify the scheme")
}

// Parse arguments
func Parse() error {
	flag.Parse()
	return validate()
}

// Usage arguments
func Usage() {
	flag.Usage()
}

func Debug() {
	fmt.Println(args)
}

func validate() error {
	switch args.PacketizeAt {
	case prePacketization:
	case postPacketization:
	default:
		return errors.New("packetizaAt must be specified")
	}

	if CodingScheme() == UnknownScheme {
		return errors.New("unknown coding scheme")
	}

	return nil
}

// FileName returns file name
func FileName() string {
	return args.FileName
}

// PacketSize returns packet size
func PacketSize() int {
	return args.PacketSize
}

// PacketizeAt returns when packetizing
func PacketizeAt() string {
	return args.PacketizeAt
}

// IsPrePacketization returns whether prePacketization or not
func IsPrePacketization() bool {
	return args.PacketizeAt == prePacketization
}

// IsPostPacketization returns whether postPacketization or not
func IsPostPacketization() bool {
	return args.PacketizeAt == postPacketization
}

// CodingScheme returns DNA coding scheme
func CodingScheme() int {
	switch args.Scheme {
	case "GC":
		return SchemeGCContent
	default:
		return UnknownScheme
	}
}
