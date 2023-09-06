package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/taimoor99/three-tier-golang/app/entities"
	"github.com/taimoor99/three-tier-golang/utills"
)

type ModelsMock struct {
	mock.Mock
}

func (m *ModelsMock) GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error) {
	users := []entities.Users{entities.Users{Email: "test", Name: "test"}}
	return users, nil
}

func (m *ModelsMock) FindUserByID(ctx context.Context, id string) (entities.Users, error) {
	var user entities.Users
	user.ID = primitive.NewObjectID()
	return user, nil
}

func (m *ModelsMock) CreateUser(ctx context.Context, user entities.UsersCreateInput) (string, error) {
	return primitive.NewObjectID().Hex(), nil
}

var r *chi.Mux

func init() {
	r = chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *rest.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	ctrl := NewUser(new(ModelsMock))
	NewRouter(r, ctrl)
}

func TestUser_PostCreateUserHandler(t *testing.T) {
	body := entities.UsersCreateInput{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "test",
	}

	payload, _ := json.Marshal(body)
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/create-user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, response.Message, utills.UserCreated)
}

func TestUser_GetUserByIdHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/get-user/"+"test", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.NotEmpty(t, response.Body)
}

func TestUser_GetAllUsersHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/users/10/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.NotEmpty(t, response.Body)
}

