# API UNGS
Esta API desarrollada en Golang nos ayuda a simular el sistema Guarani de la Universidad, nos brinda los alumnos, materias, e inscripciones.
**URL**: https://apiungs.herokuapp.com/


# Alumnos
## Insertar
Para poder generar un nuevo alumno hace falta generar los siguientes campos :

    {
		"nombre":"tomas",
		"apellidos":"gomez",
		"legajo":"12345678",
		"email":"tomasgo@gmail.com ",
		"fechaNacimiento":"1995-06-30T00:00:00Z",
		"password":"1234567"
		
    }
    
Para eso usamos PostMan o el generador de peticiones preferido, utilizando el siguiente request :  `/insertoAlumno`
## Leer

Para leer los alumnos:

https://apiungs.herokuapp.com/getAlumnos?pagina=1

# Materias
## Insertar
Para poder generar una nueva materia hace falta generar los siguientes campos:

    {
		"nombre":"tomas",
		"apellidos":"gomez"
    }
    
Para eso usamos PostMan o el generador de peticiones preferido, utilizando el siguiente request :  `/insertoMateria`
## Leer
Para leer las materias:

https://apiungs.herokuapp.com/getMaterias?pagina=1

## Inscripciones

Para poder generar una nueva inscripcion hace falta generar los siguientes campos:

	{
		"alumno":"galileo",
		"legajo":"1234567",
		"materia":"PP1",
		"codigo":"A101",
		"fecha":"2020-09-24T00:00:00Z"
	}

Para eso usamos PostMan o el generador de peticiones preferido, utilizando el siguiente request :  `/insertoInscripion`




## Leer
Las inscripciones son paginadas de  a 20 en orden cronológico, para acceder a esto nosotros podemos ir leyendo de a paginas.
https://apiungs.herokuapp.com/getInscripciones?pagina=1