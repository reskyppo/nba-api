package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// init Global Variable
var db *gorm.DB
var err error

// Struct for Teams data details
type Teams struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

// Struct for Results response API
type Results struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Struct for list endpoint Response API
type Endpoints struct {
	GetAll    string `json:"get_all"`
	GetById   string `json:"get_by_id"`
	SaveTeams string `json:"save_teams"`
}

// Func to handle content from home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint for homepage")
}

// Func to handle post request
func handleNewTeams(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var teams Teams
	json.Unmarshal(reqBody, &teams)
	db.Create(&teams)
	fmt.Println("Endpoint for post a new teams")
	json.NewEncoder(w).Encode(teams)
}

// Func to get all Teams data
func handleTeams(w http.ResponseWriter, r *http.Request) {
	teams := []Teams{}
	db.Find(&teams)
	fmt.Println("Endpoint for get all teams")
	json.NewEncoder(w).Encode(teams)
}

// Get teams by id
func getTeamsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	teams := []Teams{}
	db.Find(&teams)
	for _, teams := range teams {
		// string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if teams.Id == s {
				fmt.Println("Endpoint for teams with id:", key)
				json.NewEncoder(w).Encode(teams)
			}
		}
	}
}

// Func to handle request and create router
func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:8008/")
	log.Println("Quit the server with CONTROL-C.")

	// create a router using mux
	myRouter := mux.NewRouter().StrictSlash(true)
	//endpoint "/" and the content is from func homePage
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/team/save", handleNewTeams).Methods("POST")
	myRouter.HandleFunc("/team", handleTeams)
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
	db.AutoMigrate(&Teams{})

	//initiate func for request
	handleRequests()

}
