package imagenes

import "github.com/go-chi/chi/v5"

func Router(r *chi.Mux) {
	r.Post("/imagenes", insertImagenes)
	r.Get("/imagenes", selectAllImagenes)
	r.Put("/imagenes/{numIdentificacion}/{nombreImagen}", updateImagen)
	r.Delete("/imagenes/{numIdentificacion}/{nombreImagen}", deleteImagen)
}
