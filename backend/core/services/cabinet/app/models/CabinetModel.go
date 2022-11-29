package models

func GetCabinetsByCompany(companyName string) (Cabinet, error) {
	return Cabinet{Company: companyName}, nil
}

func GetCabinetInfoById(cabinetId string) (Cabinet, error) {
	return Cabinet{Id: cabinetId}, nil
}

func CreateCabinet(companyName, cabinetNumber, cabinetSize, cabinetPropreties string) (Cabinet, error) {
	return Cabinet{
		Company:  companyName,
		Number:   cabinetNumber,
		Size:     cabinetSize,
		Property: cabinetPropreties,
	}, nil
}
