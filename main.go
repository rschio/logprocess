package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		jsonlog := sc.Bytes()
		l := Log{}
		json.Unmarshal(jsonlog, &l)
		fmt.Println(l)
	}
	if err := sc.Err(); err != nil {
		log.Println(err)
	}
}
