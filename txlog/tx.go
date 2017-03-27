package txlog

import (
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/btcsuite/fastsha256"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/ipkg/difuse/gentypes"
	"github.com/ipkg/difuse/keypairs"
	"github.com/ipkg/difuse/utils"
)

// Signator is used to sign a transaction
type Signator interface {
	Sign([]byte) (*keypairs.Signature, error)
	PublicKey() keypairs.PublicKey
	Verify(pubkey, signature, hash []byte) error
}

// TxHeader contains header info for a transaction.
type TxHeader struct {
	PrevHash    []byte
	Source      []byte // from pubkey
	Destination []byte // to pubkey
	Timestamp   uint64 // Timestamp when tx was created
	Height      uint64 // Index from first tx in the chain for a key.
}

// Tx represents a single transaction
type Tx struct {
	*TxHeader
	Signature []byte
	Key       []byte
	Data      []byte
}

// NewTx given the previous tx hash, data and optional public keys
func NewTx(key, prevHash, data []byte) *Tx {
	tx := &Tx{
		Key: key,
		TxHeader: &TxHeader{
			PrevHash:  prevHash,
			Timestamp: uint64(time.Now().UnixNano()),
		},
		Data: data,
	}
	if utils.IsZeroHash(prevHash) {
		tx.TxHeader.Height = 1
	}
	return tx
}

// DataHash of the tx data
func (tx *Tx) DataHash() []byte {
	s := fastsha256.Sum256(tx.Data)
	return s[:]
}

// bytesToGenHash returns the byte slice that should be used to generate the hash
func (tx *Tx) bytesToGenHash() []byte {
	// key + data hash + previous hash + src pub key + dst pub key
	return concat(tx.Key, tx.DataHash(), tx.PrevHash, tx.Source, tx.Destination)
}

// Hash of the whole Tx
func (tx *Tx) Hash() []byte {
	d := tx.bytesToGenHash()
	s := fastsha256.Sum256(d)
	return s[:]
}

// Sign transaction
func (tx *Tx) Sign(signer Signator) error {
	tx.Source = signer.PublicKey().Bytes()

	sig, err := signer.Sign(tx.Hash())
	if err == nil {
		tx.Signature = sig.Bytes()
	}

	return err
}

// VerifySignature of the transaction
func (tx *Tx) VerifySignature(verifier Signator) error {
	return verifier.Verify(tx.Source, tx.Signature, tx.Hash())
}

// MarshalJSON is a custom JSON marshaller.  It properly formats the hashes
// and includes everything except the transaction data.
func (tx *Tx) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"prev":      hex.EncodeToString(tx.PrevHash),
		"id":        hex.EncodeToString(tx.Hash()),
		"key":       string(tx.Key),
		"timestamp": tx.Timestamp,
		"height":    tx.Height,
	})
}

// Serialize serializes the transaction into the given flatbuffer.
func (tx *Tx) Serialize(fb *flatbuffers.Builder) flatbuffers.UOffsetT {
	kp := fb.CreateByteString(tx.Key)
	pp := fb.CreateByteString(tx.PrevHash)
	sp := fb.CreateByteString(tx.Source)
	dp := fb.CreateByteString(tx.Destination)
	ddp := fb.CreateByteString(tx.Data)
	ssp := fb.CreateByteString(tx.Signature)

	gentypes.TxStart(fb)
	gentypes.TxAddKey(fb, kp)
	gentypes.TxAddPrevHash(fb, pp)
	gentypes.TxAddSource(fb, sp)
	gentypes.TxAddDestination(fb, dp)
	gentypes.TxAddData(fb, ddp)
	gentypes.TxAddSignature(fb, ssp)
	gentypes.TxAddTimestamp(fb, tx.Timestamp)
	gentypes.TxAddHeight(fb, tx.Height)
	return gentypes.TxEnd(fb)
}

// Deserialize deserializes the given flatbuffer object into the transaction struct.
func (tx *Tx) Deserialize(obj *gentypes.Tx) {
	tx.TxHeader = &TxHeader{
		Timestamp:   obj.Timestamp(),
		PrevHash:    obj.PrevHashBytes(),
		Source:      obj.SourceBytes(),
		Destination: obj.DestinationBytes(),
		Height:      obj.Height(),
	}

	tx.Key = obj.KeyBytes()
	tx.Data = obj.DataBytes()
	tx.Signature = obj.SignatureBytes()
}
