// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
)

// ResolvData returns ResolvData from request
func ResolvData(src *pb.ResolvData) (dst dnsutil.ResolvData, err error) {
	dst.ID = src.Id
	dst.Timestamp, _ = ptypes.Timestamp(src.GetTs())
	dst.Duration = time.Duration(src.GetDuration())
	dst.Server = net.ParseIP(src.GetServerIp())
	if dst.Server == nil {
		err = fmt.Errorf("bad server ip: %v", src.ServerIp)
		return
	}
	dst.Client = net.ParseIP(src.GetClientIp())
	if dst.Client == nil {
		err = errors.New("bad client ip")
		return
	}
	// get query info
	dst.QID = uint16(src.GetQid())
	dst.Name = src.GetName()
	dst.IsIPv6 = src.GetIsIpv6()
	srcQueryFlags := src.GetQueryFlags()
	if srcQueryFlags == nil {
		err = errors.New("query flags is empty")
		return
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
				err = errors.New("bad resolved ip")
				return
			}
			dst.ResolvedIPs = append(dst.ResolvedIPs, ip)
		}
	}
	srcResolvedCNAMEs := src.GetResolvedCnames()
	if len(srcResolvedCNAMEs) > 0 {
		dst.ResolvedCNAMEs = make([]string, 0, len(srcResolvedCNAMEs))
		for _, r := range srcResolvedCNAMEs {
			dst.ResolvedCNAMEs = append(dst.ResolvedCNAMEs, r)
		}
	}
	if srcResponseFlags == nil {
		err = errors.New("response flags is empty")
		return
	}
	dst.ResponseFlags = dnsutil.ResolvResponseFlags{
		AuthenticatedData: srcResponseFlags.GetAuthenticatedData(),
	}
	dst.TLD = src.GetTld()
	dst.TLDPlusOne = src.GetTldPlusOne()
	return
}

// ResolvDataPB returns a new protobuf request from data
func ResolvDataPB(src *dnsutil.ResolvData) (*pb.ResolvData, error) {
	tstamp, _ := ptypes.TimestampProto(src.Timestamp)
	dst := &pb.ResolvData{
		Id:         src.ID,
		Ts:         tstamp,
		Duration:   int64(src.Duration),
		ServerIp:   src.Server.String(),
		ClientIp:   src.Client.String(),
		Qid:        int32(src.QID),
		Name:       src.Name,
		IsIpv6:     src.IsIPv6,
		ReturnCode: int32(src.ReturnCode),
		Tld:        src.TLD,
		TldPlusOne: src.TLDPlusOne,
	}
	dst.QueryFlags = &pb.ResolvData_QueryFlags{
		Do:                src.QueryFlags.Do,
		AuthenticatedData: src.QueryFlags.AuthenticatedData,
		CheckingDisabled:  src.QueryFlags.CheckingDisabled,
	}
	dst.ResponseFlags = &pb.ResolvData_ResponseFlags{
		AuthenticatedData: src.ResponseFlags.AuthenticatedData,
	}
	if len(src.ResolvedIPs) > 0 {
		dst.ResolvedIps = make([]string, 0, len(src.ResolvedIPs))
		for _, r := range src.ResolvedIPs {
			dst.ResolvedIps = append(dst.ResolvedIps, r.String())
		}
	}
	if len(src.ResolvedCNAMEs) > 0 {
		dst.ResolvedCnames = make([]string, 0, len(src.ResolvedCNAMEs))
		for _, r := range src.ResolvedCNAMEs {
			dst.ResolvedCnames = append(dst.ResolvedCnames, r)
		}
	}
	return dst, nil
}
