package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rschio/logprocess/processor"
	"github.com/rschio/logprocess/processor/storage/mysql"
)

var (
	usrenv  = os.Getenv("MYSQL_USER")
	passenv = os.Getenv("MYSQL_PASSWORD")
	dbenv   = os.Getenv("MYSQL_DATABASE")
)

var (
	function = flag.Uint("f", 4, "What function to perform:\n\t"+
		"0 - insert json logs (read from stdin)\n\t"+
		"1 - consumer report\n\t"+
		"2 - services report\n\t"+
		"3 - average latencies report")
	timeout = flag.Duration("t", 5*time.Minute, "Timeout")
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

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("processor: ")

	db := connectDB()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, *timeout)
	defer cancel()

	var id string
	var err error
	if 1 <= *function && *function <= 2 {
		if flag.NArg() != 1 {
			log.Fatalf("provide one and just one id to generate the report")
		}
		id = flag.Arg(0)
	}

	w := os.Stdout
	switch *function {
	case 0:
		err = processor.InsertBatch(ctx, db, os.Stdin)
	case 1:
		err = processor.ConsumerReportCSV(ctx, w, db, id)
	case 2:
		err = processor.ServiceReportCSV(ctx, w, db, id)
	case 3:
		err = processor.AvgServicesLatenciesCSV(ctx, w, db)
	default:
		log.Println("flag f is required")
		flag.Usage()
	}
	if err != nil {
		log.Fatal(err)
	}
}
