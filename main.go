package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	ty "github.com/reskyppo/nba-api/types"
)

// init Global Variable
var db *gorm.DB
var err error

// Func to handle content from home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for homepage")
	var data = ty.Endpoints{
		GetAll:      "/team",
		GetById:     "/team/{id}",
		SaveTeams:   "/team/save",
		UpdateTeams: "/team/save/{id}",
		DeleteTeams: "/team/delete/{id}",
	}

	res := ty.Results{Code: http.StatusOK, Data: data, Message: "Created by Reskyppo"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// Func to handle post request
func addTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for post a new teams")
	reqBody, _ := ioutil.ReadAll(r.Body)

	var teams ty.Teams
	json.Unmarshal(reqBody, &teams)
	db.Create(&teams)

	res := ty.Results{Code: http.StatusOK, Data: teams, Message: "Success create new Teams"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// Func to get all Teams data
func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for get all teams")
	teams := []ty.Teams{}
	db.Find(&teams)
	res := ty.Results{Code: http.StatusOK, Message: "Sukses get all data", Data: teams}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// Get teams by id
func getTeamsById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for get team by id")

	vars := mux.Vars(r)
	key := vars["id"]
	teams := []ty.Teams{}

	// handle error data not found
	if db.First(&teams, "id = ?", key).RowsAffected != 0 {
		res := ty.Results{Code: http.StatusOK, Data: teams, Message: "Success get Team data"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		res := ty.Results{Code: http.StatusNotFound, Data: teams, Message: "Id Team not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
	}
}

// func to update db teams
func updateTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for update team by id")

	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var teamsUpdate ty.Teams
	json.Unmarshal(reqBody, &teamsUpdate)

	var teams ty.Teams
	// handle error data not found
	if db.First(&teams, "id = ?", key).RowsAffected != 0 {
		db.Model(&teams).Updates(teamsUpdate)

		res := ty.Results{Code: http.StatusOK, Data: teams, Message: "Success update teams"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		data := []ty.Teams{}
		res := ty.Results{Code: http.StatusNotFound, Data: data, Message: "Id Team not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
	}
}

// func to delete db teams
func deleteTeams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for delete team by id")

	vars := mux.Vars(r)
	key := vars["id"]

	var teams ty.Teams
	data := []ty.Teams{} //to make it empty array rather than null

	// handle error data not found
	if db.First(&teams, "id = ?", key).RowsAffected != 0 {
		db.Delete(&teams)

		res := ty.Results{Code: http.StatusOK, Data: data, Message: "Success delete team"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		res := ty.Results{Code: http.StatusNotFound, Data: data, Message: "Id Team not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
	}
}

// Func to handle request and create router
func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8008/")
	log.Println("Quit the server with CONTROL-C.")

	// create a router using mux
	myRouter := mux.NewRouter().StrictSlash(true)

	// handle wrong endpoint
	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		data := []ty.Results{}
		res := ty.Results{Code: http.StatusNotFound, Data: data, Message: "Enpoint not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	//endpoint "/" and the content is from func homePage
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/team/save", addTeams).Methods("POST")
	myRouter.HandleFunc("/team/save/{id}", updateTeams).Methods("PUT")
	myRouter.HandleFunc("/team/delete/{id}", deleteTeams).Methods("DELETE")
	myRouter.HandleFunc("/team", getAll)
	myRouter.HandleFunc("/team/{id}", getTeamsById)
	log.Fatal(http.ListenAndServe(":8008", myRouter)) //set the port and handler
}

func main() {
	// Define DB name, DB username, DB password
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/Basketball?charset=utf8&parseTime=True")

	// Check if it's connect to DB or not
	if err != nil {
		log.Println("Connection Failed to Open because", err)
	} else {
		log.Println("Connection Established")
	}

	// Create Table Teams
	db.AutoMigrate(&ty.Teams{})

	//initiate func for request
	handleRequests()

}
