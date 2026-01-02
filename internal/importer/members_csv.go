package importer

type MemberCSV struct {
	FullName                 string `csv:"fullName"`
	Phone                    string `csv:"phone"`
	Email                    string `csv:"email"`
	HomeAddress              string `csv:"homeAddress"`
	DateOfBirth              string `csv:"dateOfBirth"`
	Occupation               string `csv:"occupation"`
	CurrentEmployer          string `csv:"currentEmployer"`

	GuardianName             string `csv:"guardianName"`
	GuardianRelationship     string `csv:"guardianRelationship"`
	GuardianPhone            string `csv:"guardianPhone"`
	GuardianAlternativePhone string `csv:"guardianAlternativePhone"`
	GuardianEmail            string `csv:"guardianEmail"`

	SeniorCell               string `csv:"seniorCell"`
	FoundationSchoolStatus   string `csv:"foundationSchoolStatus"`
	LeadershipRole           string `csv:"leadershipRole"` // STRING
	DesignationCell          string `csv:"designationCell"`
}
