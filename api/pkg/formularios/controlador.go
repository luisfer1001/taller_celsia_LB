package formularios

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func insertFormulario(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	nombre := body["nombre"].(string)
	apellidos := body["apellidos"].(string)
	tipoIdentificacion := body["tipoIdentificacion"].(string)
	estadoCivilStr := body["estadoCivil"].(string)
	fechaNacimiento := body["fechaNacimiento"].(string)
	numBeneficiariosStr := body["numBeneficiarios"].(string)
	fechaIngreso := body["fechaIngreso"].(string)

	estadoCivil, err := strconv.Atoi(estadoCivilStr)
	if err != nil {
		estadoCivil = 0
	}

	numBeneficiarios, err := strconv.Atoi(numBeneficiariosStr)
	if err != nil {
		numBeneficiarios = 0
	}

	//log.Println(estadoCivil)
	//log.Println(numBeneficiarios)

	result, err := InsertFormulario(nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)

	// Verificar si se produjo un error durante la inserción de datos
	if err != nil {
		// Si hay un error, responder al cliente con un código de estado HTTP 500 (Internal Server Error) y un mensaje de error
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)

}

func selectAllFormularios(w http.ResponseWriter, r *http.Request) {
	result, err := SelectAllFormularios()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}

func updateFormulario(w http.ResponseWriter, r *http.Request) {
	numIdentificacionStr := chi.URLParam(r, "numIdentificacion")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		numIdentificacion = 0
	}

	body, _ := res.GetBody(r)
	nombre := body["nombre"].(string)
	apellidos := body["apellidos"].(string)
	tipoIdentificacion := body["tipoIdentificacion"].(string)
	estadoCivilStr := body["estadoCivil"].(string)
	fechaNacimiento := body["fechaNacimiento"].(string)
	numBeneficiariosStr := body["numBeneficiarios"].(string)
	fechaIngreso := body["fechaIngreso"].(string)

	estadoCivil, err := strconv.Atoi(estadoCivilStr)
	if err != nil {
		estadoCivil = 0
	}

	numBeneficiarios, err := strconv.Atoi(numBeneficiariosStr)
	if err != nil {
		numBeneficiarios = 0
	}

	log.Println(numIdentificacion)

	result, err := UpdateFormulario(numIdentificacion, nombre, apellidos, tipoIdentificacion, estadoCivil, fechaNacimiento, numBeneficiarios, fechaIngreso)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}

func deleteFormulario(w http.ResponseWriter, r *http.Request) {
	numIdentificacionStr := chi.URLParam(r, "numIdentificacion")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		numIdentificacion = 0
	}

	result, err := DeleteFormulario(numIdentificacion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}
