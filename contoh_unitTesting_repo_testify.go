package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"employeeleave/model"
	"employeeleave/repository"
)

type StatusLeaveRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repository.StatusLeaveRepository
}

func (suite *StatusLeaveRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(suite.T(), err)

	err = db.AutoMigrate(&model.StatusLeave{})
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.repo = repository.NewStatusLeaveRepository(db)
}

func (suite *StatusLeaveRepositoryTestSuite) TearDownTest() {
	suite.db.Migrator().DropTable(&model.StatusLeave{})
	suite.db.Close()
}

func (suite *StatusLeaveRepositoryTestSuite) TestCreateAndGetStatusLeave_Success() {
	statusLeave := model.StatusLeave{
		ID:   "1",
		Name: "Approved",
	}

	err := suite.repo.Create(statusLeave)
	assert.NoError(suite.T(), err)

	fetchedStatusLeave, err := suite.repo.Get(statusLeave.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), statusLeave.Name, fetchedStatusLeave.Name)
}

func (suite *StatusLeaveRepositoryTestSuite) TestListStatusLeaves_Success() {
	// Insert dummy status leaves into the database for testing
	dummyStatusLeaves := []model.StatusLeave{
		{ID: "1", Name: "Approved"},
		{ID: "2", Name: "Pending"},
	}

	for _, statusLeave := range dummyStatusLeaves {
		err := suite.repo.Create(statusLeave)
		assert.NoError(suite.T(), err)
	}

	// Retrieve the list of status leaves
	statusLeaves, err := suite.repo.List()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), statusLeaves, len(dummyStatusLeaves))
}

func (suite *StatusLeaveRepositoryTestSuite) TestUpdateStatusLeave_Success() {
	statusLeave := model.StatusLeave{
		ID:   "1",
		Name: "Approved",
	}

	err := suite.repo.Create(statusLeave)
	assert.NoError(suite.T(), err)

	// Update the status leave
	updatedStatusLeave := model.StatusLeave{
		ID:   "1",
		Name: "Modified",
	}

	err = suite.repo.Update(updatedStatusLeave)
	assert.NoError(suite.T(), err)

	fetchedStatusLeave, err := suite.repo.Get(statusLeave.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedStatusLeave.Name, fetchedStatusLeave.Name)
}

func (suite *StatusLeaveRepositoryTestSuite) TestDeleteStatusLeave_Success() {
	statusLeave := model.StatusLeave{
		ID:   "1",
		Name: "Approved",
	}

	err := suite.repo.Create(statusLeave)
	assert.NoError(suite.T(), err)

	// Delete the status leave
	err = suite.repo.Delete(statusLeave.ID)
	assert.NoError(suite.T(), err)

	// Try to fetch the deleted status leave
	_, err = suite.repo.Get(statusLeave.ID)
	assert.Error(suite.T(), err)
}

func (suite *StatusLeaveRepositoryTestSuite) TestUpdateStatusLeave_Fail() {
	// Attempt to update a status leave that doesn't exist
	updatedStatusLeave := model.StatusLeave{
		ID:   "999",
		Name: "Modified",
	}

	err := suite.repo.Update(updatedStatusLeave)
	assert.Error(suite.T(), err)
}

func (suite *StatusLeaveRepositoryTestSuite) TestDeleteStatusLeave_Fail() {
	// Attempt to delete a status leave that doesn't exist
	err := suite.repo.Delete("999")
	assert.Error(suite.T(), err)
}

func TestStatusLeaveRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StatusLeaveRepositoryTestSuite))
}
