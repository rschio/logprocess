package processor

import (
	"context"
	"encoding/json"
	"io"
)

func readBatch(r io.Reader) ([]Record, error) {
	dec := json.NewDecoder(r)
	records := make([]Record, 0, 200)
	var err error
	for {
		rec := new(Record)
		if err = dec.Decode(rec); err != nil {
			break
		}
		if err = ValidRecord(rec); err != nil {
			return nil, err
		}
		records = append(records, *rec)
	}
	if err != io.EOF {
		return nil, err
	}
	return records, nil
}

func insertRecordBatch(ctx context.Context, db Storage, recs []Record) error {
	const bsize = 1000
	if len(recs) < 1 {
		return nil
	}
	batches := len(recs) / bsize
	rem := len(recs) % bsize

	i := 0
	for i = 0; i < batches; i++ {
		err := db.InsertRecordBatch(ctx, recs[i*bsize:(i+1)*bsize])
		if err != nil {
			return err
		}
	}
	if rem > 0 {
		return db.InsertRecordBatch(ctx, recs[i*bsize:])
	}
	return nil
}

// InsertBatch inserts a stream of log Records into the db Storage.
// The insertion is made in batches.
func InsertBatch(ctx context.Context, db Storage, r io.Reader) error {
	recs, err := readBatch(r)
	if err != nil {
		return err
	}
	return insertRecordBatch(ctx, db, recs)
}
