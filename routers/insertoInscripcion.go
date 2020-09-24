package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/galileoluna/apiUNGS/bd"
	"gitlab.com/galileoluna/apiUNGS/models"
)

/*InsertoInscripcion es la funcion para crear en la BD el registro de Inscripcion */
func InsertoInscripcion(w http.ResponseWriter, r *http.Request) {

	//creo un Inscripcion
	var t models.Inscripcion
	//El body es un Stream, es un dato que se lee una vez
	err := json.NewDecoder(r.Body).Decode(&t)
	 
	
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Alumno) == 0 {
		http.Error(w, "El Alumno  es requerido", 400)
		return
	}
	if len(t.Legajo) == 0 {
		http.Error(w, "El Legajo  es requerido", 400)
		return
	}
	if len(t.Materia) == 0 {
		http.Error(w, "La  Materia  es requerida", 400)
		return
	}
	if len(t.Codigo) == 0 {
		http.Error(w, "El codigo de la materia es requerido", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteInscripcion(t.Codigo)
	if encontrado == true {
		http.Error(w, "Ya existe una Inscripcion registrado con ese codigo", 400)
		return
	}

	_, status, err := bd.InsertoInscripcionBD(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de Inscripcion "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del Inscripcion", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}