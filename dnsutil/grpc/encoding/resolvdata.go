// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package encoding

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
)

// ResolvData copy info from pb
func ResolvData(src *pb.ResolvData, dst *dnsutil.ResolvData) (err error) {
	if src.Id != "" {
		dst.ID, err = uuid.Parse(src.Id)
		if err != nil {
			err = fmt.Errorf("bad id format: %v", err)
			return
		}
	}
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
	if srcQueryFlags != nil {
		dst.QueryFlags.Do = srcQueryFlags.GetDo()
		dst.QueryFlags.AuthenticatedData = srcQueryFlags.GetAuthenticatedData()
		dst.QueryFlags.CheckingDisabled = srcQueryFlags.GetCheckingDisabled()
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
	if srcResponseFlags != nil {
		dst.ResponseFlags.AuthenticatedData = srcResponseFlags.GetAuthenticatedData()
	}
	dst.TLD = src.GetTld()
	dst.TLDPlusOne = src.GetTldPlusOne()
	return
}

// ResolvDataPB returns a new protobuf request from data
func ResolvDataPB(src *dnsutil.ResolvData, dst *pb.ResolvData) (err error) {
	tstamp, _ := ptypes.TimestampProto(src.Timestamp)
	dst.Id = src.ID.String()
	dst.Ts = tstamp
	dst.Duration = int64(src.Duration)
	dst.ServerIp = src.Server.String()
	dst.ClientIp = src.Client.String()
	dst.Qid = int32(src.QID)
	dst.Name = src.Name
	dst.IsIpv6 = src.IsIPv6
	dst.ReturnCode = int32(src.ReturnCode)
	dst.Tld = src.TLD
	dst.TldPlusOne = src.TLDPlusOne
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
	return
}
