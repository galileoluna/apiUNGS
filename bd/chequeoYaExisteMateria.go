package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExistMateria recibe un codigo de parámetro y chequea si ya está en la BD */
func ChequeoYaExisteMateria(codigo string) (models.Materia, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ungs")
	col := db.Collection("materias")

	condicion := bson.M{"codigo": codigo}

	var resultado models.Materia

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}