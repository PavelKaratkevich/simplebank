package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

// metadata from HTTP request
// map[grpcgateway-accept:[*/*] grpcgateway-cache-control:[no-cache] grpcgateway-content-type:[application/json] grpcgateway-user-agent:[PostmanRuntime/7.29.2] x-forwarded-for:[::1] x-forwarded-host:[localhost:8080]]

// metadata from RPC request:
//  map[:authority:[localhost:9090] accept-encoding:[identity] content-type:[application/grpc] grpc-accept-encoding:[identity,deflate,gzip] user-agent:[grpc-node-js/1.6.7]]


func ExtractMetadata(ctx context.Context) *Metadata {
	mtd := &Metadata{}

	// this part works for HTTP requests only
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %v", md)
		if userAgent := md.Get("grpcgateway-user-agent"); len(userAgent) > 0 {
			mtd.UserAgent = userAgent[0]
		}

		// this part relates to gRPC request only
		if userAgent := md.Get("user-agent"); len(userAgent) > 0 {
			mtd.UserAgent = userAgent[0]
		}

		// http again
		if clientIP := md.Get("x-forwarded-for"); len(clientIP) > 0 {
			mtd.ClientIP = clientIP[0]
		}
	}

	// Get client IP address out of metadata from gRPC request
	if p, ok := peer.FromContext(ctx); ok {
		mtd.ClientIP = p.Addr.String()
	}

	return mtd
}
