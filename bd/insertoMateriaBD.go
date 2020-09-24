package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models" 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoMateriaBD es la parada final con la BD para insertar los datos del Materia */
func InsertoMateriaBD(u models.Materia) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ungs")
	col := db.Collection("materias")

	

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}