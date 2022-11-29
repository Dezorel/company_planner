package models

import (
	"errors"
)

func GetCompanyByName(companyName string) (Company, error) {

	db := DBConnect()

	defer db.Close()

	query := "SELECT 1 as result FROM `Companies` WHERE company_name = '" + companyName + "'"

	resultQuery := DBQueryRow(db, query)

	if resultQuery.Result == "1" {
		return Company{Name: companyName}, nil
	}

	return Company{}, errors.New("Not found company!")
}

func CreateCompany(companyName string) (Company, error) {

	db := DBConnect()

	defer db.Close()

	query := "INSERT INTO `Companies` (`company_name`) VALUES ('" + companyName + "')"

	resultQuery := DBQuery(db, query)

	if resultQuery == true {
		return Company{Name: companyName}, nil
	}

	return Company{}, errors.New("Error on create company!")
}
