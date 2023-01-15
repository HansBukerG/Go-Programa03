package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {

	/*Variables de entorno para pruebas*/

	movies = append(movies, Movie{ID: "1",Isbn: "438227",Title: "El Vilchez",Director: &Director {Firstname: "Hans",Lastname: "Vilchez"}})
	movies = append(movies, Movie{ID: "2",Isbn: "438228",Title: "El vilchez 2",Director: &Director{Firstname: "Hans",Lastname: "Vilchez"}})
	movies = append(movies, Movie{ID: "3",Isbn: "465654",Title: "El vilchez 3", Director: &Director{Firstname: "Hans",Lastname: "Vilchez"}})

	/*Fin variables de entorno para pruebas*/

	puerto := ":8000"

	fmt.Println("Programa 03 - Web server Peliculas")

	router := mux.NewRouter()

	router.HandleFunc("/movies",getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	router.HandleFunc("/movies",createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Iniciando Servidor en el puerto: "+puerto+"\n")

	log.Fatal(http.ListenAndServe(puerto,router))
}

func getMovies(writer http.ResponseWriter,request *http.Request) {
	writer.Header().Set("content-type","application/json")
	json.NewEncoder(writer).Encode(movies)
}
func getMovie(writer http.ResponseWriter,request *http.Request) {
	writer.Header().Set("content-type","application/json")
	params := mux.Vars(request)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter,request *http.Request) {
	writer.Header().Set("content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}
func updateMovie(writer http.ResponseWriter,request *http.Request) {
	writer.Header().Set("content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	for index, item := range movies{
		if item.ID == movie.ID{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movies)
}
func deleteMovie(writer http.ResponseWriter,request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}
