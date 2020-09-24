  
package bd

import (
	"context"
	"time"

	"gitlab.com/galileoluna/apiUNGS/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*getMaterias lee las materias*/
func GetMaterias( pagina int64) ([]*models.Materia, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("ungs")
	col := db.Collection("materias")

	var resultados []*models.Materia

	condicion := bson.M{
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.Materia
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}