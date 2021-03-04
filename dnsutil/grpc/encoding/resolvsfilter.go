// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"net"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
)

// ResolvsFilter copy info from pb
func ResolvsFilter(src *pb.ResolvsFilter, dst *dnsutil.ResolvsFilter) (err error) {
	if src.Since != nil {
		dst.Since, _ = ptypes.Timestamp(src.GetSince())
	}
	if src.To != nil {
		dst.To, _ = ptypes.Timestamp(src.GetTo())
	}
	dst.Server = net.ParseIP(src.GetServerIp())
	dst.Client = net.ParseIP(src.GetClientIp())
	dst.Name = src.GetName()
	dst.ResolvedIP = net.ParseIP(src.GetResolvedIp())
	dst.ResolvedCNAME = src.GetResolvedCname()
	dst.QID = int(src.GetQid())
	dst.ReturnCode = int(src.GetReturnCode())
	dst.TLD = src.GetTld()
	dst.TLDPlusOne = src.GetTldPlusOne()
	return
}

// ResolvsFilterPB copy info to pb
func ResolvsFilterPB(src *dnsutil.ResolvsFilter, dst *pb.ResolvsFilter) (err error) {
	if !src.Since.IsZero() {
		dst.Since, _ = ptypes.TimestampProto(src.Since)
	}
	if !src.To.IsZero() {
		dst.To, _ = ptypes.TimestampProto(src.To)
	}
	if src.Server != nil {
		dst.ServerIp = src.Server.String()
	}
	if src.Client != nil {
		dst.ClientIp = src.Client.String()
	}
	dst.Name = src.Name
	if src.ResolvedIP != nil {
		dst.ResolvedIp = src.ResolvedIP.String()
	}
	dst.ResolvedCname = src.ResolvedCNAME
	dst.Qid = int32(src.QID)
	dst.ReturnCode = int32(src.ReturnCode)
	dst.Tld = src.TLD
	dst.TldPlusOne = src.TLDPlusOne
	return
}
