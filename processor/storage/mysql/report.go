package mysql

import (
	"context"

	"github.com/rschio/logprocess/processor"
)

const reportRequests = `
SELECT
	r.id,
	r.consumer_id,
	r.upstream_uri AS upstream_URI,
	r.response_id,
	r.request_id,
	r.route_id,
	r.service_id,
	r.proxy_latency,
	r.gateway_latency,
	r.request_latency,
	r.client_ip as client_IP,
	r.started_at,
	rp.status AS rsp_status,
	rp.size AS rsp_size,
	rp.content_length AS rsp_content_length,
	rp.via AS rsp_via,
	rp.connection AS rsp_connection,
	rp.access_control_allow_credentials AS rsp_access_control_allow_credentials,
	rp.access_control_allow_origin AS rsp_access_control_allow_origin,
	rp.content_type AS rsp_content_type,
	rp.server AS rsp_server,
	rq.method AS req_method,
	rq.uri AS req_URI,
	rq.url AS req_URL,
	rq.size AS req_Size,
	rq.querystring AS req_querystring,
	rq.header_accept AS req_header_accept,
	rq.header_host AS req_header_host,
	rq.header_user_agent AS req_user_agent
	FROM records r
	INNER JOIN responses rp ON r.response_id = rp.id
	INNER JOIN requests rq ON r.request_id = rq.id
`

// GetConsumerRequests return information about consumer requests.
func (q *Queries) GetConsumerRequests(ctx context.Context, consumerID string) ([]processor.ReportRow, error) {
	stmt := reportRequests + " WHERE r.consumer_id = ?"
	return q.getReportRequests(ctx, stmt, consumerID)
}

// GetServiceRequests return information about service requests.
func (q *Queries) GetServiceRequests(ctx context.Context, serviceID string) ([]processor.ReportRow, error) {
	stmt := reportRequests + " WHERE r.service_id = ?"
	return q.getReportRequests(ctx, stmt, serviceID)
}

func (q *Queries) getReportRequests(ctx context.Context, stmt, id string) ([]processor.ReportRow, error) {
	rows, err := q.db.QueryContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []processor.ReportRow
	for rows.Next() {
		var i processor.ReportRow
		if err := rows.Scan(
			&i.ID,
			&i.ConsumerID,
			&i.UpstreamURI,
			&i.ResponseID,
			&i.RequestID,
			&i.RouteID,
			&i.ServiceID,
			&i.ProxyLatency,
			&i.GatewayLatency,
			&i.RequestLatency,
			&i.ClientIP,
			&i.StartedAt,
			&i.RspStatus,
			&i.RspSize,
			&i.RspContentLength,
			&i.RspVia,
			&i.RspConnection,
			&i.RspAccessControlAllowCredentials,
			&i.RspAccessControlAllowOrigin,
			&i.RspContentType,
			&i.RspServer,
			&i.ReqMethod,
			&i.ReqURI,
			&i.ReqURL,
			&i.ReqSize,
			&i.ReqQuerystring,
			&i.ReqHeaderAccept,
			&i.ReqHeaderHost,
			&i.ReqUserAgent,
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
