package encoding

import (
	"errors"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/luids-io/api/protogen/eventpb"
	"github.com/luids-io/core/event"
)

// NotifyEventRequest returns a new protobuf request from event
func NotifyEventRequest(e event.Event) (*pb.NotifyEventRequest, error) {
	var err error
	req := &pb.NotifyEventRequest{}
	req.Code = int32(e.Code)
	req.Level = pb.EventLevel(e.Level)
	req.CreatedTs, _ = ptypes.TimestampProto(e.Created)
	req.Source = SourcePB(e.Source)
	req.Data, err = EventDataPB(e.Data)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FromNotifyEventRequest returns event
func FromNotifyEventRequest(req *pb.NotifyEventRequest) (event.Event, error) {
	var err error
	e := event.Event{}
	e.Code = event.Code(req.GetCode())
	e.Level = event.Level(req.GetLevel())
	e.Created, _ = ptypes.Timestamp(req.GetCreatedTs())
	// get source
	pbsource := req.GetSource()
	if pbsource == nil {
		return event.Event{}, errors.New("source is empty")
	}
	e.Source = Source(pbsource)
	//decode event data
	pbdata := req.GetData()
	if pbdata == nil {
		return event.Event{}, errors.New("data is empty")
	}
	e.Data, err = EventData(pbdata)
	if err != nil {
		return event.Event{}, err
	}
	return e, nil
}
