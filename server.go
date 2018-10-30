package main

import (
	"net/http"
	"io"
	//"fmt"
	"encoding/json"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Payment_Gateway struct {
    Payment_gateway_id int `json:"payment_gateway_id"`
    Name  string `json:"name"`
    Status int `json:"status"`
}
var database *sql.DB

func hello(res http.ResponseWriter, req *http.Request) {
    rows, err := database.Query("SELECT * FROM hockey_payment.payment_gateway")
    if err != nil {
		http.Error(res, http.StatusText(500), 500)
        return
    }
    defer rows.Close()

    var payment_gateway_list = []Payment_Gateway{}

    for rows.Next(){
        p := Payment_Gateway{}
        err := rows.Scan(&p.Payment_gateway_id, &p.Name, &p.Status)
        if err != nil{
            continue
        }
        payment_gateway_list = append(payment_gateway_list, p)
    }

    data, err := json.Marshal(payment_gateway_list)

	res.Header().Set(
        "Content-Type",
        "application/json",
    )

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
	db, err := sql.Open("mysql", "root@/hockey_payment")
	if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    database = db

	log.Println("Run server")

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/bye", bye)
    log.Fatal(http.ListenAndServe(":9000", nil))
}
