// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/event"
	"github.com/luids-io/api/event/grpc/pb"
)

// ProcessInfo returs event.ProcessInfo
func ProcessInfo(src *pb.ProcessInfo) event.ProcessInfo {
	dst := event.ProcessInfo{}
	dst.Received, _ = ptypes.Timestamp(src.GetReceivedTs())
	dst.Processor = Source(src.GetProcessor())
	return dst
}

// ProcessInfoPB returns source protobuf
func ProcessInfoPB(src event.ProcessInfo) *pb.ProcessInfo {
	dst := &pb.ProcessInfo{}
	dst.ReceivedTs, _ = ptypes.TimestampProto(src.Received)
	dst.Processor = SourcePB(src.Processor)
	return dst
}

// Source returs event.Source
func Source(src *pb.EventSource) event.Source {
	dst := event.Source{}
	dst.Hostname = src.GetHostname()
	dst.Instance = src.GetInstance()
	dst.Program = src.GetProgram()
	dst.PID = int(src.GetPid())
	return dst
}

// SourcePB returs source protobuf
func SourcePB(src event.Source) *pb.EventSource {
	dst := &pb.EventSource{}
	dst.Hostname = src.Hostname
	dst.Program = src.Program
	dst.Instance = src.Instance
	dst.Pid = uint32(src.PID)
	return dst
}

// EventData returs map with data
func EventData(src *pb.EventData) (map[string]interface{}, error) {
	var dst map[string]interface{}
	switch src.GetDataEnc() {
	case pb.EventData_JSON:
		rawData := src.GetData()
		err := json.Unmarshal(rawData, &dst)
		if err != nil {
			return dst, fmt.Errorf("unmarshalling data: %v", err)
		}
	case pb.EventData_NODATA:
		dst = make(map[string]interface{})
	default:
		return dst, errors.New("invalid econding data")
	}
	return dst, nil
}

// EventDataPB returns data in pb
func EventDataPB(src map[string]interface{}) (*pb.EventData, error) {
	dst := &pb.EventData{}
	dst.DataEnc = pb.EventData_NODATA
	if len(src) > 0 {
		// encode data to json
		encoded, err := json.Marshal(src)
		if err != nil {
			return nil, fmt.Errorf("cannot encode data to json: %v", err)
		}
		dst.DataEnc = pb.EventData_JSON
		dst.Data = encoded
	}
	return dst, nil
}
