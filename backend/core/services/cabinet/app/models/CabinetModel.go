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

	Logger(3).Println("Cabinets found: ", answer)

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

	Logger(3).Println("Cabinet found: ", answerCabinet)

	return answerCabinet, nil
}

func CreateCabinet(companyName, cabinetNumber, cabinetSize, cabinetProperties string) (Cabinet, error) {

	db := DBConnect()

	defer db.Close()

	query := "INSERT INTO `Cabinets` (`number`,`company_id`,`specifications`,`size_id`) " +
		"VALUES (" +
		"'" + cabinetNumber + "'," +
		"( SELECT id FROM Companies WHERE company_name = '" + companyName + "' LIMIT 1)," +
		"'" + cabinetProperties + "'," +
		"( SELECT id FROM Cabinet_size WHERE cabinet_size >= '" + cabinetSize + "' ORDER BY cabinet_size ASC LIMIT 1)" +
		")"

	resultQuery := DBQuery(db, query)

	if resultQuery == true {

		Logger(3).Println("Cabinets successfully created. companyName: " +
			companyName + ", cabinetNumber: " + cabinetNumber + ", cabinetSize: " + cabinetSize +
			", cabinetProperties: " + cabinetProperties)

		return Cabinet{
			Company:  companyName,
			Number:   cabinetNumber,
			Size:     cabinetSize,
			Property: cabinetProperties,
		}, nil
	}

	return Cabinet{}, errors.New("Error on create cabinet!")
}
