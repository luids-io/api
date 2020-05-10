package encoding

import (
	"errors"
	"net"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/pb"
)

// MessageRequestPB serialize to PB
func MessageRequestPB(src *tlsutil.Msg) *pb.SendMessageRequest {
	dst := &pb.SendMessageRequest{}
	dst.Type = pb.MsgType(src.Type)
	dst.StreamId = src.StreamID
	if src.Type == tlsutil.OpenMsg && src.Open != nil {
		dst.Open = &pb.MsgOpen{}
		dst.Open.SourceIp = src.Open.SrcIP.String()
		dst.Open.SourcePort = uint32(src.Open.SrcPort)
		dst.Open.DestinationIp = src.Open.DstIP.String()
		dst.Open.DestinationPort = uint32(src.Open.DstPort)
		return dst
	}
	if src.Type == tlsutil.DataMsg && src.Data != nil {
		dst.Data = &pb.MsgData{}
		dst.Data.Timestamp, _ = ptypes.TimestampProto(src.Data.Timestamp)
		dst.Data.Bytes = int32(src.Data.Bytes)
		dst.Data.SawStart = src.Data.SawStart
		dst.Data.SawEnd = src.Data.SawEnd
		if len(src.Data.Records) > 0 {
			dst.Data.Records = make([][]byte, 0, len(src.Data.Records))
			for _, records := range src.Data.Records {
				dst.Data.Records = append(dst.Data.Records, records)
			}
		}
		if src.Data.Error != nil {
			dst.Data.MsgErr = src.Data.Error.Error()
		}
		return dst
	}
	return dst
}

// MessageRequest unserialize  PB
func MessageRequest(src *pb.SendMessageRequest) *tlsutil.Msg {
	dst := &tlsutil.Msg{}
	dst.Type = tlsutil.MsgType(src.GetType())
	dst.StreamID = src.GetStreamId()
	if dst.Type == tlsutil.OpenMsg && src.GetOpen() != nil {
		openpb := src.GetOpen()
		dst.Open = &tlsutil.MsgOpen{}
		dst.Open.SrcIP = net.ParseIP(openpb.GetSourceIp())
		dst.Open.DstIP = net.ParseIP(openpb.GetDestinationIp())
		dst.Open.SrcPort = int(openpb.GetSourcePort())
		dst.Open.DstPort = int(openpb.GetDestinationPort())
		return dst
	}
	if dst.Type == tlsutil.DataMsg && src.GetData() != nil {
		datapb := src.GetData()
		dst.Data = &tlsutil.MsgData{}
		dst.Data.Timestamp, _ = ptypes.Timestamp(datapb.GetTimestamp())
		dst.Data.Bytes = int(datapb.GetBytes())
		dst.Data.SawStart = datapb.GetSawStart()
		dst.Data.SawEnd = datapb.GetSawEnd()
		if len(datapb.GetRecords()) > 0 {
			pbrecords := datapb.GetRecords()
			dst.Data.Records = make([][]byte, 0, len(pbrecords))
			for _, records := range pbrecords {
				dst.Data.Records = append(dst.Data.Records, records)
			}
		}
		if datapb.GetMsgErr() != "" {
			dst.Data.Error = errors.New(datapb.GetMsgErr())
		}
		return dst
	}
	return dst
}
