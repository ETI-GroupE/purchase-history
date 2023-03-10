package Api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type History struct {
	//Item_id     int     `json:"item_id"`
	Order_id    int     `json:"order_id"`
	User_id     int     `json:"user_id"`
	Discount_id int     `json:"discount_id"`
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
	Quantity    int     `json:"Quantity"`
	Final_price float32 `json:"final_price"`
}

// type User struct {
// 	User_id int `json:"user_id"`
// }

type Status struct {
	Location string `json:"location"`
	Status   string `json:"status"`
}

type OrderProducts struct {
	Product_Name        string `json:"product_name"`
	Product_Description string `json:"product_description"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/allpurchase", getAllPurchase).Methods("GET")
	router.HandleFunc("/api/v1/updatehistory", updatePurchaseHistory).Methods("POST")
	router.HandleFunc("/api/v1/viewAllBusinessPurchase", viewAllBusinessPurchase).Methods("GET")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(": 5000", router))
}

func getAllPurchase(w http.ResponseWriter, r *http.Request) {
	var products product                 //WK
	var purchasehistory History          //B
	var shoppingcart shopping_cart_items //LC
	var status Status                    //H

	if r.Method == "GET" {

		//=====================================
		//Calling Products endpoint
		response, err := http.Get("https://localhost:5000/api/v1/products")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of Products :", err)
			return
		}

		err = json.Unmarshal(body, &products)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of Products:", err)
			return
		}

		//=====================================
		//Calling Shopping cart endpoint
		response, err = http.Get("https://localhost:5000/api/v1/shopping_cart_items")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of shopping cart:", err)
			return
		}

		err = json.Unmarshal(body, &shoppingcart)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of shopping cart:", err)
			return
		}

		//=====================================
		//Calling location endpoint
		response, err = http.Get("https://localhost:6327/discounts/{discount_id}/{shop_Cart_ID}")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of location:", err)
			return
		}

		err = json.Unmarshal(body, &status)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of location:", err)
			return
		}

		//Calling of database
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/eti_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
		}
		defer db.Close()

		//Checking for value in database
		result, err := db.Query("select * from purchasehistory where user_id = ?", purchasehistory.User_id)
		if err != nil {
			fmt.Println("Error with getting data from database")
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)

		} else {
			for result.Next() {

				//Checking for database items
				err = result.Scan(&purchasehistory.Order_id, &purchasehistory.User_id, &purchasehistory.Final_price, &purchasehistory.Quantity, &purchasehistory.Product_id, &purchasehistory.Status, &purchasehistory.Location)
				if err != nil {
					fmt.Printf("No purchase history available")
					http.Error(w, err.Error(), http.StatusBadRequest)
					w.WriteHeader(http.StatusBadRequest)

				} else {
					//Print out database items
					w.WriteHeader(http.StatusOK)
					fmt.Println(purchasehistory.Order_id, products.Product_Name, purchasehistory.Final_price, purchasehistory.Quantity, products.Product_Description, purchasehistory.Status, purchasehistory.Location)
				}
			}
		}
	}
}

func updatePurchaseHistory(w http.ResponseWriter, r *http.Request) {
	var products product                 //WK
	var purchasehistory History          //B
	var shoppingcart shopping_cart_items //LC
	var status Status                    //H

	if r.Method == "POST" {

		//=====================================
		//Calling Products endpoint
		response, err := http.Get("https://localhost:5000/api/v1/products")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of Products :", err)
			return
		}

		err = json.Unmarshal(body, &products)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of Products:", err)
			return
		}

		//=====================================
		//Calling Shopping cart endpoint
		response, err = http.Get("https://localhost:5000/api/v1/shopping_cart_items")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of shopping cart:", err)
			return
		}

		err = json.Unmarshal(body, &shoppingcart)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of shopping cart:", err)
			return
		}

		//=====================================
		//Calling location endpoint
		response, err = http.Get("https://localhost:6327/discounts/{discount_id}/{shop_Cart_ID}")
		if err != nil {
			fmt.Println("Error making the API call:", err)
			return
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading the response body of location:", err)
			return
		}

		err = json.Unmarshal(body, &status)
		if err != nil {
			fmt.Println("Error unmarshaling the JSON data of location:", err)
			return
		}

		//Testing only
		//fmt.Printf("post hello")

		//Calling of database
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/eti_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer db.Close()

		//Inserting values into database
		_, err = db.Exec("insert into purchasehistory (user_id, final_price,quantity,product_id,status,location) values(?,?,?,?,?,?)",
			purchasehistory.User_id, shoppingcart.Final_price, shoppingcart.Quantity, products.Product_id, status.Status, status.Location)
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
	var products product //WK
	var orderProduct OrderProducts
	var purchasehistory History //B

	// Map to store the products

	//=====================================
	//Calling Products endpoint
	response, err := http.Get("https://localhost:5000/api/v1/products")
	if err != nil {
		fmt.Println("Error making the API call:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body of Products :", err)
		return
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Error unmarshaling the JSON data of Products:", err)
		return
	}

	if r.Method == "GET" {
		///Calling of database
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/eti_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer db.Close()

		//Checking for value in database
		result, err := db.Query("select * from purchasehistory where product_id = 3")
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
					var orderProductMap = make(map[int]map[int]OrderProducts)
					orderProductMap[purchasehistory.Order_id] = make(map[int]OrderProducts)
					orderProductMap[purchasehistory.Order_id][purchasehistory.Product_id] = OrderProducts{Product_Name: orderProduct.Product_Name, Product_Description: orderProduct.Product_Description}
					fmt.Println(orderProductMap)
					fmt.Fprintln(w, "Status OK")
					// orderProductMap[purchasehistory.Order_id] = append(orderProductMap[purchasehistory.Order_id], purchasehistory.Product_id)
					// var orderProductArray []OrderProducts
					// for order_id, product_id := range orderProductMap {
					// 	orderProductArray = append(orderProductArray, OrderProducts{Order_id: order_id, Product_id: product_id})
					// 	fmt.Println(orderProductArray)
					// }
					// //Print out database items
					// //w.WriteHeader(http.StatusOK)

					//To display out something
					//fmt.Println(purchasehistory.Order_id, purchasehistory.User_id, purchasehistory.Final_price, purchasehistory.Quantity, purchasehistory.Product_id, purchasehistory.Status, purchasehistory.Location)
				}
			}

		}

	}
}
