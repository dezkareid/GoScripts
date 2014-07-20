package main

import "fmt" //Paquete para entrada y salida
import "net/http" //Paquete para hacer conexiones Http
import "encoding/json"
import "io/ioutil"



type Curso struct {
   	Nombre string `json:"curso"`
   	Escuela string `json:"escuela"`
}

type Cursos struct {
   	Cursos []Curso `json:"cursos"`
   
}


func main() {

	resp, err := http.Get("https://script.google.com/macros/s/AKfycbyx9FM2Athv6v0yuKzES1jZI7sCdqbY3mbtcnzrdnSblsmh0AQ/exec")
	contents, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(contents))
	var curso Cursos
	if err != nil {
	// handle error
	}
	if err := json.Unmarshal(contents, &curso); err != nil {
        panic(err)
    }
	
	
	
	
	fmt.Println(curso.Cursos[3].Nombre)

  	resp.Body.Close()
}