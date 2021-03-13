-- name: GetRecords :many
SELECT * 
	FROM records r
	JOIN responses rp ON r.response_id = rp.id
	JOIN requests rq ON r.request_id = rq.id
	JOIN routes rt ON r.route_id = rt.id
	JOIN services s ON r.service_id = s.id
;

-- name: AverageLatencyByService :many
SELECT 
	AVG(proxy_latency) AS proxy_latency, 
	AVG(gateway_latency) AS gateway_latency, 
	AVG(request_latency) AS request_latency
	FROM records 
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
);

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
);
	
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
