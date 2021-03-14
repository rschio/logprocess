package mysql

import (
	"strconv"
	"strings"
	"time"

	"github.com/rschio/logprocess/processor"
)

func toInsertServiceParams(s *processor.Service) *InsertServiceParams {
	return &InsertServiceParams{
		ID:             s.ID,
		CreatedAt:      time.Unix(s.CreatedAt, 0),
		UpdatedAt:      time.Unix(s.UpdatedAt, 0),
		Host:           s.Host,
		Name:           s.Name,
		Path:           s.Path,
		Port:           s.Port,
		Protocol:       s.Protocol,
		ReadTimeout:    s.ReadTimeout,
		WriteTimeout:   s.WriteTimeout,
		ConnectTimeout: s.ConnectTimeout,
		Retries:        s.Retries,
	}
}

func toInsertRouteParams(r *processor.Route) *InsertRouteParams {
	p := &InsertRouteParams{
		ID:            r.ID,
		CreatedAt:     time.Unix(r.CreatedAt, 0),
		UpdatedAt:     time.Unix(r.UpdatedAt, 0),
		Hosts:         strings.Join(r.Hosts, ","),
		Methods:       strings.Join(r.Methods, ","),
		Protocols:     strings.Join(r.Protocols, ","),
		RegexPriority: r.RegexPriority,
		ServiceID:     r.Service.ID,
	}
	if r.PreserveHost == true {
		p.PreserveHost = 1
	}
	if r.StripPath == true {
		p.StripPath = 1
	}
	return p
}

func toInsertResponseParams(r *processor.Response) (*InsertResponseParams, error) {
	h := r.Headers
	p := &InsertResponseParams{
		Status:                        r.Status,
		Via:                           h.Via,
		Size:                          int64(r.Size),
		Connection:                    h.Connection,
		AccessControlAllowCredentials: h.AccessControlAllowCredentials,
		AccessControlAllowOrigin:      h.AccessControlAllowOrigin,
		ContentType:                   h.ContentType,
		Server:                        h.Server,
	}
	cl, err := strconv.ParseInt(h.ContentLength, 10, 64)
	if err != nil {
		return nil, err
	}
	p.ContentLength = cl
	return p, nil
}

func toInsertRequestParams(r *processor.Request) (*InsertRequestParams, error) {
	p := &InsertRequestParams{
		Method:          r.Method,
		Uri:             r.URI,
		Url:             r.URL,
		Size:            int64(r.Size),
		HeaderAccept:    r.Headers.Accept,
		HeaderHost:      r.Headers.Host,
		HeaderUserAgent: r.Headers.UserAgent,
	}
	return p, nil
}

func toServiceLatencies(avgLat AverageLatencyByServiceRow) processor.ServiceLatencies {
	sl := processor.ServiceLatencies{ID: avgLat.ID, Name: avgLat.Name}
	sl.AvgLatencies.Proxy = int64(avgLat.AvgProxyLatency)
	sl.AvgLatencies.Gateway = int64(avgLat.AvgGatewayLatency)
	sl.AvgLatencies.Request = int64(avgLat.AvgRequestLatency)
	return sl
}

func toServicesLatencies(avgLats []AverageLatencyByServiceRow) []processor.ServiceLatencies {
	slice := make([]processor.ServiceLatencies, len(avgLats))
	for i, avgLat := range avgLats {
		slice[i] = toServiceLatencies(avgLat)
	}
	return slice
}

func toConsumerReportRow(row GetConsumerRequestsRow) processor.ConsumerReportRow {
	return processor.ConsumerReportRow(row)
}

func toConsumerReportRows(rows []GetConsumerRequestsRow) []processor.ConsumerReportRow {
	slice := make([]processor.ConsumerReportRow, len(rows))
	for i, row := range rows {
		slice[i] = toConsumerReportRow(row)
	}
	return slice
}
