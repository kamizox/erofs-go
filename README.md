# erofs-go

`erofs-go` is a Go library for parsing EROFS (Enhanced Read-Only File System) images.
It is focused on providing clear data structures and binary parsing helpers for core
filesystem metadata, starting with the EROFS superblock.

## Current Features

- Parses EROFS superblock magic number
- Calculates block size from `BlockSizeBits`
- Displays total blocks, inodes, volume name and UUID
- Cross-platform: works on Windows, Linux, macOS

## Installation

```bash
go get github.com/kamizox/erofs-go
```

## Usage

The example below shows how to parse a superblock from an EROFS image file.

```go
package main

import (
	"fmt"
	"log"

	"github.com/kamizox/erofs-go"
)

func main() {
	sb, err := erofs.ParseSuperblock("rootfs.erofs")
	if err != nil {
		log.Fatalf("failed to parse superblock: %v", err)
	}

	fmt.Printf("Magic: 0x%X\n", sb.Magic)
	fmt.Printf("Block size bits: %d\n", sb.BlockSizeBits)
	fmt.Printf("Inodes: %d\n", sb.Inos)
	fmt.Printf("Blocks: %d\n", sb.Blocks)
}
```

## Demo Output

Example from running the `cmd/erofs-go` CLI against an image:

```
Magic:        0xE0F5E1E2
Block size:   4096 bytes
Total blocks: 3276800
Total inodes: 28147497671065600
Volume name:  
UUID:         00000000-0000-0000-0000-000000000000
```

## Roadmap

- Image injection for fuzz testing
- Inode table parser
- Directory entry walker
- GitHub Actions CI integration

## GSoC 2026 Alignment

This library is being developed to support the goals of the **GSoC 2026 EROFS project**,
including better tooling for reading, inspecting, and building EROFS filesystem images in Go.
