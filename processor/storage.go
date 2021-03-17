package processor

import (
	"context"
)

type Storage interface {
	InsertRecord(context.Context, *Record) error
	InsertRecordBatch(context.Context, []Record) error
	AverageServicesLatencies(context.Context) ([]ServiceLatencies, error)
	ConsumerReport(ctx context.Context, id string) ([]ReportRow, error)
	ServiceReport(ctx context.Context, id string) ([]ReportRow, error)
}
