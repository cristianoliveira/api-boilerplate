package main

import (
  "os"
	"log"
	"net/http"
  "encoding/json"

	"github.com/gorilla/mux"
  "gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

var DATABASE string = "findhotel"

type Hotel struct {
  ID        string   `json:"id,omitempty"`
  Firstname string   `json:"firstname,omitempty"`
  Lastname  string   `json:"lastname,omitempty"`
}

func main() {
  // Create an index for Full Text Search
  // index := mgo.Index{
  //   Key: []string{"$text:name", "$text:about"},
  // }
  // err = hotels.EnsureIndex(index)
  // if err != nil { panic(err) }

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/hotels.json", GetHotelsHandler).Methods("GET")
  router.HandleFunc("/hotels", CreateHotelHandler).Methods("POST")
  router.HandleFunc("/hotels/{id}", GetHotelHandler).Methods("GET")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func GetHotelsHandler(w http.ResponseWriter, r *http.Request) {
  session, err := mgo.Dial(os.Getenv("database_PORT"))
  if err != nil { panic(err) }

  hotels := session.DB(DATABASE).C("hotels")

  var results []Hotel
  err = hotels.Find(nil).All(&results)
  if err != nil { panic(err) }

  hotels.Insert(hotel)

  json.NewEncoder(w).Encode(results)
}

func CreateHotelHandler(w http.ResponseWriter, r *http.Request) {
  session, err := mgo.Dial(os.Getenv("database_PORT"))
  if err != nil { panic(err) }

  hotels := session.DB(DATABASE).C("hotels")

  params := mux.Vars(r)

  var hotel Hotel
  err = json.NewDecoder(r.Body).Decode(&hotel)
  if err != nil { panic(err) }

  hotels.Insert(hotel)

  json.NewEncoder(w).Encode(people)
}
func GetHotelHandler(w http.ResponseWriter, r *http.Request) {
  session, err := mgo.Dial(os.Getenv("database_PORT"))
  if err != nil { panic(err) }

  hotels := session.DB(DATABASE).C("hotels")

  var results []Hotel
  err = hotels.Find(nil).All(&results)
  if err != nil { panic(err) }

  json.NewEncoder(w).Encode(results)
}
