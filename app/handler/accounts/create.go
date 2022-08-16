package accounts

import (
	"encoding/json"
	"log"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // https://pkg.go.dev/context#Context

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}
	log.Default().Printf("req: %+v", req)

	// req: {Username:john Password:P@ssw0rd}
	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	db := h.app.Dao.Account() // domain/repository の取得
	if _, err := db.Insert(ctx, account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	// panic("Must Implement Account Registration")

	// account: &{
	// 	ID:0
	// 	Username:john
	// 	PasswordHash:$2a$10$swGUoUJfJUZnHAZRSzH6ZuNVpiQ2azG4dNprg4FhNrT6uvOBogcY6
	// 	DisplayName:<nil>
	// 	Avatar:<nil>
	// 	Header:<nil>
	// 	Note:<nil>
	// 	CreateAt:0001-01-01 00:00:00 +0000 UTC

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
