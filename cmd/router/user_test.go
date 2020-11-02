package router_test

import (
	"go_chi/internal/db"
	"go_chi/internal/feature"
	"go_chi/internal/feature/user"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

//TestGetUserByID test user by ID
func TestGetUser(t *testing.T) {
	testCases := []struct {
		desc           string
		givenUserID    string
		expectedResult user.User
		expectedError  *feature.ResponseError
	}{
		{
			desc:        "Successfully",
			givenUserID: "2",
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

			id, _ := strconv.Atoi(i.givenUserID)
			result, err := user.GetUser(db, id)
			if err != nil {
				require.Equal(t, i.expectedError, err)
			} else {
				require.Nil(t, err)
				require.Equal(t, i.expectedResult, result)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		desc          string
		givenUser     user.User
		expectedError *feature.ResponseError
	}{
		{
			desc: "Successfully",

			givenUser: user.User{

				Address:  "Dong Thap",
				Birthday: "1996-01-23T00:00:00Z",
				Name:     "Hong",
			},
		},
	}
	for _, i := range testCases {
		db := db.InitDatabase()
		defer db.Close()

		t.Run(i.desc, func(t *testing.T) {
			err := user.AddUser(db, &i.givenUser)
			if err != nil {
				require.Equal(t, i.expectedError, err)
			}
			require.Nil(t, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	testCases := []struct {
		desc          string
		givenUserID   string
		givenUser     user.User
		expectedError *feature.ResponseError
	}{
		{
			desc:        "Successfully",
			givenUserID: "38",
			givenUser: user.User{

				Address:  "Dong Thap",
				Birthday: "1996-01-23T00:00:00Z",
				Name:     "Super",
			},
		},
	}
	for _, i := range testCases {
		db := db.InitDatabase()
		defer db.Close()
		id, _ := strconv.Atoi(i.givenUserID)

		err := user.UpdateUser(db, &i.givenUser, id)
		if err != nil {
			require.Equal(t, i.expectedError, err)
		}
		require.Nil(t, err)
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		desc          string
		givenUserID   string
		expectedError *feature.ResponseError
	}{
		{
			desc:        "Successfully",
			givenUserID: "38",
		},
	}
	for _, i := range testCases {
		db := db.InitDatabase()
		defer db.Close()
		// TODO
		t.Run(i.desc, func(t *testing.T) {
			id, _ := strconv.Atoi(i.givenUserID)
			err := user.DeleteUser(db, id)
			if err != nil {
				require.Equal(t, i.expectedError, err)
			}
			require.Nil(t, err)

		})
	}
}
