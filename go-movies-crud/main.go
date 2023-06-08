package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"fristname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //sets a new value for an existing header inside a Headers object
	json.NewEncoder(w).Encode(movies)                  //w에 movies를 json 형태로 encoding
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // create a mapp of route variables
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //w에 item 를 json 형태로 encode
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)    //JSON encoded된 r.body를 변수 Movie type movie 변수로 저장함
	movie.ID = strconv.Itoa(rand.Intn(100000000)) //id를 랜덤 생성 후 string으로 변환 후 movie.ID에 저장
	movies = append(movies, movie)                //created된 movie를 movies에 append
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // create a mapp of route variables
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //해당 index에 해당하는 데이터를 빼고 movies에 반환
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //w에 item 를 json 형태로 encode
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r) // create a mapp of route variables
	//loop over the movies, range
	//delte the movie with the id that you've sent
	//add anew movie- the movie that we send in the body of postman
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) //JSON encoded된 r.body를 변수 Movie type movie 변수로 저장함
			movie.ID = params["id"]                    //기존 id 값을 읽어와 새로운 ID에 덮어 씌움
			movies = append(movies, movie)             //기본에 업데이트될 data 빼놓은 movies에 updated 된 movie를 추가
			json.NewEncoder(w).Encode(movie)           //업데이트 된movie를 w에 encoding
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})
	fmt.Println(movies)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("Delete")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
