package mysql

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

// createQuery assembly a query from using a stmt, a optional
// suffix and the total number of args.
func createQuery(stmt, suffix string, totalArgs int) string {
	narg := strings.Count(stmt, ",") + 1
	rows := totalArgs / narg
	vals := strings.Repeat("?,", narg)
	// Trim the last comma.
	vals = "(" + vals[:len(vals)-1] + "),"
	stmt += "VALUES "
	stmt += strings.Repeat(vals, rows)
	// Trim the last comma again.
	stmt = stmt[:len(stmt)-1]
	stmt += " " + suffix
	return stmt
}

func (q *Queries) execBatch(ctx context.Context, stmt, suffix string,
	args ...interface{}) (sql.Result, error) {

	totalArgs := len(args)
	stmt = createQuery(stmt, suffix, totalArgs)
	return q.db.ExecContext(ctx, stmt, args...)
}

const suffixDupKey = `ON DUPLICATE KEY UPDATE id=id`

const insertServiceBatch = `
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
)`

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

func (q *Queries) InsertServiceBatch(ctx context.Context, args []InsertServiceParams,
) (sql.Result, error) {
	sliceArgs := make([]interface{}, 0, len(args)*12)
	for _, arg := range args {
		sliceArgs = append(sliceArgs,
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
	stmt := insertServiceBatch
	return q.execBatch(ctx, stmt, suffixDupKey, sliceArgs...)
}

const insertRouteBatch = `
INSERT INTO routes (
	id,
	created_at,
	updated_at,
	hosts,
	methods,
	paths,
	preserve_host,
	protocols,
	regex_priority,
	service_id,
	strip_path
)`

type InsertRouteParams struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Hosts         string
	Methods       string
	Paths         string
	PreserveHost  int32
	Protocols     string
	RegexPriority int64
	ServiceID     string
	StripPath     int32
}

func (q *Queries) InsertRouteBatch(ctx context.Context, args []InsertRouteParams) (sql.Result, error) {
	sliceArgs := make([]interface{}, 0, len(args)*11)
	for _, arg := range args {
		sliceArgs = append(sliceArgs,
			arg.ID,
			arg.CreatedAt,
			arg.UpdatedAt,
			arg.Hosts,
			arg.Methods,
			arg.Paths,
			arg.PreserveHost,
			arg.Protocols,
			arg.RegexPriority,
			arg.ServiceID,
			arg.StripPath,
		)
	}
	stmt := insertRouteBatch
	return q.execBatch(ctx, stmt, suffixDupKey, sliceArgs...)
}

const insertResponseBatch = `
INSERT INTO responses (
	id,
	status,
	size,
	content_length,
	via,
	connection, access_control_allow_credentials,
	access_control_allow_origin,
	content_type,
	server
)`

type InsertResponseParams struct {
	ID                            string
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

func (q *Queries) InsertResponseBatch(ctx context.Context, args []InsertResponseParams) (sql.Result, error) {
	sliceArgs := make([]interface{}, 0, len(args)*10)
	for _, arg := range args {
		sliceArgs = append(sliceArgs,
			arg.ID,
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
	stmt := insertResponseBatch
	return q.execBatch(ctx, stmt, "", sliceArgs...)
}

const insertRequestBatch = `
INSERT INTO requests (
	id,
	method,
	uri,
	url,
	size,
	querystring,
	header_accept,
	header_host,
	header_user_agent
)`

type InsertRequestParams struct {
	ID              string
	Method          string
	Uri             string
	Url             string
	Size            int64
	Querystring     string
	HeaderAccept    string
	HeaderHost      string
	HeaderUserAgent string
}

func (q *Queries) InsertRequestBatch(ctx context.Context, args []InsertRequestParams) (sql.Result, error) {
	sliceArgs := make([]interface{}, 0, len(args)*9)
	for _, arg := range args {
		sliceArgs = append(sliceArgs,
			arg.ID,
			arg.Method,
			arg.Uri,
			arg.Url,
			arg.Size,
			arg.Querystring,
			arg.HeaderAccept,
			arg.HeaderHost,
			arg.HeaderUserAgent,
		)
	}
	stmt := insertRequestBatch
	return q.execBatch(ctx, stmt, "", sliceArgs...)
}

const insertRecordBatch = `
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
)`

type InsertRecordParams struct {
	ConsumerID     string
	UpstreamUri    string
	ResponseID     string
	RequestID      string
	RouteID        string
	ServiceID      string
	ProxyLatency   int64
	GatewayLatency int64
	RequestLatency int64
	ClientIp       string
	StartedAt      time.Time
}

func (q *Queries) InsertRecordBatch(ctx context.Context, args []InsertRecordParams) (sql.Result, error) {
	sliceArgs := make([]interface{}, 0, len(args)*11)
	for _, arg := range args {
		sliceArgs = append(sliceArgs,
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
	stmt := insertRecordBatch
	return q.execBatch(ctx, stmt, "", sliceArgs...)
}
