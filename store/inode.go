package store

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/fastsha256"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/ipkg/difuse/fbtypes"
)

// Inode represents a single unit of data around the ring.
type Inode struct {
	Id []byte
	// total size of the data
	Size int64
	// Inline data indicates that the data in Blocks is the actual
	// data rather than addresses to the data.
	Inline bool
	// This holds the address to physical data.  The address can be of any type
	// i.e. hash, key, url etc..
	Blocks [][]byte
}

func NewInode(id []byte) *Inode {
	return &Inode{
		Id:     id,
		Blocks: make([][]byte, 0),
	}
}

// NewInodeFromData instantiates a new inode setting the hash from the provided data
// to blocks
func NewInodeFromData(key, data []byte) *Inode {
	rk := &Inode{Id: key}
	rk.Size = int64(len(data))
	sh := fastsha256.Sum256(data)
	rk.Blocks = [][]byte{sh[:]}
	return rk
}

// MarshalJSON is for user legibility
func (r *Inode) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"size":   r.Size,
		"key":    string(r.Id),
		"inline": r.Inline,
	}
	bhs := make([]string, len(r.Blocks))
	for i, v := range r.Blocks {
		bhs[i] = fmt.Sprintf("%x", v)
	}
	m["blocks"] = bhs

	return json.Marshal(m)
}

// Deserialize deserializes the buffer into a Inode
func (r *Inode) Deserialize(buf []byte) {
	//var rk RingKey
	fbk := fbtypes.GetRootAsInode(buf, 0)
	r.Id = fbk.IdBytes()
	r.Size = fbk.Size()
	r.Inline = (fbk.Inline() == byte(1))

	l := fbk.BlocksLength()
	r.Blocks = make([][]byte, l)

	for i := 0; i < l; i++ {
		var hj fbtypes.ByteSlice
		fbk.Blocks(&hj, i)
		r.Blocks[i] = hj.BBytes()
	}
}

// Serialize serializes the struct into the flatbuffer returning the offset.
func (r *Inode) Serialize(fb *flatbuffers.Builder) flatbuffers.UOffsetT {
	obh := make([]flatbuffers.UOffsetT, len(r.Blocks))

	for i, v := range r.Blocks {
		bhp := fb.CreateByteString(v)
		fbtypes.ByteSliceStart(fb)
		fbtypes.ByteSliceAddB(fb, bhp)
		obh[i] = fbtypes.ByteSliceEnd(fb)
	}

	fbtypes.InodeStartBlocksVector(fb, len(r.Blocks))
	for _, v := range obh {
		fb.PrependUOffsetT(v)
	}
	bh := fb.EndVector(len(r.Blocks))
	kp := fb.CreateByteString(r.Id)

	fbtypes.InodeStart(fb)
	fbtypes.InodeAddId(fb, kp)
	fbtypes.InodeAddBlocks(fb, bh)
	fbtypes.InodeAddSize(fb, r.Size)
	if r.Inline {
		fbtypes.InodeAddInline(fb, byte(1))
	}
	return fbtypes.InodeEnd(fb)
}