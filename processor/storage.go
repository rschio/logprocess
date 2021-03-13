package processor

import "context"

type Storage interface {
	InsertRecord(context.Context, *Record) error
}
