package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoAlumnoBD es la parada final con la BD para insertar los datos del Alumno */
func InsertoAlumnoBD(u models.Alumno) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("alumnos")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}