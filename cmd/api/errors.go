package main

import (
	"fmt"
	"net/http"
)

// logError 함수: 에러 로깅 헬퍼
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// errorResponse 함수: 에러 제네릭 헬퍼
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse 함수: 서버 에러 담당
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and couyld not process your request"

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse 함수: 404 담당
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse 함수: 405 담당
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
