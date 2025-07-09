package service

import (
	"SI/model"
	"time"

	"github.com/xuri/excelize/v2"
)

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

func (s *Service) CreateExcel(input []model.Member) bool {
	f := excelize.NewFile()
	sheet := "Membres"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{
		"ID", "Prénom", "Nom", "Sexe", "Date de naissance", "Date de décès",
		"Baptême", "Date Baptême", "Renouvellement Baptême", "Date Renouvellement",
		"Confession", "Date Confession", "Communion", "Date Communion",
		"Confirmation", "Date Confirmation", "Mariage", "Couple", "Date Mariage", "Faritra",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, m := range input {
		row := i + 2
		values := []interface{}{
			m.Id,
			m.FirstName,
			m.LastName,
			m.Sexe,
			m.BirthDate,
			m.DeathDate,
			m.Baptism,
			m.BaptismDate,
			m.BaptismRenew,
			m.BaptismRenewDate,
			m.Confession,
			m.ConfessionDate,
			m.Communion,
			m.CommunionDate,
			m.Confirmation,
			m.ConfirmationDate,
			m.Wedding,
			m.WeddingDate,
			m.Couple,
			m.Faritra,
		}

		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
	}

	return true
}
