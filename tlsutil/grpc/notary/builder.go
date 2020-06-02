// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notary

import (
	"errors"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"github.com/luids-io/core/apiservice"
	"github.com/luids-io/core/grpctls"
	"github.com/luids-io/core/option"
	"github.com/luids-io/core/yalogi"
)

// ClientBuilder returns builder function
func ClientBuilder(opt ...ClientOption) apiservice.BuildFn {
	return func(def apiservice.ServiceDef, logger yalogi.Logger) (apiservice.Service, error) {
		//validates definition
		err := def.Validate()
		if err != nil {
			return nil, err
		}
		dopts := make([]grpc.DialOption, 0)
		if def.Metrics {
			dopts = append(dopts, grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor))
			dopts = append(dopts, grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
		}
		//dial grpc
		dial, err := grpctls.Dial(def.Endpoint, def.ClientCfg(), dopts...)
		if err != nil {
			return nil, err
		}
		if def.Log {
			opt = append(opt, SetLogger(logger))
		}
		if len(def.Opts) > 0 {
			// parse and set cache options
			ttl, err := parseCacheOpts(def.Opts)
			if err != nil {
				return nil, err
			}
			if ttl > 0 {
				opt = append(opt, SetCache(time.Duration(ttl)*time.Second))
			}
		}
		//creates client
		client := NewClient(dial, opt...)
		return client, nil
	}
}

func parseCacheOpts(opts map[string]interface{}) (int, error) {
	var ttl int
	// get ttl
	value, ok, err := option.Int(opts, "ttl")
	if err != nil {
		return 0, err
	}
	if ok {
		if value < 0 {
			return 0, errors.New("invalid 'ttl'")
		}
		ttl = value
	}
	return ttl, nil
}

func init() {
	apiservice.RegisterBuilder(ServiceName(), ClientBuilder())
}
