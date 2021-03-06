/*
 * Proddx API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package router

import (
	"fmt"
	"net/http"

	"api.proddx.com/storage"
	"api.proddx.com/tokens"
	"github.com/julienschmidt/httprouter"
)

func New(us storage.User, cs storage.Company, ps storage.Product, rs storage.Review) *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/", Logger(Index(), "Index"))

	router.HandlerFunc(http.MethodOptions, "/login", cors)
	router.Handler(http.MethodPost, "/login", Logger(corsHandler(login(us)), "LoginUser"))
	router.HandlerFunc(http.MethodOptions, "/register", cors)
	router.Handler(http.MethodPost, "/register", Logger(corsHandler(register(us, cs)), "RegisterUser"))

	router.HandlerFunc(http.MethodOptions, "/companies", cors)
	router.HandlerFunc(http.MethodOptions, "/companies/:id", cors)
	router.Handler(http.MethodGet, "/companies", Logger(corsHandler(tokens.Validation(listCompanies(cs))), "ListCompanies"))
	router.Handler(http.MethodPost, "/companies", Logger(corsHandler(tokens.Validation(insertCompany(cs))), "InsertCompany"))
	router.Handler(http.MethodGet, "/companies/:id", Logger(corsHandler(tokens.Validation(findCompany(cs))), "FindCompany"))
	router.Handler(http.MethodPut, "/companies/:id", Logger(corsHandler(tokens.Validation(updateCompany(cs))), "UpdateCompany"))
	router.Handler(http.MethodDelete, "/companies/:id", Logger(corsHandler(tokens.Validation(deleteCompany(cs))), "DeleteCompany"))

	router.HandlerFunc(http.MethodOptions, "/products", cors)
	router.HandlerFunc(http.MethodOptions, "/products/:id", cors)
	router.Handler(http.MethodGet, "/products", Logger(corsHandler(tokens.Validation(listProducts(ps))), "ListProducts"))
	router.Handler(http.MethodPost, "/products", Logger(corsHandler(tokens.Validation(insertProduct(ps))), "InsertProduct"))
	router.Handler(http.MethodGet, "/products/:id", Logger(corsHandler(findProduct(ps)), "FindProduct"))
	router.Handler(http.MethodPut, "/products/:id", Logger(corsHandler(tokens.Validation(updateProduct(ps))), "UpdateProduct"))
	router.Handler(http.MethodDelete, "/products/:id", Logger(corsHandler(tokens.Validation(deleteProduct(ps))), "DeleteProduct"))

	router.HandlerFunc(http.MethodOptions, "/reviews", cors)
	router.HandlerFunc(http.MethodOptions, "/reviews/:id", cors)
	router.Handler(http.MethodGet, "/reviews", Logger(corsHandler(tokens.Validation(listReviews(rs))), "ListReviews"))
	router.Handler(http.MethodPost, "/reviews", Logger(corsHandler(insertReview(rs)), "InsertReview"))
	router.Handler(http.MethodGet, "/reviews/:id", Logger(corsHandler(tokens.Validation(findReview(rs))), "FindReview"))
	router.Handler(http.MethodPut, "/reviews/:id", Logger(corsHandler(tokens.Validation(updateReview(rs))), "UpdateReview"))
	router.Handler(http.MethodDelete, "/reviews/:id", Logger(corsHandler(tokens.Validation(deleteReview(rs))), "DeleteReview"))

	return router
}

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}
}
