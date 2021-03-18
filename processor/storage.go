package processor

import (
	"context"
)

// Storage is used to persist data, usually a database.
type Storage interface {
	// InserRecord inserts the Record into the Storage.
	InsertRecord(context.Context, *Record) error

	// InserRecordBatch inserts a batch of Records into the Storage.
	InsertRecordBatch(context.Context, []Record) error

	// AverageServicesLatencies returns the latencies of all services.
	AverageServicesLatencies(context.Context) ([]ServiceLatencies, error)

	// ConsumerReport returns information about the specified consumer's
	// requests.
	ConsumerReport(ctx context.Context, id string) ([]ReportRow, error)

	//  ServiceReport returns information about the specified service's
	// requests.
	ServiceReport(ctx context.Context, id string) ([]ReportRow, error)
}
