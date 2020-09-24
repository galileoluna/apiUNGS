package routers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/galileoluna/apiUNGS/bd"
	"gitlab.com/galileoluna/apiUNGS/models"
)

/*InsertoAlumno es la funcion para crear en la BD el registro de Alumno */
func InsertoAlumno(w http.ResponseWriter, r *http.Request) {

	//creo un alumno
	var t models.Alumno
	//El body es un Stream, es un dato que se lee una vez
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Nombre) == 0 {
		http.Error(w, "El nombre de Alumno es requerido", 400)
		return
	}
	if len(t.Apellidos) == 0 {
		http.Error(w, "El Apellido de Alumno es requerido", 400)
		return
	}
	if len(t.Legajo) == 0 {
		http.Error(w, "El Legajo de Alumno es requerido", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de Alumno es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser mayor a  6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteAlumno(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un Alumno registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoAlumnoBD(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de Alumno "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del Alumno", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}