package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/johnbenedictyan/One-CV-Test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var testDb *gorm.DB

func TestMain(m *testing.M) {
	testDBName := "test"
	testDBUser := "admin"
	testDBPassword := "123"
	testDBHost := "localhost"
	testDBPort := "5433"
	testDBSslMode := "disable"
	testDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		testDBHost, testDBUser, testDBPassword, testDBName, testDBPort, testDBSslMode,
	)

	db, err := gorm.Open(postgres.Open(testDSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	testDb = db

	db.Exec("TRUNCATE TABLE students CASCADE")
	db.Exec("TRUNCATE TABLE teachers CASCADE")

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestCreateStudent(t *testing.T) {
	if testDb == nil {
		t.Errorf("Expected testDb to be initialized")
	}
	// Create Test Students
	testStudent1 := models.Student{
		Email: "testStudent1@email.com",
	}

	testDb.Create(&testStudent1)

	// Find Test Student
	var student models.Student
	testDb.First(&student, testStudent1.ID)

	// Assert Test Student
	assert.Equal(t, testStudent1.Email, student.Email)
}

func TestUpdateStudent(t *testing.T) {
	if testDb == nil {
		t.Errorf("Expected testDb to be initialized")
	}
	// Create Test Students
	testStudent2 := models.Student{
		Email: "testStudent2@email.com",
	}

	testDb.Create(&testStudent2)

	// Update Test Student
	testStudent2.Email = "newTestStudent2@email.com"
	testDb.Save(&testStudent2)

	// Find Test Student
	var student models.Student
	testDb.First(&student, testStudent2.ID)

	// Assert Test Student
	assert.Equal(t, testStudent2.Email, student.Email)
}

func TestDeleteStudent(t *testing.T) {
	if testDb == nil {
		t.Errorf("Expected testDb to be initialized")
	}
	// Create Test Students
	testStudent3 := models.Student{
		Email: "testStudent3@email.com",
	}

	testDb.Create(&testStudent3)

	// Delete Test Student
	testDb.Delete(&testStudent3)

	// Find Test Student
	var student models.Student
	testDb.First(&student, testStudent3.ID)

	// Assert Test Student
	assert.Equal(t, student.ID, uint(0))

}
