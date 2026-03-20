package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/kamizox/erofs-go"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <erofs-image>\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	sb, err := erofs.ParseSuperblock(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: parse superblock: %v\n", err)
		os.Exit(1)
	}

	blockSize := uint64(1) << sb.BlockSizeBits
	volName := strings.TrimRight(string(sb.VolumeName[:]), "\x00")
	uuidStr := formatUUID(sb.UUID)

	fmt.Printf("Magic:           0x%X\n", sb.Magic)
	fmt.Printf("Block size:      %d bytes\n", blockSize)
	fmt.Printf("Total blocks:    %d\n", sb.Blocks)
	fmt.Printf("Total inodes:    %d\n", sb.Inos)
	fmt.Printf("Volume name:     %s\n", volName)
	fmt.Printf("UUID:            %s\n", uuidStr)
}

// formatUUID renders a 16-byte UUID as the standard hyphenated hex string.
func formatUUID(b [16]byte) string {
	s := hex.EncodeToString(b[:])
	if len(s) != 32 {
		return s
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s", s[0:8], s[8:12], s[12:16], s[16:20], s[20:32])
}
