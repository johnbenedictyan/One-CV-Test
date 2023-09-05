//go:build e2e
// +build e2e

package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/johnbenedictyan/One-CV-Test/infra/database"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

type RegisterSuite struct {
	suite.Suite
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
}

func (s *RegisterSuite) SetupTest() {
	database.GetTestDB().Exec("TRUNCATE teacher_students")
	database.GetTestDB().Exec("TRUNCATE students")
	database.GetTestDB().Exec("TRUNCATE teachers")
}

func (s *RegisterSuite) TestRegisterStudentsSuccessfully() {
	c := http.Client{}

	// Create Test Teacher
	testTeacher := models.Teacher{
		Email: "testTeacher@email.com"
	}

	// Create Test Students
	testStudent1 := models.Student{
		Email: "testStudent1@email.com"
	}

	testStudent2 := models.Student{
		Email: "testStudent2@email.com"
	}


	database.GetTestDB().model(&models.Teacher).Create(&testTeacher)
	database.GetTestDB().model(&models.Student).Create(&testStudent1)
	database.GetTestDB().model(&models.Student).Create(&testStudent2)


	testRegisterPostData := {
		"teacher": testTeacher.Email,
		"students": [testStudent1.Email, testStudent2.Email]
	}

	r, _ := c.Post("/register", "application/json", testRegisterPostData)
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)
	s.JSONEq(`{"code": "000", "msg": "Success"}`, string(body))
}

func (s *RegisterSuite) TestGetBookThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(`{"code": "001", "msg": "No book with ISBN 123456789"}`, string(body))
}

func (s *RegisterSuite) TestGetBookWithInvalidISBNGiven() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/1234C6789")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(`{"code": "003", "msg": "ISBN is invalid"}`, string(body))
}

func (s *RegisterSuite) TestGetBookThatDoesExist() {
	s.T().Skip("Pact Demo")
	db.Exec("INSERT INTO book (isbn, name, image, genre, year_published) VALUES ('987654321', 'Testing All The Things', 'testing.jpg', 'Computing', 2021)")

	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/987654321")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)

	expBody := `{
	"isbn": "987654321",
	"title": "Testing All The Things",
	"image": "testing.jpg",
	"genre": "Computing",
	"year_published": 2021
}`

	s.JSONEq(expBody, string(body))
}
