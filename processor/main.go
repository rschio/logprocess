package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	usrenv  = os.Getenv("MYSQL_USER")
	passenv = os.Getenv("MYSQL_PASSWORD")
	dbenv   = os.Getenv("MYSQL_DATABASE")
)

func connectDB() {
	info := fmt.Sprintf("%s:%s@/%s", usrenv, passenv, dbenv)
	conn, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
}

func main() {
	connectDB()
	//	sc := bufio.NewScanner(os.Stdin)
	//	logs := make([]Log, 0, 200)
	//	for sc.Scan() {
	//		jsonlog := sc.Bytes()
	//		l := Log{}
	//		json.Unmarshal(jsonlog, &l)
	//		logs = append(logs, l)
	//	}
	//	if err := sc.Err(); err != nil {
	//		log.Println(err)
	//	}
	//	latencies(logs)
}

func latencies(logs []Record) {
	m := make(map[string][]Latencies)
	for _, l := range logs {
		id := l.Service.ID
		m[id] = append(m[id], l.Latencies)
	}

	for id, lats := range m {
		ls := Latencies{}
		if len(lats) == 0 {
			fmt.Println(ls)
			continue
		}
		for i, l := range lats {
			j := i + 1
			ls.Proxy = avg(ls.Proxy, l.Proxy, j)
			ls.Gateway = avg(ls.Gateway, l.Gateway, j)
			ls.Request = avg(ls.Request, l.Request, j)
		}
		fmt.Printf("ID: %s:\n\t%v\n", id, ls)
	}
}

func avg(avg, new int64, i int) int64 {
	return avg + (new-avg)/int64(i)
}
