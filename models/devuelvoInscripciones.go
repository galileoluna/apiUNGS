package models

import (
	

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweets es la estructura con la que devolveremos los Tweets */
type DevuelvoInscripciones struct {
	ID      	   primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Alumno         string    `bson:"alumno" json:"alumno,omitempty"`
	Legajo  	   string    `bson:"legajo" json:"legajo,omitempty"`
	Materia 	   string    `bson:"materia" json:"materia,omitempty"` 		
	Codigo 	  	   string    `bson:"codigo" json:"codigo,omitempty"` 
	
}