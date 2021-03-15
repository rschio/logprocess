package mysql

import (
	"context"
	"database/sql"
	"strings"
)

const insertServiceBatch = `-- name: InsertService :execresult
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
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`

const suffixDupKey = `ON DUPLICATE KEY UPDATE id=id`

func (q *Queries) InsertServiceBatch(ctx context.Context, args []InsertServiceParams,
) (sql.Result, error) {
	if len(args) == 1 {
		return q.InsertService(ctx, args[0])
	}
	stmt := insertServiceBatch
	argStmt := ", (?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt += strings.Repeat(argStmt, len(args)-1)

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
	stmt += suffixDupKey
	return q.db.ExecContext(ctx, stmt, sliceArgs...)
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
) VALUES (?,?,?,?,?,?,?,?,?,?,?)`

func (q *Queries) InsertRouteBatch(ctx context.Context, args []InsertRouteParams) (sql.Result, error) {
	if len(args) == 1 {
		return q.InsertRoute(ctx, args[0])
	}
	stmt := insertRouteBatch
	argStmt := ", (?,?,?,?,?,?,?,?,?,?,?)"
	stmt += strings.Repeat(argStmt, len(args)-1)
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
	stmt += suffixDupKey
	return q.db.ExecContext(ctx, stmt, sliceArgs...)
}

const insertResponseBatch = `
INSERT INTO responses (
	id, status, size, content_length,
	via, connection, access_control_allow_credentials,
	access_control_allow_origin, content_type, server
) VALUES (?,?,?,?,?,?,?,?,?,?)`

func (q *Queries) InsertResponseBatch(ctx context.Context, args []InsertResponseParams) (sql.Result, error) {
	if len(args) == 1 {
		return q.InsertResponse(ctx, args[0])
	}
	stmt := insertResponseBatch
	argStmt := ", (?,?,?,?,?,?,?,?,?,?)"
	stmt += strings.Repeat(argStmt, len(args)-1)
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
	return q.db.ExecContext(ctx, stmt, sliceArgs...)
}

const insertRequestBatch = `
INSERT INTO requests (
	id, method, uri, url, size, querystring,
	header_accept, header_host,
	header_user_agent
) VALUES (?,?,?,?,?,?,?,?,?)`

func (q *Queries) InsertRequestBatch(ctx context.Context, args []InsertRequestParams) (sql.Result, error) {
	if len(args) == 1 {
		return q.InsertRequest(ctx, args[0])
	}
	stmt := insertRequestBatch
	argStmt := ", (?,?,?,?,?,?,?,?,?)"
	stmt += strings.Repeat(argStmt, len(args)-1)
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
	return q.db.ExecContext(ctx, stmt, sliceArgs...)
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
) VALUES (?,?,?,?,?,?,?,?,?,?,?)`

func (q *Queries) InsertRecordBatch(ctx context.Context, args []InsertRecordParams) (sql.Result, error) {
	if len(args) == 1 {
		return q.InsertRecord(ctx, args[0])
	}
	stmt := insertRecordBatch
	argStmt := ", (?,?,?,?,?,?,?,?,?,?,?)"
	stmt += strings.Repeat(argStmt, len(args)-1)
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
	return q.db.ExecContext(ctx, stmt, sliceArgs...)
}
