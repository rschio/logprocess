package processor

import (
	"context"
	"io"
	"time"

	"github.com/dnlo/struct2csv"
)

type ServiceLatencies struct {
	ID           string
	Name         string
	AvgLatencies Latencies
}

type ReportRow struct {
	ID                               int64
	ConsumerID                       string
	UpstreamURI                      string
	ResponseID                       string
	RequestID                        string
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
	ReqQuerystring                   string
	ReqHeaderAccept                  string
	ReqHeaderHost                    string
	ReqUserAgent                     string
}

func writeReportCSV(w io.Writer, report []ReportRow) error {
	wr := struct2csv.NewWriter(w)
	return wr.WriteStructs(report)
}

func ConsumerReportCSV(ctx context.Context, w io.Writer, db Storage, id string) error {
	report, err := db.ConsumerReport(ctx, id)
	if err != nil {
		return err
	}
	return writeReportCSV(w, report)
}

func ServiceReportCSV(ctx context.Context, w io.Writer, db Storage, id string) error {
	report, err := db.ServiceReport(ctx, id)
	if err != nil {
		return err
	}
	return writeReportCSV(w, report)
}

func AvgServicesLatenciesCSV(ctx context.Context, w io.Writer, db Storage) error {
	latencies, err := db.AverageServicesLatencies(ctx)
	if err != nil {
		return err
	}
	wr := struct2csv.NewWriter(w)
	return wr.WriteStructs(latencies)
}
