package res

import (
	"encoding/json"
	"net/http"
)

type Json map[string]interface{}

type Result struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

func GetBody(r *http.Request) (Json, error) {
	var result Json
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return Json{}, err
	}
	defer r.Body.Close()

	return result, nil
}

func WriteResponse(w http.ResponseWriter, statusCode int, e []byte) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(e)

	return nil
}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, dt interface{}) error {
	if dt == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}

	result := Result{
		Ok:     http.StatusOK == statusCode,
		Result: dt,
	}

	e, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return WriteResponse(w, statusCode, e)
}
