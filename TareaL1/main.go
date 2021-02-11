package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var data Datos

type Datos struct{
	Mensajes []Mensajes `json:"Mensajes"`
}

type Mensajes struct{
	Origen string `json:"Origen"`
	Destino string `json:"Destino"`
	Msg []Msg `json:"Msg"`
}

type Msg struct {
	Fecha string `json:"Fecha"`
	Texto string `json:"Texto"`
}

func Inicial(w http.ResponseWriter, r *http.Request){
	for i := 0; i < len(data.Mensajes); i++ {
		fmt.Println("Indice: " + data.Mensajes[i].Destino)
		fmt.Println("Indice: " + data.Mensajes[i].Origen)
		for j:=0; j<len(data.Mensajes[i].Msg);j++{
			fmt.Println("Departamento: " + data.Mensajes[i].Msg[j].Fecha)
			fmt.Println("Departamento: " + data.Mensajes[i].Msg[j].Texto)
		}
	}
	json.NewEncoder(w).Encode(data)
}

func mostrar(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(body, &data)

}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", Inicial).Methods("GET")
	router.HandleFunc("/", mostrar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000",router))
}
