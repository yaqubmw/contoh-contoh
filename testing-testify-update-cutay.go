func (suite *EmployeeRepositorySuite) TestUpdateMenstrualLeave() {
    employeeID := employeeDummy[0].ID
    availableDays := 0

    expectedQuery := `UPDATE "employee" SET "menstrual_leave"=\$1 WHERE id = \$2`

	suite.mocksql.ExpectBegin()
    suite.mocksql.ExpectExec(expectedQuery).WithArgs(availableDays, employeeID).WillReturnResult(sqlmock.NewResult(0, 1))
	suite.mocksql.ExpectCommit()

    err := suite.repo.UpdateMenstrualLeave(employeeID, availableDays)
    assert.NoError(suite.T(), err)
}
