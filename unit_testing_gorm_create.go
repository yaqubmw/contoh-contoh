package repository

import (
	"employeeleave/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StatusLeaveRepositorySuite struct {
	suite.Suite
	repo    StatusLeaveRepository
	mockDB  *gorm.DB
	mocksql sqlmock.Sqlmock
}

func (suite *StatusLeaveRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	suite.mockDB = gormDB
	suite.mocksql = mock
	suite.repo = NewStatusLeaveRepository(gormDB)
}

var dataDummy = []model.StatusLeave{
		{
			ID:              "1",
			StatusLeaveName: "Pending",
		},
		{
			ID:              "2",
			StatusLeaveName: "Approved",
		},
	}

func (suite *StatusLeaveRepositorySuite) TestCreate() {
	payload := dataDummy[0]

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"status_leave\" (.+)").WithArgs(payload.ID, payload.StatusLeaveName).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveRepositorySuite) TestGet() {
	statusLeaveID := "1"
	expectedStatusLeave := dataDummy[0]

	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
	rows.AddRow(expectedStatusLeave.ID, expectedStatusLeave.StatusLeaveName)
	expectedQuery := `SELECT \* FROM "status_leave" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(statusLeaveID).WillReturnRows(rows)

	result, err := suite.repo.Get(statusLeaveID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatusLeave, result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func (suite *StatusLeaveRepositorySuite) TestList() {
	expectedStatusLeaves := dataDummy

	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
	for _, statusLeave := range expectedStatusLeaves {
		rows.AddRow(statusLeave.ID, statusLeave.StatusLeaveName)
	}

	expectedQuery := `SELECT \* FROM "status_leave"`
	suite.mocksql.ExpectQuery(expectedQuery).WillReturnRows(rows)

	result, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatusLeaves, result)

	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}

func TestStatusLeaveRepositorySuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveRepositorySuite))
}
