package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Payment_Gateway struct{
    payment_gateway_id int `json:"payment_gateway_id"`
    name  string `json:"name"`
    status int `json:"status"`
}
var db *sql.DB

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "application/json",
    )

    rows, err := db.Query("SELECT * FROM payment_gateway")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    fmt.Println(rows)

    payment_gateway_list := []Payment_Gateway{}

    for rows.Next(){
        payment_gateway := Payment_Gateway{}
        err := rows.Scan(&payment_gateway.payment_gateway_id, &payment_gateway.name, &payment_gateway.status)
        if err != nil{
            fmt.Println(err)
            continue
        }
        payment_gateway_list = append(payment_gateway_list, payment_gateway)
    }

    data, err := json.Marshal(payment_gateway_list)

    io.WriteString(
        res,
		string(data),
    )
}

func bye(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    io.WriteString(
        res,
        `<doctype html>
			<html>
				<head>
					<title>Bye</title>
				</head>
				<body>
					Test!
				</body>
			</html>`,
    )
}
func main() {
	db, err := sql.Open("mysql", "root:@127.0.0.1:3306/hockey_payment")
	if err != nil {
        log.Println(err)
    }
    defer db.Close()

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/bye", bye)
    http.ListenAndServe(":9000", nil)
}
