package database

import (
	"alura-backend-7/src/models"
	"errors"
)

var MockDb = []models.Deposition{
	{
		Id:         1,
		Picture:    "https://pictures-repository.s3.aws.com/fabiano.png",
		Deposition: "Good product",
		Name:       "Fabiano",
	},
	{
		Id:         2,
		Picture:    "https://pictures-repository.s3.aws.com/maria.png",
		Deposition: "The product become with defects",
		Name:       "Maria",
	},
	{
		Id:         3,
		Picture:    "https://pictures-repository.s3.aws.com/bruno.png",
		Deposition: "Good product! I recommend this seller",
		Name:       "Bruno",
	},
	{
		Id:         4,
		Picture:    "https://pictures-repository.s3.aws.com/pedro.png",
		Deposition: "Bad delivery service, the shipment delayed 3 days",
		Name:       "Pedro",
	},
	{
		Id:         5,
		Picture:    "https://pictures-repository.s3.aws.com/alice.png",
		Deposition: "Nothing compares to that quality and technology of this product",
		Name:       "Alice",
	},
}

func CreateNewDeposition(deposition models.Deposition) models.Deposition {
	MockDb = append(MockDb, deposition)
	return deposition
}

func GetAllDepositions() []models.Deposition {
	return MockDb
}

func GetDepositionById(id int) models.Deposition {
	for _, deposition := range MockDb {
		if deposition.Id == id {
			return deposition
		}
	}
	return models.Deposition{}
}

func UpdateDeposition(id int, deposition models.Deposition) error {
	for i, d := range MockDb {
		if d.Id == id {
			MockDb[i] = deposition
			return nil
		}
	}

	return errors.New("deposition not found")
}

func DeleteDeposition(id int) error {
	for i, d := range MockDb {
		if d.Id == id {
			MockDb = append(MockDb[:i], MockDb[i+1:]...)
			return nil
		}
	}

	return errors.New("deposition not found")
}
