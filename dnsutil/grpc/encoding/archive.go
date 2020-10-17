// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"errors"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
)

// SaveResolvRequest returns a new protobuf request from data
func SaveResolvRequest(src *dnsutil.ResolvData) (*pb.SaveResolvRequest, error) {
	tstamp, _ := ptypes.TimestampProto(src.Timestamp)
	dst := &pb.SaveResolvRequest{
		Ts:         tstamp,
		Duration:   int64(src.Duration),
		ServerIp:   src.Server.String(),
		ClientIp:   src.Client.String(),
		Qid:        int32(src.QID),
		Name:       src.Name,
		IsIpv6:     src.IsIPv6,
		ReturnCode: int32(src.ReturnCode),
	}
	dst.QueryFlags = &pb.SaveResolvRequest_QueryFlags{
		Do:                src.QueryFlags.Do,
		AuthenticatedData: src.QueryFlags.AuthenticatedData,
		CheckingDisabled:  src.QueryFlags.CheckingDisabled,
	}
	dst.ResponseFlags = &pb.SaveResolvRequest_ResponseFlags{
		AuthenticatedData: src.ResponseFlags.AuthenticatedData,
	}
	if len(src.ResolvedIPs) > 0 {
		dst.ResolvedIps = make([]string, 0, len(src.ResolvedIPs))
		for _, r := range src.ResolvedIPs {
			dst.ResolvedIps = append(dst.ResolvedIps, r.String())
		}
	}
	return dst, nil
}

// FromSaveResolvRequest returns ResolvData from request
func FromSaveResolvRequest(src *pb.SaveResolvRequest) (*dnsutil.ResolvData, error) {
	dst := &dnsutil.ResolvData{}
	dst.Timestamp, _ = ptypes.Timestamp(src.GetTs())
	dst.Duration = time.Duration(src.GetDuration())
	dst.Server = net.ParseIP(src.GetServerIp())
	if dst.Server == nil {
		return dst, errors.New("bad server ip")
	}
	dst.Client = net.ParseIP(src.GetClientIp())
	if dst.Client == nil {
		return dst, errors.New("bad client ip")
	}
	// get query info
	dst.QID = uint16(src.GetQid())
	dst.Name = src.GetName()
	dst.IsIPv6 = src.GetIsIpv6()
	srcQueryFlags := src.GetQueryFlags()
	if srcQueryFlags == nil {
		return nil, errors.New("query flags is empty")
	}
	dst.QueryFlags = dnsutil.ResolvQueryFlags{
		Do:                srcQueryFlags.GetDo(),
		AuthenticatedData: srcQueryFlags.GetAuthenticatedData(),
		CheckingDisabled:  srcQueryFlags.GetCheckingDisabled(),
	}
	// get response info
	dst.ReturnCode = int(src.GetReturnCode())
	srcResponseFlags := src.GetResponseFlags()
	srcResolvedIPs := src.GetResolvedIps()
	if len(srcResolvedIPs) > 0 {
		dst.ResolvedIPs = make([]net.IP, 0, len(srcResolvedIPs))
		for _, r := range srcResolvedIPs {
			ip := net.ParseIP(r)
			if ip == nil {
				return dst, errors.New("bad resolved ip")
			}
			dst.ResolvedIPs = append(dst.ResolvedIPs, ip)
		}
	}
	if srcResponseFlags == nil {
		return nil, errors.New("response flags is empty")
	}
	dst.ResponseFlags = dnsutil.ResolvResponseFlags{
		AuthenticatedData: srcResponseFlags.GetAuthenticatedData(),
	}
	return dst, nil
}
