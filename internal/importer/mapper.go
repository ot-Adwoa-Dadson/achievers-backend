package importer

import (
	"time"

	"fellowship-backend/internal/models"
)

func CsvToMember(csv MemberCSV) models.Member {
	dob, _ := time.Parse("2006-01-02", csv.DateOfBirth)

	return models.Member{
		FullName:    csv.FullName,
		Phone:       csv.Phone,
		Email:       csv.Email,
		HomeAddress: csv.HomeAddress,
		DateOfBirth: dob,

		Occupation:      csv.Occupation,
		CurrentEmployer: csv.CurrentEmployer,

		Guardian: models.Guardian{
			Name:             csv.GuardianName,
			Relationship:     csv.GuardianRelationship,
			Phone:            csv.GuardianPhone,
			AlternativePhone: csv.GuardianAlternativePhone,
			Email:            csv.GuardianEmail,
		},

		Fellowship: models.FellowshipInfo{
			SeniorCell:             csv.SeniorCell,
			FoundationSchoolStatus: csv.FoundationSchoolStatus,
			LeadershipRole:         csv.LeadershipRole,
			DesignationCell:        csv.DesignationCell,
		},

		IsNewMember: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
