package router

import (
	"go_chi/internal/db"
	"go_chi/internal/feature/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestGetUserByID test user by ID
func TestGetUserByID(t *testing.T) {
	db := db.InitDatabase()
	defer db.Close()

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUser(db))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"address":"HCM","birthday":"2000-11-11T00:00:00Z","name":"Teo"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		desc           string
		givenUserID    string
		expectedResult interface{}
	}{
		{
			desc:        "Successfully",
			givenUserID: "1",
			expectedResult: user.User{
				ID:       1,
				Address:  "HCM",
				Birthday: "2000-11-11T00:00:00Z",
				Name:     "Teo",
			},
		},
	}

	for _, i := range testCases {
		db := db.InitDatabase()
		defer db.Close()
		// TODO
		t.Run(i.desc, func(t *testing.T) {

		})
	}
}
