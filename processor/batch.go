package processor

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
)

func readRecord(data []byte) (*Record, error) {
	rec := new(Record)
	if err := json.Unmarshal(data, rec); err != nil {
		return nil, err
	}
	if err := ValidRecord(rec); err != nil {
		return nil, err
	}
	return rec, nil
}

func readBatch(r io.Reader) ([]Record, error) {
	sc := bufio.NewScanner(r)
	records := make([]Record, 0, 200)
	for i := 0; sc.Scan(); i++ {
		rec, err := readRecord(sc.Bytes())
		if err != nil {
			return nil, err
		}
		records = append(records, *rec)
	}
	if err := sc.Err(); err != nil {
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

func InsertBatch(ctx context.Context, db Storage, r io.Reader) error {
	recs, err := readBatch(r)
	if err != nil {
		return err
	}
	return insertRecordBatch(ctx, db, recs)
}
