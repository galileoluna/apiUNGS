package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gitlab.com/galileoluna/apiUNGS/middlew"
	"gitlab.com/galileoluna/apiUNGS/routers"
	"github.com/rs/cors"
)


/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	router := mux.NewRouter()


	router.HandleFunc("/insertoAlumno", middlew.ChequeoBD(routers.InsertoAlumno)).Methods("POST")

//	router.HandleFunc("/insertoAlumno", middlew.ChequeoBD(routers.InsertoAlumno)).Methods("POST")

//	router.HandleFunc("/insertoAlumno", middlew.ChequeoBD(routers.InsertoAlumno)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//DAMOS Permiso a todo el mundo
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}