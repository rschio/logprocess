package processor

import (
	"context"
	"time"
)

type Storage interface {
	InsertRecord(context.Context, *Record) error
	AverageServicesLatencies(context.Context) ([]ServiceLatencies, error)
	ConsumerReport(ctx context.Context, id string) ([]ConsumerReportRow, error)
}

type ServiceLatencies struct {
	ID           string
	Name         string
	AvgLatencies Latencies
}

type ConsumerReportRow struct {
	ID                               int64
	ConsumerID                       string
	UpstreamURI                      string
	ResponseID                       int64
	RequestID                        int64
	RouteID                          string
	ServiceID                        string
	ProxyLatency                     int64
	GatewayLatency                   int64
	RequestLatency                   int64
	ClientIP                         string
	StartedAt                        time.Time
	RspStatus                        int64
	RspSize                          int64
	RspContentLength                 int64
	RspVia                           string
	RspConnection                    string
	RspAccessControlAllowCredentials string
	RspAccessControlAllowOrigin      string
	RspContentType                   string
	RspServer                        string
	ReqMethod                        string
	ReqURI                           string
	ReqURL                           string
	ReqSize                          int64
	ReqHeaderAccept                  string
	ReqHeaderHost                    string
	ReqUserAgent                     string
}
