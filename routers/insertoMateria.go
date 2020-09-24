package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/galileoluna/apiUNGS/bd"
	"gitlab.com/galileoluna/apiUNGS/models"
)

/*InsertoMateria es la funcion para crear en la BD el registro de Materia */
func InsertoMateria(w http.ResponseWriter, r *http.Request) {

	//creo un Materia
	var t models.Materia
	//El body es un Stream, es un dato que se lee una vez
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Nombre) == 0 {
		http.Error(w, "El nombre de Materia es requerido", 400)
		return
	}
	if len(t.Codigo) == 0 {
		http.Error(w, "El codigo de Materia es requerido", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteMateria(t.Codigo)
	if encontrado == true {
		http.Error(w, "Ya existe una Materia registrado con ese codigo", 400)
		return
	}

	_, status, err := bd.InsertoMateriaBD(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de Materia "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del Materia", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}