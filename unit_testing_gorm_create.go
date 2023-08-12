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

func (suite *StatusLeaveRepositorySuite) TestCreate_Success() {
	payload := model.StatusLeave{
		ID:              "1",
		StatusLeaveName: "Pending",
	}

	suite.mocksql.ExpectBegin()
	suite.mocksql.ExpectExec("INSERT INTO \"status_leave\" (.+)").WithArgs(payload.ID, payload.StatusLeaveName).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mocksql.ExpectCommit()

	err := suite.repo.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *StatusLeaveRepositorySuite) TestGet_Success() {
	statusLeaveID := "1"
	expectedStatusLeave := model.StatusLeave{
		ID:              statusLeaveID,
		StatusLeaveName: "Pending",
	}
	rows := sqlmock.NewRows([]string{"id", "status_leave_name"})
	rows.AddRow(expectedStatusLeave.ID, expectedStatusLeave.StatusLeaveName)
	expectedQuery := `SELECT \* FROM "status_leave" WHERE id = \$1`
	suite.mocksql.ExpectQuery(expectedQuery).WithArgs(statusLeaveID).WillReturnRows(rows)

	result, err := suite.repo.Get(statusLeaveID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStatusLeave, result)

	// Ensure all expectations were met
	assert.NoError(suite.T(), suite.mocksql.ExpectationsWereMet())
}


func TestStatusLeaveRepositorySuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveRepositorySuite))
}
