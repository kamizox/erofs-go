package erofs

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	// SuperblockMagic is the fixed EROFS superblock magic value.
	SuperblockMagic uint32 = 0xE0F5E1E2
)

// Superblock represents the essential fields of an EROFS filesystem superblock.
type Superblock struct {
	// Magic identifies this as an EROFS superblock and should equal SuperblockMagic.
	Magic uint32
	// Checksum stores the superblock checksum used for integrity verification.
	Checksum uint32
	// FeatureCompat contains flags for backward-compatible filesystem features.
	FeatureCompat uint32
	// BlockSizeBits stores log2 of the filesystem block size.
	BlockSizeBits uint8
	// ExtSlots is the number of additional slots used by the extended superblock area.
	ExtSlots uint8
	// RootNid is the inode number of the root directory.
	RootNid uint16
	// Inos is the total number of inodes in the filesystem.
	Inos uint64
	// BuildTime is the filesystem build timestamp in seconds since the Unix epoch.
	BuildTime uint64
	// BuildTimeNsec is the nanosecond component of BuildTime.
	BuildTimeNsec uint32
	// Blocks is the total number of filesystem blocks.
	Blocks uint32
	// MetaBlkAddr is the starting block address of metadata area.
	MetaBlkAddr uint32
	// XattrBlkAddr is the starting block address for extended attribute data.
	XattrBlkAddr uint32
	// UUID is the 16-byte filesystem UUID.
	UUID [16]byte
	// VolumeName is the 16-byte volume label.
	VolumeName [16]byte
	// FeatureIncompat contains flags for incompatible filesystem features.
	FeatureIncompat uint32
}

// ParseSuperblock reads and parses EROFS superblock fields from a binary file.
// The fields are decoded using little-endian byte order.
func ParseSuperblock(path string) (*Superblock, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open superblock file: %w", err)
	}
	defer f.Close()

	sb := &Superblock{}
	order := binary.LittleEndian

	fields := []struct {
		name string
		ptr  any
	}{
		{name: "Magic", ptr: &sb.Magic},
		{name: "Checksum", ptr: &sb.Checksum},
		{name: "FeatureCompat", ptr: &sb.FeatureCompat},
		{name: "BlockSizeBits", ptr: &sb.BlockSizeBits},
		{name: "ExtSlots", ptr: &sb.ExtSlots},
		{name: "RootNid", ptr: &sb.RootNid},
		{name: "Inos", ptr: &sb.Inos},
		{name: "BuildTime", ptr: &sb.BuildTime},
		{name: "BuildTimeNsec", ptr: &sb.BuildTimeNsec},
		{name: "Blocks", ptr: &sb.Blocks},
		{name: "MetaBlkAddr", ptr: &sb.MetaBlkAddr},
		{name: "XattrBlkAddr", ptr: &sb.XattrBlkAddr},
		{name: "UUID", ptr: &sb.UUID},
		{name: "VolumeName", ptr: &sb.VolumeName},
		{name: "FeatureIncompat", ptr: &sb.FeatureIncompat},
	}

	for _, field := range fields {
		if err := binary.Read(f, order, field.ptr); err != nil {
			return nil, fmt.Errorf("read %s: %w", field.name, err)
		}
	}

	return sb, nil
}
