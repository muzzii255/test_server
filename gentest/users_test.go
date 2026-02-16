package gentests

import (
	"net/http"
	"testing"

	"test_server/models"

	"github.com/stretchr/testify/require"
)

func testusers(t *testing.T) {
	app := setup()

	t.Run("Create users", func(t *testing.T) {
		payloads := []struct {
			name           string
			payload        models.User
			expectedStatus int
		}{
			{
				name:           "users 0",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Diana",
					Age:  65,
					Contact: models.ContactInfo{
						Email: "user1616@example.com",
						Phone: "+18757337380",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "51333",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "29963",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "28399",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "74041",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "12003",
						},
					},
				},
			},
			{
				name:           "users 1",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Charlie",
					Age:  49,
					Contact: models.ContactInfo{
						Email: "user5184@example.com",
						Phone: "+18755965441",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "95621",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "30623",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "89039",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "83454",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "56964",
						},
					},
				},
			},
			{
				name:           "users 2",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Diana",
					Age:  46,
					Contact: models.ContactInfo{
						Email: "user6057@example.com",
						Phone: "+18754351048",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "98630",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "94832",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "61050",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "24783",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "50845",
						},
					},
				},
			},
			{
				name:           "users 3",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Charlie",
					Age:  36,
					Contact: models.ContactInfo{
						Email: "user2861@example.com",
						Phone: "+18751478534",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "28156",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "38740",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "96329",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "83516",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "98736",
						},
					},
				},
			},
			{
				name:           "users 4",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Eve",
					Age:  19,
					Contact: models.ContactInfo{
						Email: "user8650@example.com",
						Phone: "+18758342817",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "30246",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "35276",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "76149",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "19318",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "41253",
						},
					},
				},
			},
			{
				name:           "users 5",
				expectedStatus: http.StatusCreated,
				payload: models.User{
					Name: "Eve",
					Age:  20,
					Contact: models.ContactInfo{
						Email: "user9444@example.com",
						Phone: "+18759341376",
					},
					Address: []models.Address{
						{
							Street:  "your mom blvd",
							Zipcode: "13447",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "81460",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "84815",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "58932",
						},
						{
							Street:  "your mom blvd",
							Zipcode: "61313",
						},
					},
				},
			},
		}
		for _, pl := range payloads {
			t.Run(pl.name, func(t *testing.T) {
				resp := makeReq(t, app, http.MethodPost, "/api/v1/users", pl.payload)
				require.Equal(t, http.StatusOK, resp.StatusCode)
			})
		}
	})

	t.Run("Update users", func(t *testing.T) {
		payloads := []struct {
			name           string
			payload        models.User
			expectedStatus int
		}{
			{name: "users 0", expectedStatus: http.StatusOK, payload: models.User{
				Name: "Alice",
				Age:  65,
				Contact: models.ContactInfo{
					Email: "user6634@example.com",
					Phone: "+18756392003",
				},
				Address: []models.Address{
					{
						Street:  "your mom blvd",
						Zipcode: "21793",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "72113",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "55721",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "13895",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "27917",
					},
				},
			}},
			{name: "users 1", expectedStatus: http.StatusOK, payload: models.User{
				Name: "Bob",
				Age:  30,
				Contact: models.ContactInfo{
					Email: "user5073@example.com",
					Phone: "+18752133459",
				},
				Address: []models.Address{
					{
						Street:  "your mom blvd",
						Zipcode: "75050",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "61972",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "19318",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "96214",
					},
					{
						Street:  "your mom blvd",
						Zipcode: "35170",
					},
				},
			}},
		}
		for _, pl := range payloads {
			t.Run(pl.name, func(t *testing.T) {
				resp := makeReq(t, app, http.MethodPut, "/api/v1/users", pl.payload)
				require.Equal(t, pl.expectedStatus, resp.StatusCode)
			})
		}
	})

	t.Run("Get users", func(t *testing.T) {
		testCases := []struct {
			name           string
			path           string
			expectedStatus int
		}{
			{name: "users 0", path: "/api/v1/users", expectedStatus: http.StatusOK},
			{name: "users 1", path: "/api/v1/users", expectedStatus: http.StatusOK},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				resp := makeReq(t, app, http.MethodGet, tc.path, nil)
				require.Equal(t, tc.expectedStatus, resp.StatusCode)
			})
		}
	})
}
