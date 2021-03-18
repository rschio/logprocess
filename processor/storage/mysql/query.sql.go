// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package mysql

import (
	"context"
)

const averageLatencyByService = `-- name: AverageLatencyByService :many
SELECT 
	s.id, 
	s.name, 
	ROUND(AVG(proxy_latency),   0)   AS avg_proxy_latency,
	ROUND(AVG(gateway_latency), 0) AS avg_gateway_latency,
	ROUND(AVG(request_latency), 0) AS avg_request_latency
	FROM  records 
	INNER JOIN services s ON records.service_id = s.id 
	GROUP BY service_id
`

type AverageLatencyByServiceRow struct {
	ID                string
	Name              string
	AvgProxyLatency   float64
	AvgGatewayLatency float64
	AvgRequestLatency float64
}

// AverageLatencyByService returns the average latency of each service.
func (q *Queries) AverageLatencyByService(ctx context.Context) ([]AverageLatencyByServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, averageLatencyByService)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AverageLatencyByServiceRow
	for rows.Next() {
		var i AverageLatencyByServiceRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AvgProxyLatency,
			&i.AvgGatewayLatency,
			&i.AvgRequestLatency,
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
