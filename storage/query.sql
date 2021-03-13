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

