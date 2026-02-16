package gentests
	import (
			"net/http"
			"net/http/httptest"
			"testing"
			"github.com/stretchr/testify/require"
			)

	func testusers(t *testing.T){


app := setup()

t.Run("Create users", func(t *testing.T) {payloads := []struct{
name string
payload models.User
expectedStatus int
}{{
				name: "users 0",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Bob",
Age: 40,
Contact: models.ContactInfo{
Email: "user9168@example.com",
Phone: "+18754541090",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "81854",
},
    {
Street: "your mom blvd",
Zipcode: "36428",
},
    {
Street: "your mom blvd",
Zipcode: "46840",
},
    {
Street: "your mom blvd",
Zipcode: "82233",
},
    {
Street: "your mom blvd",
Zipcode: "15316",
},
},
},},
{
				name: "users 1",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Eve",
Age: 24,
Contact: models.ContactInfo{
Email: "user8663@example.com",
Phone: "+18759399505",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "73794",
},
    {
Street: "your mom blvd",
Zipcode: "19607",
},
    {
Street: "your mom blvd",
Zipcode: "47845",
},
    {
Street: "your mom blvd",
Zipcode: "67795",
},
    {
Street: "your mom blvd",
Zipcode: "16914",
},
},
},},
{
				name: "users 2",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Diana",
Age: 24,
Contact: models.ContactInfo{
Email: "user8187@example.com",
Phone: "+18753682096",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "11562",
},
    {
Street: "your mom blvd",
Zipcode: "98297",
},
    {
Street: "your mom blvd",
Zipcode: "75368",
},
    {
Street: "your mom blvd",
Zipcode: "14834",
},
    {
Street: "your mom blvd",
Zipcode: "15045",
},
},
},},
{
				name: "users 3",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Diana",
Age: 39,
Contact: models.ContactInfo{
Email: "user3297@example.com",
Phone: "+18756157184",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "83795",
},
    {
Street: "your mom blvd",
Zipcode: "27538",
},
    {
Street: "your mom blvd",
Zipcode: "65075",
},
    {
Street: "your mom blvd",
Zipcode: "99183",
},
    {
Street: "your mom blvd",
Zipcode: "42207",
},
},
},},
{
				name: "users 4",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Bob",
Age: 18,
Contact: models.ContactInfo{
Email: "user4669@example.com",
Phone: "+18751181828",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "10473",
},
    {
Street: "your mom blvd",
Zipcode: "39800",
},
    {
Street: "your mom blvd",
Zipcode: "30726",
},
    {
Street: "your mom blvd",
Zipcode: "59993",
},
    {
Street: "your mom blvd",
Zipcode: "79436",
},
},
},},
{
				name: "users 5",
				expectedStatus: http.StatusCreated,
				payload: models.User {
Name: "Bob",
Age: 18,
Contact: models.ContactInfo{
Email: "user7180@example.com",
Phone: "+18755252745",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "97073",
},
    {
Street: "your mom blvd",
Zipcode: "37075",
},
    {
Street: "your mom blvd",
Zipcode: "52955",
},
    {
Street: "your mom blvd",
Zipcode: "79714",
},
    {
Street: "your mom blvd",
Zipcode: "72938",
},
},
},},
}
for _, pl := range payloads {
t.Run(pl.name, func(t *testing.T) {
resp := makeReq(t, app, http.MethodPost, "/api/v1/users", pl.payload)
require.Equal(t, http.StatusOK, resp.StatusCode)
})
}
})

t.Run("Update users", func(t *testing.T) {payloads := []struct{
name string
payload models.User
expectedStatus int
}{
{name: "users 0",expectedStatus: http.StatusOK, payload: models.User {
Name: "Eve",
Age: 44,
Contact: models.ContactInfo{
Email: "user8391@example.com",
Phone: "+18751333228",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "87387",
},
    {
Street: "your mom blvd",
Zipcode: "31623",
},
    {
Street: "your mom blvd",
Zipcode: "77189",
},
    {
Street: "your mom blvd",
Zipcode: "69757",
},
    {
Street: "your mom blvd",
Zipcode: "41001",
},
},
},},
{name: "users 1",expectedStatus: http.StatusOK, payload: models.User {
Name: "Alice",
Age: 27,
Contact: models.ContactInfo{
Email: "user9179@example.com",
Phone: "+18751338161",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "59946",
},
    {
Street: "your mom blvd",
Zipcode: "86918",
},
    {
Street: "your mom blvd",
Zipcode: "54989",
},
    {
Street: "your mom blvd",
Zipcode: "67985",
},
    {
Street: "your mom blvd",
Zipcode: "75396",
},
},
},},
{name: "users 2",expectedStatus: http.StatusOK, payload: models.User {
Name: "Alice",
Age: 54,
Contact: models.ContactInfo{
Email: "user7153@example.com",
Phone: "+18754573960",
},
Address: []models.Address{
    {
Street: "your mom blvd",
Zipcode: "70574",
},
    {
Street: "your mom blvd",
Zipcode: "11030",
},
    {
Street: "your mom blvd",
Zipcode: "64567",
},
    {
Street: "your mom blvd",
Zipcode: "65890",
},
    {
Street: "your mom blvd",
Zipcode: "42505",
},
},
},},
}
for _, pl := range payloads {
t.Run(pl.name, func(t *testing.T) {
resp := makeReq(t, app, http.MethodPut, "/api/v1/users", pl.payload)
require.Equal(t, pl.expectedStatus, resp.StatusCode)
})
}
})



t.Run("Delete users", func(t *testing.T) {resp := makeReq(t, app, http.MethodDelete, "/api/v1/users", nil)
require.Equal(t, http.StatusOK, resp.StatusCode)
})

}

