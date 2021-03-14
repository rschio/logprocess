-- name: GetRecords :many
SELECT * 
	FROM records r
	JOIN responses rp ON r.response_id = rp.id
	JOIN requests rq ON r.request_id = rq.id
	JOIN routes rt ON r.route_id = rt.id
	JOIN services s ON r.service_id = s.id
;

-- name: GetConsumerRequests :many
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
	rq.header_accept AS req_header_accept,
	rq.header_host AS req_header_host,
	rq.header_user_agent AS req_user_agent
	FROM records r
	INNER JOIN responses rp ON r.response_id = rp.id
	INNER JOIN requests rq ON r.request_id = rq.id
	WHERE r.consumer_id = ?
;



-- name: AverageLatencyByService :many
SELECT 
	s.id, 
	s.name, 
	ROUND(AVG(proxy_latency),   0)   AS avg_proxy_latency,
	ROUND(AVG(gateway_latency), 0) AS avg_gateway_latency,
	ROUND(AVG(request_latency), 0) AS avg_request_latency
	FROM  records 
	INNER JOIN services s ON records.service_id = s.id 
	GROUP BY service_id
;

-- name: InsertRequest :execresult
INSERT INTO requests (
	method, uri, url, size,
	header_accept, header_host,
	header_user_agent
) VALUES (
	?,?,?,?,?,?,?
);

-- name: InsertResponse :execresult
INSERT INTO responses (
	status, size, content_length,
	via, connection, access_control_allow_credentials,
	access_control_allow_origin, content_type, server
) VALUES (
	?,?,?,?,?,?,?,?,?
);

-- name: InsertRoute :execresult
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
) ON DUPLICATE KEY UPDATE id=id;

-- name: InsertService :execresult
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
) ON DUPLICATE KEY UPDATE id=id;
	
-- name: InsertRecord :execresult
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
);
