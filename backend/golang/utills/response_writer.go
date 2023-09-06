package utills

import (
	"encoding/json"
	"net/http"

	"github.com/taimoor99/three-tier-golang/app/entities"
)

const UserCreated = "user created"
const UserAlreadyExist = "user already exist"
const UserIdNotFoundInParam = "user id not not found param"
const LimitNotFoundInParam = "limit not not found param"
const OffsetNotFoundInParam = "offset not not found param"

func WriteJsonRes(w http.ResponseWriter, statusCode int, body interface{}, message string) {
	res := entities.JsonResponse{Message: message, Body: body}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(uj)
	return
}
