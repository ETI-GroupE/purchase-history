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
	Order_id    int     `json:"order_id"`
	User_id     int     `json:"user_id"`
	Quantity    int     `json:"quantity"`
	Final_price float64 `json:"final_price"`
	Product_id  int     `json:"product_id"`
	Status      string  `json:"status"`
	Location    string  `json:"location"`
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
	shipLocation string `json:"location"`
	shipStatus   string `json:"status"`
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
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.Status, &purchasehistory.Location)
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
					err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.Status, &purchasehistory.Location)
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
	var productInfo product              //WK
	var purchasehistory History          //B
	var shoppingcart shopping_cart_items //LC
	var DeliveryInfo Delivery            //H
	//var userInfo UserID //DE

	if r.Method == "POST" {
		//querystringmap := r.URL.Query()
		//userID := querystringmap.Get("UserID")

		//=====================================
		//Calling shopping cart endpoint
		response, err := http.Get("https://buyee-shoppingcart-gukqthlh4a-as.a.run.app/api/v1/shoppingCart")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of Shopping cart :", err)
			return
		}

		err = json.Unmarshal(body, &shoppingcart)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of Shopping cart:", err)
			return
		}

		//=====================================
		//Calling product endpoint
		response, err = http.Get(" https://buyee-delivery-qqglc24h2a-as.a.run.app/api/v1")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of product :", err)
			return
		}

		err = json.Unmarshal(body, &productInfo)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of product:", err)
			return
		}

		//=====================================
		//Calling Delivery endpoint
		response, err = http.Get("https://buyee-delivery-qqglc24h2a-as.a.run.app/api/v1/status")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of Delivery :", err)
			return
		}

		err = json.Unmarshal(body, &DeliveryInfo)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of Delivery:", err)
			return
		}

		//Write
		ExodiaTheForbidden := os.Getenv("S1020")
		BodyOfExodia := os.Getenv("S8584")
		ArmsOfExodia := os.Getenv("S1029")
		LegsOfExodia := os.Getenv("S1019")
		//Calling of database
		db, err := sql.Open("mysql", ExodiaTheForbidden+":"+BodyOfExodia+"@tcp("+ArmsOfExodia+")/"+LegsOfExodia)

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer db.Close()

		//Inserting values into database
		_, err = db.Exec("insert into purchasehistory (user_id, final_price,quantity,product_id,status,location) values(?,?,?,?,?,?)",
			purchasehistory.User_id, purchasehistory.Final_price, purchasehistory.Quantity, purchasehistory.Product_id, purchasehistory.Status, purchasehistory.Location)
		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		} else {
			// To notify of new item to purchase history
			fmt.Println("====================")
			fmt.Println("New purchase history added")
		}
	}
}

func viewAllBusinessPurchase(w http.ResponseWriter, r *http.Request) {
	//var orderProductMap = make(map[int][]int)
	var productInfo product     //WK
	var purchasehistory History //B
	var userInfo UserID         //DE
	//https://buyee-delivery-qqglc24h2a-as.a.run.app //H Delivery
	//https://buyee-discount-qqglc24h2a-as.a.run.app //H Discount

	querystringmap := r.URL.Query()
	productID := querystringmap.Get("ProductID")
	if productID != "" {

	} else {

	}
	if r.Method == "GET" {
		//=====================================
		//Calling business endpoint
		response, err := http.Get("https://auth-ksbujg5hza-as.a.run.app/api/v1/verify/customer")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of User :", err)
			return
		}

		err = json.Unmarshal(body, &userInfo)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of User:", err)
			return
		}

		//=====================================
		//Calling product endpoint
		response, err = http.Get("https://buyee-catalog-ksbujg5hza-as.a.run.app/api/v1")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of product :", err)
			return
		}

		err = json.Unmarshal(body, &productInfo)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of product:", err)
			return
		}

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

		//Checking for value in database
		result, err := db.Query("select * from purchasehistory where product_id = ?", productID)
		if err != nil {
			fmt.Println("Error with getting data from database")
			http.Error(w, err.Error(), http.StatusBadRequest)

		} else {
			for result.Next() {

				//Checking for database items
				err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.Status, &purchasehistory.Location)
				if err != nil {
					fmt.Printf("No purchase history available")
					http.Error(w, err.Error(), http.StatusBadRequest)
					//w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(w, "No purchase history available")

				} else {
					w.WriteHeader(http.StatusOK)
					output, _ := json.Marshal(purchasehistory)
					w.WriteHeader(http.StatusAccepted)
					fmt.Fprintf(w, string(output))
					fmt.Println(purchasehistory.Order_id, purchasehistory.Final_price, purchasehistory.Quantity, purchasehistory.Status, purchasehistory.Location)
					// LATEST COMPOSITE KEY CODE
					// m := make(map[CompositeKey]product)
					// m[CompositeKey{ProductID: productInfo.Product_id, OrderID: purchasehistory.Order_id}] = product{Product_Name: productInfo.Product_Name, Product_Description: productInfo.Product_Description}
					// for key, value := range m {
					// 	fmt.Fprintf(w, "Product ID: %d Order ID: %d Quantity: %s Price: %s", key.ProductID, key.OrderID, value.Product_Name, value.Product_Description)
					// }

					// OLD COMPOSITE KEY CODE NO USE
					// var orderProductMap = make(map[int]map[int]product)
					// orderProductMap[purchasehistory.Order_id] = make(map[int]product)
					// orderProductMap[purchasehistory.Order_id][purchasehistory.Product_id] = product{Product_Name: productInfo.Product_Name, Product_Description: productInfo.Product_Description}
					// fmt.Println(orderProductMap)
					// fmt.Fprintln(w, "Status OK")
					// output, err := json.Marshal(orderProductMap)
					// if err != nil {
					// 	http.Error(w, err.Error(), http.StatusInternalServerError)
					// 	return
					// }
					// w.Header().Set("Content-Type", "application/json")
					// w.Write(output)
					//w.WriteHeader(http.StatusAccepted)
				} //fmt.Fprintf(w, string(output))

			}

		}

	}
}
