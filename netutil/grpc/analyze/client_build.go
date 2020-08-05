// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"errors"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"github.com/luids-io/core/apiservice"
	"github.com/luids-io/core/grpctls"
	"github.com/luids-io/core/option"
	"github.com/luids-io/core/yalogi"
)

// ClientBuilder returns builder function for the apiservice
func ClientBuilder(opt ...ClientOption) apiservice.BuildFn {
	return func(def apiservice.ServiceDef, logger yalogi.Logger) (apiservice.Service, error) {
		//validates definition
		err := def.Validate()
		if err != nil {
			return nil, err
		}
		opts := make([]grpc.DialOption, 0)
		if def.Metrics {
			opts = append(opts, grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor))
			opts = append(opts, grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
		}
		//dial grpc
		dial, err := grpctls.Dial(def.Endpoint, def.ClientCfg(), opts...)
		if err != nil {
			return nil, err
		}
		if def.Log {
			opt = append(opt, SetLogger(logger))
		}
		if len(def.Opts) > 0 {
			value, ok, err := option.Int(def.Opts, "buffer")
			if err != nil {
				return nil, errors.New("invalid 'buffer'")
			}
			if ok && value > 0 {
				opt = append(opt, SetPacketBuffer(value))
			}
		}
		//creates client
		client := NewClient(dial, opt...)
		return client, nil
	}
}

func init() {
	apiservice.RegisterBuilder(ServiceName(), ClientBuilder())
}
