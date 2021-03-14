// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package mysql

import (
	"context"
	"database/sql"
	"time"
)

const averageLatencyByService = `-- name: AverageLatencyByService :many
SELECT 
	AVG(proxy_latency) AS proxy_latency, 
	AVG(gateway_latency) AS gateway_latency, 
	AVG(request_latency) AS request_latency
	FROM records 
	GROUP BY service_id
`

type AverageLatencyByServiceRow struct {
	ProxyLatency   interface{}
	GatewayLatency interface{}
	RequestLatency interface{}
}

func (q *Queries) AverageLatencyByService(ctx context.Context) ([]AverageLatencyByServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, averageLatencyByService)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AverageLatencyByServiceRow
	for rows.Next() {
		var i AverageLatencyByServiceRow
		if err := rows.Scan(&i.ProxyLatency, &i.GatewayLatency, &i.RequestLatency); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecords = `-- name: GetRecords :many
SELECT r.id, consumer_id, upstream_uri, response_id, request_id, route_id, r.service_id, proxy_latency, gateway_latency, request_latency, client_ip, started_at, rp.id, status, rp.size, content_length, via, connection, access_control_allow_credentials, access_control_allow_origin, content_type, server, rq.id, method, uri, url, rq.size, header_accept, header_host, header_user_agent, rt.id, rt.created_at, rt.updated_at, hosts, methods, preserve_host, protocols, regex_priority, rt.service_id, strip_path, s.id, s.created_at, s.updated_at, host, name, path, port, protocol, read_timeout, write_timeout, connect_timeout, retries 
	FROM records r
	JOIN responses rp ON r.response_id = rp.id
	JOIN requests rq ON r.request_id = rq.id
	JOIN routes rt ON r.route_id = rt.id
	JOIN services s ON r.service_id = s.id
`

type GetRecordsRow struct {
	ID                            int64
	ConsumerID                    string
	UpstreamUri                   string
	ResponseID                    int64
	RequestID                     int64
	RouteID                       string
	ServiceID                     string
	ProxyLatency                  int64
	GatewayLatency                int64
	RequestLatency                int64
	ClientIp                      string
	StartedAt                     time.Time
	ID_2                          int64
	Status                        int64
	Size                          int64
	ContentLength                 int64
	Via                           string
	Connection                    string
	AccessControlAllowCredentials string
	AccessControlAllowOrigin      string
	ContentType                   string
	Server                        string
	ID_3                          int64
	Method                        string
	Uri                           string
	Url                           string
	Size_2                        int64
	HeaderAccept                  string
	HeaderHost                    string
	HeaderUserAgent               string
	ID_4                          string
	CreatedAt                     time.Time
	UpdatedAt                     time.Time
	Hosts                         string
	Methods                       string
	PreserveHost                  int32
	Protocols                     string
	RegexPriority                 int64
	ServiceID_2                   string
	StripPath                     int32
	ID_5                          string
	CreatedAt_2                   time.Time
	UpdatedAt_2                   time.Time
	Host                          string
	Name                          string
	Path                          string
	Port                          int64
	Protocol                      string
	ReadTimeout                   int64
	WriteTimeout                  int64
	ConnectTimeout                int64
	Retries                       int64
}

func (q *Queries) GetRecords(ctx context.Context) ([]GetRecordsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecords)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecordsRow
	for rows.Next() {
		var i GetRecordsRow
		if err := rows.Scan(
			&i.ID,
			&i.ConsumerID,
			&i.UpstreamUri,
			&i.ResponseID,
			&i.RequestID,
			&i.RouteID,
			&i.ServiceID,
			&i.ProxyLatency,
			&i.GatewayLatency,
			&i.RequestLatency,
			&i.ClientIp,
			&i.StartedAt,
			&i.ID_2,
			&i.Status,
			&i.Size,
			&i.ContentLength,
			&i.Via,
			&i.Connection,
			&i.AccessControlAllowCredentials,
			&i.AccessControlAllowOrigin,
			&i.ContentType,
			&i.Server,
			&i.ID_3,
			&i.Method,
			&i.Uri,
			&i.Url,
			&i.Size_2,
			&i.HeaderAccept,
			&i.HeaderHost,
			&i.HeaderUserAgent,
			&i.ID_4,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Hosts,
			&i.Methods,
			&i.PreserveHost,
			&i.Protocols,
			&i.RegexPriority,
			&i.ServiceID_2,
			&i.StripPath,
			&i.ID_5,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Host,
			&i.Name,
			&i.Path,
			&i.Port,
			&i.Protocol,
			&i.ReadTimeout,
			&i.WriteTimeout,
			&i.ConnectTimeout,
			&i.Retries,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertRecord = `-- name: InsertRecord :execresult
INSERT INTO records (
	consumer_id,
	upstream_uri,
	response_id,
	request_id,
	route_id,
	service_id,
	proxy_latency,
	gateway_latency,
	request_latency,
	client_ip,
	started_at
) VALUES (
	?,?,?,?,?,?,?,?,?,?,?
)
`

type InsertRecordParams struct {
	ConsumerID     string
	UpstreamUri    string
	ResponseID     int64
	RequestID      int64
	RouteID        string
	ServiceID      string
	ProxyLatency   int64
	GatewayLatency int64
	RequestLatency int64
	ClientIp       string
	StartedAt      time.Time
}

func (q *Queries) InsertRecord(ctx context.Context, arg InsertRecordParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertRecord,
		arg.ConsumerID,
		arg.UpstreamUri,
		arg.ResponseID,
		arg.RequestID,
		arg.RouteID,
		arg.ServiceID,
		arg.ProxyLatency,
		arg.GatewayLatency,
		arg.RequestLatency,
		arg.ClientIp,
		arg.StartedAt,
	)
}

const insertRequest = `-- name: InsertRequest :execresult
INSERT INTO requests (
	method, uri, url, size,
	header_accept, header_host,
	header_user_agent
) VALUES (
	?,?,?,?,?,?,?
)
`

type InsertRequestParams struct {
	Method          string
	Uri             string
	Url             string
	Size            int64
	HeaderAccept    string
	HeaderHost      string
	HeaderUserAgent string
}

func (q *Queries) InsertRequest(ctx context.Context, arg InsertRequestParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertRequest,
		arg.Method,
		arg.Uri,
		arg.Url,
		arg.Size,
		arg.HeaderAccept,
		arg.HeaderHost,
		arg.HeaderUserAgent,
	)
}

const insertResponse = `-- name: InsertResponse :execresult
INSERT INTO responses (
	status, size, content_length,
	via, connection, access_control_allow_credentials,
	access_control_allow_origin, content_type, server
) VALUES (
	?,?,?,?,?,?,?,?,?
)
`

type InsertResponseParams struct {
	Status                        int64
	Size                          int64
	ContentLength                 int64
	Via                           string
	Connection                    string
	AccessControlAllowCredentials string
	AccessControlAllowOrigin      string
	ContentType                   string
	Server                        string
}

func (q *Queries) InsertResponse(ctx context.Context, arg InsertResponseParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertResponse,
		arg.Status,
		arg.Size,
		arg.ContentLength,
		arg.Via,
		arg.Connection,
		arg.AccessControlAllowCredentials,
		arg.AccessControlAllowOrigin,
		arg.ContentType,
		arg.Server,
	)
}

const insertRoute = `-- name: InsertRoute :execresult
INSERT INTO routes (
	id,
	created_at,
	updated_at,
	hosts,
	methods,
	preserve_host,
	protocols,
	regex_priority,
	service_id,
	strip_path
) VALUES (
	?,?,?,?,?,?,?,?,?,?
) ON DUPLICATE KEY UPDATE id=id
`

type InsertRouteParams struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Hosts         string
	Methods       string
	PreserveHost  int32
	Protocols     string
	RegexPriority int64
	ServiceID     string
	StripPath     int32
}

func (q *Queries) InsertRoute(ctx context.Context, arg InsertRouteParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertRoute,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Hosts,
		arg.Methods,
		arg.PreserveHost,
		arg.Protocols,
		arg.RegexPriority,
		arg.ServiceID,
		arg.StripPath,
	)
}

const insertService = `-- name: InsertService :execresult
INSERT INTO services (
	id,
	created_at,
	updated_at,
	host,
	name,
	path,
	port,
	protocol,
	read_timeout,
	write_timeout,
	connect_timeout,
	retries
) VALUES (
	?,?,?,?,?,?,?,?,?,?,?,?
) ON DUPLICATE KEY UPDATE id=id
`

type InsertServiceParams struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Host           string
	Name           string
	Path           string
	Port           int64
	Protocol       string
	ReadTimeout    int64
	WriteTimeout   int64
	ConnectTimeout int64
	Retries        int64
}

func (q *Queries) InsertService(ctx context.Context, arg InsertServiceParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertService,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Host,
		arg.Name,
		arg.Path,
		arg.Port,
		arg.Protocol,
		arg.ReadTimeout,
		arg.WriteTimeout,
		arg.ConnectTimeout,
		arg.Retries,
	)
}
