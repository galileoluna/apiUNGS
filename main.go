package main


import (
	"log"
	"gitlab.com/galileoluna/apiUNGS/bd"
	"gitlab.com/galileoluna/apiUNGS/handlers"
	

)

func main()  {
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a la BD")
		return 
	}
	handlers.Manejadores()
}