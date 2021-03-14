package processor

import "context"

type Storage interface {
	InsertRecord(context.Context, *Record) error
	AverageServicesLatencies(context.Context) ([]ServiceLatencies, error)
}

type ServiceLatencies struct {
	ID           string
	Name         string
	AvgLatencies Latencies
}
