// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"errors"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/event"
	"github.com/luids-io/api/event/grpc/pb"
)

// SaveEventRequest returns a new protobuf request from event
func SaveEventRequest(e event.Event) (*pb.SaveEventRequest, error) {
	var err error
	req := &pb.SaveEventRequest{}
	req.Id = e.ID
	req.Type = pb.EventType(e.Type)
	req.Code = int32(e.Code)
	req.Level = pb.EventLevel(e.Level)
	req.CreatedTs, _ = ptypes.TimestampProto(e.Created)
	req.Source = SourcePB(e.Source)
	if len(e.Processors) > 0 {
		req.Processors = make([]*pb.ProcessInfo, 0, len(e.Processors))
		for _, p := range e.Processors {
			req.Processors = append(req.Processors, ProcessInfoPB(p))
		}
	}
	req.Data, err = EventDataPB(e.Data)
	if err != nil {
		return nil, err
	}
	req.Codename = e.Codename
	req.Description = e.Description
	if len(e.Tags) > 0 {
		req.Tags = make([]string, len(e.Tags), len(e.Tags))
		copy(req.Tags, e.Tags)
	}
	return req, nil
}

// FromSaveEventRequest returns event
func FromSaveEventRequest(req *pb.SaveEventRequest) (event.Event, error) {
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
	e.Codename = req.GetCodename()
	e.Description = req.GetDescription()
	tags := req.GetTags()
	if len(tags) > 0 {
		e.Tags = make([]string, len(tags), len(tags))
		copy(e.Tags, tags)
	}
	return e, nil
}
