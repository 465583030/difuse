// automatically generated by the FlatBuffers compiler, do not modify

package gentypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RequestOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsRequestOptions(buf []byte, offset flatbuffers.UOffsetT) *RequestOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RequestOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RequestOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RequestOptions) Consistency() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RequestOptions) MutateConsistency(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func RequestOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RequestOptionsAddConsistency(builder *flatbuffers.Builder, Consistency int8) {
	builder.PrependInt8Slot(0, Consistency, 0)
}
func RequestOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
