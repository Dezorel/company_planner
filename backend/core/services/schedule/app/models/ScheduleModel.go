package models

import (
	"errors"
	"strconv"
)

func GetScheduleByCompanyId(companyName string) ([]Schedule, error) {

	db := DBConnect()

	defer db.Close()

	query := "SELECT " +
		" cab.id, " +
		" cs.date_time_start, " +
		" cs.date_time_end, " +
		" cab.number " +
		" FROM " +
		" `Cabinets_schedule` cs " +
		" INNER JOIN Cabinets cab ON cs.cabinet_id = cab.id " +
		" WHERE " +
		" cab.company_id = (SELECT id FROM Companies WHERE company_name = '" + companyName + "')"

	selectResult, err := DBQueryMultiRow(db, query)

	if err != nil {
		return []Schedule{}, errors.New("Schedule not found!")
	}

	var answer []Schedule

	for selectResult.Next() {

		var id int
		var number, startDate, endDate string

		selectResult.Scan(&id, &startDate, &endDate, &number)

		answer = append(answer, Schedule{
			CabinetId:     strconv.Itoa(id),
			StartDate:     startDate,
			EndDate:       endDate,
			CabinetNumber: number,
			CompanyName:   companyName,
		})
	}

	selectResult.Close()

	Logger(3).Println("Schedule found: ", answer)

	return answer, nil
}

func CreateSchedule(startDate, endDate, cabinetId string) (Schedule, error) {

	db := DBConnect()

	defer db.Close()

	query := "INSERT INTO `Cabinets_schedule` (`cabinet_id`,`date_time_start`,`date_time_end`) " +
		"VALUES ( '" + cabinetId + "', '" + startDate + "', '" + endDate + "' )"

	resultQuery := DBQuery(db, query)

	if resultQuery == true {
		Logger(3).Println("Schedule successfully created. cabinetId: " + cabinetId +
			", startDate: " + startDate + ", endDate: " + endDate)

		return Schedule{
			CabinetId: cabinetId,
			StartDate: startDate,
			EndDate:   endDate,
		}, nil
	}

	return Schedule{}, errors.New("Error on create schedule!")
}
