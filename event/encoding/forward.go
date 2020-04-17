package encoding

import (
	"errors"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/luids-io/api/protogen/eventpb"
	"github.com/luids-io/core/event"
)

// ForwardEventRequest returns a new protobuf request from event
func ForwardEventRequest(e event.Event) (*pb.ForwardEventRequest, error) {
	var err error
	req := &pb.ForwardEventRequest{}
	req.Id = e.ID
	req.Type = pb.EventType(e.Type)
	req.Code = int32(e.Code)
	req.Level = pb.EventLevel(e.Level)
	req.CreatedTs, _ = ptypes.TimestampProto(e.Created)
	req.Source = SourcePB(e.Source)
	if len(e.Processors) > 0 {
		req.Processors = make([]*pb.ProcessInfo, 0, len(e.Processors)+1)
		for _, p := range e.Processors {
			req.Processors = append(req.Processors, ProcessInfoPB(p))
		}
	}
	req.Data, err = EventDataPB(e.Data)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FromForwardEventRequest returns event
func FromForwardEventRequest(req *pb.ForwardEventRequest) (event.Event, error) {
	var err error
	e := event.Event{}
	e.ID = req.GetId()
	e.Type = event.Type(req.GetType())
	e.Code = event.Code(req.GetCode())
	e.Level = event.Level(req.GetLevel())
	e.Created, _ = ptypes.Timestamp(req.GetCreatedTs())
	// get source
	pbsource := req.GetSource()
	if pbsource == nil {
		return event.Event{}, errors.New("source is empty")
	}
	e.Source = Source(pbsource)
	// get processors
	pbprocessors := req.GetProcessors()
	if len(pbprocessors) > 0 {
		e.Processors = make([]event.ProcessInfo, 0, len(pbprocessors))
		for _, p := range pbprocessors {
			e.Processors = append(e.Processors, ProcessInfo(p))
		}
	}
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
