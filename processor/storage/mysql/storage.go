package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/rschio/logprocess/processor"
)

type MySQL struct {
	db *sql.DB
	q  *Queries
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db, q: New(db)}
}

func (m *MySQL) InsertRecordBatch(ctx context.Context, rs []processor.Record) error {
	ss := make([]InsertServiceParams, len(rs))
	srp := make([]InsertResponseParams, len(rs))
	srq := make([]InsertRequestParams, len(rs))
	srt := make([]InsertRouteParams, len(rs))
	srr := make([]InsertRecordParams, len(rs))
	var err error
	for i, r := range rs {
		ss[i] = *toInsertServiceParams(&r.Service)
		v0, err := toInsertResponseParams(&r.Response)
		if err != nil {
			return err
		}
		srp[i] = *v0
		v1, err := toInsertRequestParams(&r.Request)
		if err != nil {
			return err
		}
		srq[i] = *v1
		srt[i] = *toInsertRouteParams(&r.Route)

		srr[i] = InsertRecordParams{
			ConsumerID:     r.AuthenticatedEntity.ConsumerID.UUID,
			UpstreamUri:    r.UpstreamURI,
			ResponseID:     srp[i].ID,
			RequestID:      srq[i].ID,
			RouteID:        r.Route.ID,
			ServiceID:      r.Service.ID,
			ProxyLatency:   r.Latencies.Proxy,
			GatewayLatency: r.Latencies.Gateway,
			RequestLatency: r.Latencies.Request,
			ClientIp:       r.ClientIP,
			StartedAt:      time.Unix(r.StartedAt, 0),
		}

	}
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := m.q.WithTx(tx)
	_, err = q.InsertServiceBatch(ctx, ss)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = q.InsertResponseBatch(ctx, srp)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = q.InsertRequestBatch(ctx, srq)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = q.InsertRouteBatch(ctx, srt)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = q.InsertRecordBatch(ctx, srr)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (m *MySQL) ServiceReport(ctx context.Context, id string) ([]processor.ReportRow, error) {
	rows, err := m.q.GetServiceRequests(ctx, id)
	if err != nil {
		return nil, err
	}
	return toReportRowsService(rows), nil
}

func (m *MySQL) ConsumerReport(ctx context.Context, id string) ([]processor.ReportRow, error) {
	rows, err := m.q.GetConsumerRequests(ctx, id)
	if err != nil {
		return nil, err
	}
	return toReportRows(rows), nil
}

func (m *MySQL) AverageServicesLatencies(ctx context.Context) ([]processor.ServiceLatencies, error) {
	avgLats, err := m.q.AverageLatencyByService(ctx)
	if err != nil {
		return nil, err
	}
	return toServicesLatencies(avgLats), nil
}

func (m *MySQL) InsertRecord(ctx context.Context, r *processor.Record) error {
	respParams, err := toInsertResponseParams(&r.Response)
	if err != nil {
		return err
	}
	reqParams, err := toInsertRequestParams(&r.Request)
	if err != nil {
		return err
	}
	routeParams := toInsertRouteParams(&r.Route)
	serviceParams := toInsertServiceParams(&r.Service)

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := m.q.WithTx(tx)
	_, err = q.InsertResponse(ctx, *respParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = q.InsertRequest(ctx, *reqParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = q.InsertRoute(ctx, *routeParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = q.InsertService(ctx, *serviceParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	params := InsertRecordParams{
		ConsumerID:     r.AuthenticatedEntity.ConsumerID.UUID,
		UpstreamUri:    r.UpstreamURI,
		ResponseID:     respParams.ID,
		RequestID:      reqParams.ID,
		RouteID:        r.Route.ID,
		ServiceID:      r.Service.ID,
		ProxyLatency:   r.Latencies.Proxy,
		GatewayLatency: r.Latencies.Gateway,
		RequestLatency: r.Latencies.Request,
		ClientIp:       r.ClientIP,
		StartedAt:      time.Unix(r.StartedAt, 0),
	}
	_, err = q.InsertRecord(ctx, params)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
