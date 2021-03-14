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
	res, err := q.InsertResponse(ctx, *respParams)
	if err != nil {
		tx.Rollback()
		return err
	}
	respID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err = q.InsertRequest(ctx, *reqParams)
	if err != nil {
		tx.Rollback()
		return err
	}
	reqID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err = q.InsertRoute(ctx, *routeParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err = q.InsertService(ctx, *serviceParams)
	if err != nil {
		tx.Rollback()
		return err
	}

	params := InsertRecordParams{
		ConsumerID:     r.AuthenticatedEntity.ConsumerID.UUID,
		UpstreamUri:    r.UpstreamURI,
		ResponseID:     respID,
		RequestID:      reqID,
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
