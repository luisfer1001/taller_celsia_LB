package imagenes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func insertImagenes(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	nombreImagen := body["nombreImagen"].(string)
	fecha := body["fecha"].(string)
	numIdentificacionStr := body["numIdentificacion"].(string)

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		numIdentificacion = 0
	}

	result, err := InsertImagen(numIdentificacion, nombreImagen, fecha)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)

}

func selectAllImagenes(w http.ResponseWriter, r *http.Request) {
	result, err := SelectAllImagen()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}

func updateImagen(w http.ResponseWriter, r *http.Request) {
	numIdentificacionStr := chi.URLParam(r, "numIdentificacion")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		numIdentificacion = 0
	}

	nombreImagen := chi.URLParam(r, "nombreImagen")

	body, _ := res.GetBody(r)
	fecha := body["fecha"].(string)

	result, err := UpdateImagen(numIdentificacion, nombreImagen, fecha)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}

func deleteImagen(w http.ResponseWriter, r *http.Request) {
	numIdentificacionStr := chi.URLParam(r, "numIdentificacion")

	numIdentificacion, err := strconv.Atoi(numIdentificacionStr)
	if err != nil {
		numIdentificacion = 0
	}

	nombreImagen := chi.URLParam(r, "nombreImagen")

	result, err := DeleteImagen(numIdentificacion, nombreImagen)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{
			"message": err.Error(),
		})
		return
	}

	res.JSON(w, r, http.StatusOK, result)
}
