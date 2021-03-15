package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dnlo/struct2csv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rschio/logprocess/processor"
	"github.com/rschio/logprocess/processor/storage/mysql"
)

var (
	usrenv  = os.Getenv("MYSQL_USER")
	passenv = os.Getenv("MYSQL_PASSWORD")
	dbenv   = os.Getenv("MYSQL_DATABASE")
)

func connectDB() processor.Storage {
	info := fmt.Sprintf("%s:%s@/%s?parseTime=true", usrenv, passenv, dbenv)
	conn, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
	return mysql.NewMySQL(conn)
}

func genCSV(ls []processor.ServiceLatencies) {
	w := struct2csv.NewWriter(os.Stdout)
	err := w.WriteStructs(ls)
	if err != nil {
		log.Fatal(err)
	}
}
func genReportCSV(ls []processor.ReportRow) {
	w := struct2csv.NewWriter(os.Stdout)
	err := w.WriteStructs(ls)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := connectDB()
	sc := bufio.NewScanner(os.Stdin)
	logs := make([]processor.Record, 0, 200)
	for i := 0; sc.Scan(); i++ {
		jsonlog := sc.Bytes()
		l := processor.Record{}
		err := json.Unmarshal(jsonlog, &l)
		if err != nil {
			log.Printf("%d: %v", i, err)
			continue
		}
		if err := processor.ValidRecord(&l); err != nil {
			log.Println(err)
		} else {
			logs = append(logs, l)
		}
	}
	if err := sc.Err(); err != nil {
		log.Println(err)
	}
	ctx := context.Background()
	s := len(logs) / 1000
	for i := 0; i < s; i++ {
		err := db.InsertRecordBatch(ctx, logs[i*1000:(i+1)*1000])
		if err != nil {
			log.Fatal(err)
		}
	}
	//genCSV(as)
	//genReportCSV(as)
	//for _, a := range as {
	//	fmt.Printf("%+v\n", a)
	//}
	//for _, l := range logs {
	//	err := db.InsertRecord(ctx, &l)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}
