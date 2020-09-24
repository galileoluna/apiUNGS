package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExistInscripcion recibe un codigo de parámetro y chequea si ya está en la BD */
func ChequeoYaExisteInscripcion(codigo string) (models.Inscripcion, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ungs")
	col := db.Collection("Inscripcions")

	condicion := bson.M{"codigo": codigo}

	var resultado models.Inscripcion

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}