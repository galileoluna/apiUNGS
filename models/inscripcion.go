package models

import "time"

/*GraboTweet es el formato o estructura que tendrá nuestro Tweet en la BD */
type Inscripcion struct {
	InscripcionID  string    `bson:"userid" json:"userid,omitempty"`
	Alumno         string    `bson:"alumno" json:"alumno,omitempty"`
	Legajo  	   string    `bson:"legajo" json:"legajo,omitempty"`
	Materia 	   string    `bson:"materia" json:"materia,omitempty"` 		
	Codigo 	  	   string    `bson:"codigo" json:"codigo,omitempty"` 		
	Fecha   	   time.Time `bson:"fecha" json:"fecha,omitempty"`
}