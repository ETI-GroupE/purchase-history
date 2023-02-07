package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type History struct {
	Order_id     int     `json:"order_id"`
	User_id      int     `json:"user_id"`
	Quantity     int     `json:"quantity"`
	Final_price  float64 `json:"final_price"`
	Product_id   int     `json:"product_id"`
	ShipStatus   string  `json:"shipstatus"`
	ShipLocation string  `json:"shiplocation"`
}

type product struct {
	Product_id          int    `json:"product_id"`
	Product_Name        string `json:"product_name"`
	Product_Description string `json:"product_description"`
}

type shopping_cart_items struct {
	Quantity int `json:"Quantity"`
}

type UserID struct {
	User_id int `json:"user_id"`
}

type Delivery struct {
	ShipLocation string `json:"shiplocation"`
	ShipStatus   string `json:"shipstatus"`
}

type OrderProducts struct {
	Product_Name        string `json:"product_name"`
	Product_Description string `json:"product_description"`
}

type CompositeKey struct {
	ProductID int `json:"product_id"`
	OrderID   int `json:"order_id"`
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/allpurchase", getAllPurchase).Methods("GET")
	router.HandleFunc("/api/v1/updatehistory", updatePurchaseHistory).Methods("POST")
	router.HandleFunc("/api/v1/viewAllBusinessPurchase", viewAllBusinessPurchase).Methods("GET")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}

func getAllPurchase(w http.ResponseWriter, r *http.Request) {
	//var products product        //WK
	//var purchasehistory History //B
	//var shoppingcart shopping_cart_items //LC
	//var status Status                    //H

	if r.Method == "GET" {
		querystringmap := r.URL.Query()
		userID := querystringmap.Get("UserID")

		// Execute when userID in query string is given a value
		if userID != "" {
			//Read
			ExodiaTheForbidden := os.Getenv("S1020")
			BodyOfExodia := os.Getenv("S8584")
			ArmsOfExodia := os.Getenv("S1090")
			LegsOfExodia := os.Getenv("S1019")
			//Calling of database
			db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)

			// Error handling
			if err != nil {
				fmt.Println("Error in connecting to database")
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			}
			defer db.Close()

			var purchasehi []History
			//Checking for value in database
			result, err := db.Query("select * from purchasehistory where user_id = ?", userID)
			if err != nil {
				fmt.Println("Error with getting data from database")
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)

			} else {
				for result.Next() {

					var purchasehistory History
					//Checking for database items
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.ShipStatus, &purchasehistory.ShipLocation)
					if err != nil {
						fmt.Printf("No purchase history available")
						http.Error(w, err.Error(), http.StatusBadRequest)
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, ("Invalid table in database"))

					} else {
						//Print out database items
						purchasehi = append(purchasehi, purchasehistory)
					}
				}
				output, _ := json.Marshal(purchasehi)
				w.WriteHeader(http.StatusAccepted)
				fmt.Fprintf(w, string(output))
			}

			// Execute when userID in query string is not given a value or not giving userID in query string
		} else {
			//Read
			ExodiaTheForbidden := os.Getenv("S1020")
			BodyOfExodia := os.Getenv("S8584")
			ArmsOfExodia := os.Getenv("S1090")
			LegsOfExodia := os.Getenv("S1019")
			//Calling of database
			db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)

			// Error handling
			if err != nil {
				fmt.Println("Error in connecting to database")
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
			}
			defer db.Close()

			var purchasehi []History
			//Checking for value in database
			result, err := db.Query("select * from purchasehistory ")
			if err != nil {
				fmt.Println("Error with getting data from database")
				http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)

			} else {
				for result.Next() {

					var purchasehistory History
					//Checking for database items
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.ShipStatus, &purchasehistory.ShipLocation)
					if err != nil {
						fmt.Printf("No purchase history available")
						http.Error(w, err.Error(), http.StatusBadRequest)
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, ("Invalid table in database"))

					} else {
						//Print out database items
						purchasehi = append(purchasehi, purchasehistory)
					}
				}
				output, _ := json.Marshal(purchasehi)
				w.WriteHeader(http.StatusAccepted)
				fmt.Fprintf(w, string(output))
			}
		}

	}
}

func updatePurchaseHistory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var purchasehistory History //B
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			if err := json.Unmarshal(body, &purchasehistory); err == nil {
				//Write
				ExodiaTheForbidden := os.Getenv("S1020")
				BodyOfExodia := os.Getenv("S8584")
				ArmsOfExodia := os.Getenv("S1029")
				LegsOfExodia := os.Getenv("S1019")
				//Calling of database
				db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)
				// handle error upon failure
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				defer db.Close()

				//inserting values into passenger table
				_, err = db.Exec("insert into purchasehistory (user_id, final_price, quantity, product_id, shipstatus,shiplocation) values(?,?,?,?,?,?)",
					purchasehistory.User_id, purchasehistory.Final_price, purchasehistory.Quantity, purchasehistory.Product_id, purchasehistory.ShipStatus, purchasehistory.ShipLocation)
				//Handling error of SQL statement
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}

}

func viewAllBusinessPurchase(w http.ResponseWriter, r *http.Request) {
	//var orderProductMap = make(map[int][]int)
	//var productInfo product     //WK
	//var userInfo UserID         //DE
	//https://buyee-delivery-qqglc24h2a-as.a.run.app //H Delivery
	//https://buyee-discount-qqglc24h2a-as.a.run.app //H Discount

	querystringmap := r.URL.Query()
	productID := querystringmap.Get("ProductID")
	if r.Method == "GET" {
		if productID != "" {
			//Read
			ExodiaTheForbidden := os.Getenv("S1020")
			BodyOfExodia := os.Getenv("S8584")
			ArmsOfExodia := os.Getenv("S1090")
			LegsOfExodia := os.Getenv("S1019")
			///Calling of database
			db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)

			// Error handling
			if err != nil {
				fmt.Println("Error in connecting to database")
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			defer db.Close()

			var purchasehi []History
			//Checking for value in database
			result, err := db.Query("select * from purchasehistory where product_id = ?", productID)
			if err != nil {
				fmt.Println("Error with getting data from database")
				http.Error(w, err.Error(), http.StatusBadRequest)

			} else {
				for result.Next() {

					var purchasehistory History
					//Checking for database items
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.ShipStatus, &purchasehistory.ShipLocation)
					if err != nil {
						fmt.Printf("No purchase history available")
						http.Error(w, err.Error(), http.StatusBadRequest)
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, ("Invalid table in database"))

					} else {
						//Print out database items
						purchasehi = append(purchasehi, purchasehistory)

					}
					output, _ := json.Marshal(purchasehi)
					w.WriteHeader(http.StatusAccepted)
					fmt.Fprintf(w, string(output))
				}

			}
		} else {
			//Read
			ExodiaTheForbidden := os.Getenv("S1020")
			BodyOfExodia := os.Getenv("S8584")
			ArmsOfExodia := os.Getenv("S1090")
			LegsOfExodia := os.Getenv("S1019")
			///Calling of database
			db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)

			// Error handling
			if err != nil {
				fmt.Println("Error in connecting to database")
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			defer db.Close()

			var purchasehi []History
			//Checking for value in database
			result, err := db.Query("select * from purchasehistory")
			if err != nil {
				fmt.Println("Error with getting data from database")
				http.Error(w, err.Error(), http.StatusBadRequest)

			} else {
				for result.Next() {

					var purchasehistory History
					//Checking for database items
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.ShipStatus, &purchasehistory.ShipLocation)
					if err != nil {
						fmt.Printf("No purchase history available")
						http.Error(w, err.Error(), http.StatusBadRequest)
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, ("Invalid table in database"))

					} else {
						//Print out database items
						purchasehi = append(purchasehi, purchasehistory)

					}
					output, _ := json.Marshal(purchasehi)
					w.WriteHeader(http.StatusAccepted)
					fmt.Fprintf(w, string(output))
				}

			}
		}

	}
}
