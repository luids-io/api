// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"errors"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/encoding"
	"github.com/luids-io/api/tlsutil/grpc/pb"
)

func certificateToRequest(cert *tlsutil.CertificateData) (*pb.SaveCertificateRequest, error) {
	req := &pb.SaveCertificateRequest{}
	req.Certificate = encoding.CertificateDataPB(cert)
	return req, nil
}

func certificateFromRequest(req *pb.SaveCertificateRequest) (*tlsutil.CertificateData, error) {
	c := req.GetCertificate()
	if c == nil {
		return nil, errors.New("certificate data is empty")
	}
	return encoding.CertificateData(c), nil
}

func connectionToRequest(cn *tlsutil.ConnectionData) (*pb.SaveConnectionRequest, error) {
	req := &pb.SaveConnectionRequest{}
	req.Connection = encoding.ConnectionDataPB(cn)
	return req, nil
}

func connectionFromRequest(req *pb.SaveConnectionRequest) (*tlsutil.ConnectionData, error) {
	c := req.GetConnection()
	if c == nil {
		return nil, errors.New("connection data is empty")
	}
	return encoding.ConnectionData(c), nil
}

func recordToRequest(r *tlsutil.RecordData) *pb.SaveRecordRequest {
	req := &pb.SaveRecordRequest{}
	req.Record = encoding.RecordDataPB(r)
	return req
}

func recordFromRequest(req *pb.SaveRecordRequest) (*tlsutil.RecordData, error) {
	r := req.GetRecord()
	if r == nil {
		return nil, errors.New("record data is empty")
	}
	return encoding.RecordData(r), nil
}
