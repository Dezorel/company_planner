package models

import (
	"errors"
	"strconv"
)

func GetCabinetsByCompany(companyName string) ([]Cabinet, error) {

	db := DBConnect()

	defer db.Close()

	query := "SELECT cab.id, cab.number, cab.size_id, cab.specifications " +
		"FROM Cabinets cab " +
		"INNER JOIN Companies comp ON cab.company_id = comp.id " +
		"WHERE company_name = '" + companyName + "'"

	selectResult, err := DBQueryMultiRow(db, query)

	if err != nil {
		return []Cabinet{}, errors.New("Cabinets not found!")
	}

	var answer []Cabinet

	for selectResult.Next() {

		var id, sizeId int
		var number, specification string

		selectResult.Scan(&id, &number, &sizeId, &specification)

		answer = append(answer, Cabinet{
			Id:       strconv.Itoa(id),
			Number:   number,
			Size:     strconv.Itoa(sizeId),
			Property: specification,
			Company:  companyName,
		})
	}

	selectResult.Close()

	return answer, nil
}

func GetCabinetInfoById(cabinetId string) (Cabinet, error) {
	db := DBConnect()

	defer db.Close()

	query := "SELECT cab.number, cab.size_id, cab.specifications, comp.company_name " +
		"FROM Cabinets cab " +
		"INNER JOIN Companies comp ON cab.company_id = comp.id " +
		"WHERE cab.id = '" + cabinetId + "'"

	selectResult, err := DBQueryMultiRow(db, query)

	if err != nil {
		return Cabinet{}, errors.New("Cabinets not found!")
	}

	var answerCabinet Cabinet

	for selectResult.Next() {

		var sizeId int
		var number, specification, companyName string

		selectResult.Scan(&number, &sizeId, &specification, &companyName)

		answerCabinet = Cabinet{
			Id:       cabinetId,
			Number:   number,
			Size:     strconv.Itoa(sizeId),
			Property: specification,
			Company:  companyName,
		}
	}

	selectResult.Close()

	return answerCabinet, nil
}

func CreateCabinet(companyName, cabinetNumber, cabinetSize, cabinetPropreties string) (Cabinet, error) {
	return Cabinet{
		Company:  companyName,
		Number:   cabinetNumber,
		Size:     cabinetSize,
		Property: cabinetPropreties,
	}, nil
}
