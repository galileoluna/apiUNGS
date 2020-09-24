package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Alumno es el modelo de Alumno de la base de MongoDB */
type Alumno struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	Legajo          string             `bson:"legajo" json:"legajo,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
}