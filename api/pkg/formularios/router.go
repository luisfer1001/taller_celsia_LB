package formularios

import "github.com/go-chi/chi/v5"

func Router(r *chi.Mux) {
	r.Post("/formularios", insertFormulario)
	r.Get("/formularios", selectAllFormularios)
	r.Put("/formularios/{numIdentificacion}", updateFormulario)
	r.Delete("/formularios/{numIdentificacion}", deleteFormulario)
}
