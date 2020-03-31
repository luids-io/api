package encoding

import (
	"github.com/golang/protobuf/ptypes"

	pb "github.com/luids-io/api/protogen/tlsutilpb"
	"github.com/luids-io/core/tlsutil"
	"github.com/luids-io/core/tlsutil/layer"
)

// RecordData convert to model RecordData
func RecordData(src *pb.RecordData) *tlsutil.RecordData {
	dst := &tlsutil.RecordData{}
	dst.StreamID = src.GetStreamId()
	dst.Type = layer.ContentType(src.GetType())
	dst.Len = uint16(src.GetLen())
	dst.Timestamp, _ = ptypes.Timestamp(src.GetTimestamp())
	dst.Ciphered = src.GetCiphered()
	if dst.Ciphered {
		dst.Fragmented = src.GetFragmented()
		dst.NumMsg = int(src.GetMsgsCount())
	}
	return dst
}

// RecordDataPB convert to protobuf RecordData
func RecordDataPB(src *tlsutil.RecordData) *pb.RecordData {
	dst := &pb.RecordData{}
	dst.StreamId = src.StreamID
	dst.Type = int32(src.Type)
	dst.Len = int32(src.Len)
	dst.Timestamp, _ = ptypes.TimestampProto(src.Timestamp)
	dst.Ciphered = src.Ciphered
	if !src.Ciphered {
		dst.Fragmented = src.Fragmented
		dst.MsgsCount = int32(src.NumMsg)
	}
	return dst
}
