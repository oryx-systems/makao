package gorm_test

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/oryx-systems/makao/pkg/makao/application/common/testutils"
	"github.com/oryx-systems/makao/pkg/makao/application/utils"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/db/gorm"
)

var (
	fixtures  *testfixtures.Loader
	testingDB *gorm.PGInstance
	db        *sql.DB

	// Pin variables
	salt, encryptedPin string
	userID             = "6ecbbc80-24c8-421a-9f1a-e14e12678ee0"
	testPhone          = "+254722000000"
	testIdentifier     = "123456789"
)

func TestMain(m *testing.M) {
	log.Printf("Setting tests up ...")

	isLocalDB := testutils.CheckIfCurrentDBIsLocal()
	if !isLocalDB {
		fmt.Println("Cannot run tests. The current database is not a local database.")
		os.Exit(1)
	}

	log.Println("setting up test database")
	var err error

	testingDB, err = gorm.NewPGInstance()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}
	db, err = testingDB.DB.DB()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}

	// setup test variables
	salt, encryptedPin = utils.EncryptPIN("0000", nil)

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Template(),
		testfixtures.TemplateData(template.FuncMap{
			"salt":                  salt,
			"hash":                  encryptedPin,
			"valid_to":              time.Now().Add(500).String(),
			"test_user_id":          userID,
			"test_phone":            "\"" + testPhone + "\"",
			"test_identifier_value": "\"" + testIdentifier + "\"",
		}),
		// this is the directory containing the YAML files.
		// The file name should be the same as the table name
		// order of inserting values matter to avoid foreign key constraint errors
		testfixtures.Paths(
			"../../../../../../fixtures/makao_user.yml",
			"../../../../../../fixtures/makao_contact.yml",
			"../../../../../../fixtures/makao_identifier.yml",
		),
		// uncomment when running tests locally, if your db is not a test db
		// Ensure the testing db in the ci is named `test`
		// !!Warning!!: this can corrupt data, do not turn on or run tests while in non-test db
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		fmt.Println("failed to create fixtures:", err)
		os.Exit(1)

	}

	err = prepareTestDatabase()
	if err != nil {
		fmt.Println("failed to prepare test database:", err)
		os.Exit(1)
	}

	log.Printf("Running tests ...")
	os.Exit(m.Run())
}

func prepareTestDatabase() error {
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}
