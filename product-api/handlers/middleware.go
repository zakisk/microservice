package handlers

import (
	"context"
	"net/http"

	"github.com/zakisk/microservice/product-api/data"
)

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		//validate the product we've got in request
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product")
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		//add prod to this context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		//passing control to the next middleware or if there no then to final handler
		next.ServeHTTP(rw, r)
	})
}