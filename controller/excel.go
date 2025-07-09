package controller

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func checkCouple(s string) string {
	if s == "" {
		return "Pas encore marié"
	}
	return s
}

func (c *Controller) Export() (string, error) {
	allMembers, err := c.R.FindAllMember()
	if err != nil {
		return "", fmt.Errorf("erreur récupération membres : %w", err)
	}

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

	for i, m := range allMembers {
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
			checkCouple(m.Couple),
			m.WeddingDate,
			m.Faritra,
		}

		for j, v := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, v)
		}
	}

	filePath := "./members.xlsx"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	if err := f.SaveAs(filePath); err != nil {
		return "", fmt.Errorf("erreur sauvegarde fichier : %w", err)
	}

	return filePath, nil
}
