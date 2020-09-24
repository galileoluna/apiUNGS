package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExistAlumno recibe un email de parámetro y chequea si ya está en la BD */
func ChequeoYaExisteAlumno(email string) (models.Alumno, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apiungs")
	col := db.Collection("alumnos")

	condicion := bson.M{"email": email}

	var resultado models.Alumno

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}