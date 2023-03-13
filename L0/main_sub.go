package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/btcsuite/btcd/btcutil"

	_ "github.com/lib/pq"
	_ "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	_ "github.com/nats-io/stan.go"
)

type From_db struct {
	id    int
	value string
}

type jsonModel struct {
	Order_uid    string `json:"order_uid"`
	Track_number string `json:"track_number"`
	Entry        string `json:"entry"`
	Delivery     struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction   string `json:"transaction"`
		Request_id    string `json:"request_id"`
		Currency      string `json:"currency"`
		Provider      string `json:"provider"`
		Amount        int    `json:"amount"`
		Payment_dt    int    `json:"payment_dt"`
		Bank          string `json:"bank"`
		Delivery_cost int    `json:"delivery_cost"`
		Goods_total   int    `json:"goods_total"`
		Custom_fee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		Chrt_id      int    `json:"chrt_id"`
		Track_number string `json:"track_number"`
		Price        int    `json:"price"`
		Rid          string `json:"rid"`
		Name         string `json:"name"`
		Sale         int    `json:"sale"`
		Size         string `json:"size"`
		Total_price  int    `json:"total_price"`
		Nm_id        int    `json:"nm_id"`
		Brand        string `json:"brand"`
		Status       int    `json:"status"`
	} `json:"items"`
	Locale             string `json:"locale"`
	Internal_signature string `json:"internal_signature"`
	Customer_id        string `json:"customer_id"`
	Delivery_service   string `json:"delivery_service"`
	Shardkey           string `json:"shardkey"`
	Sm_id              int    `json:"sm_id"`
	Date_created       string `json:"date_created"`
	Oof_shard          string `json:"oof_shard"`
}

var myCash = make(map[int]jsonModel, 10)

func main() {
	clusterID := "prod"
	clientID := "subscribe_client"
	channel := "future"
	id := 0
	value := ""
	var jSon = jsonModel{}
	var sub stan.Subscription

	//	POSTGRES
	connSTR := "user=andy password=cGBRYrd23 dbname=wildberries sslmode=disable"
	db, err := sql.Open("postgres", connSTR)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Evriiiica!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM order_id")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &value)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(value), &jSon)
		if err != nil {
			panic(err)
		}
		myCash[id] = jSon
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//	POSTGRESS

	//	NATS-STREAMING
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println("Hey sir! I cant connect(..")
	}

	go func() {
		sub, _ = sc.Subscribe(channel, func(m *stan.Msg) {
			model := jsonModel{}
			err = json.Unmarshal(m.Data, &model)
			if err != nil {
				fmt.Println("Incorrect json!")
				panic(err)
			}
			if model.Oof_shard == "" {
				fmt.Println("Struct is empty!")
			} else {
				fmt.Println(model)
				id++
				myCash[id] = model
				_, err = db.Exec("INSERT INTO order_id (value) VALUES ($1)", m.Data)
				fmt.Println("Value is correct")
				if err != nil {
					panic(err)
				}
			}
		}, stan.DeliverAllAvailable())
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "index.html")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "id parameter is missing", http.StatusBadRequest)
			return
		}
		idInInt, _ := strconv.Atoi(id)
		val, ok := myCash[idInInt]
		if ok {
			jsonData, err := json.Marshal(val)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
		http.Error(w, "order_id not found", http.StatusNotFound)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	sub.Unsubscribe()
	sc.Close()
	//	NATS-STREAMING
}
