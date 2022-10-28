package commons

const (
	Admin    Roles         = "Admin"
	Normal   Roles         = "Normal"
	Disable  Roles         = "Disable"
	Male     Genders       = "Male"
	Female   Genders       = "Female"
	Pending  AccountStatus = "Password Pending"
	Active   AccountStatus = "Active"
	Inactive AccountStatus = "Inactive"
	Jan      Months        = "01"
	Feb      Months        = "02"
	Mar      Months        = "03"
	Apr      Months        = "04"
	May      Months        = "05"
	Jun      Months        = "06"
	Jul      Months        = "07"
	Aug      Months        = "08"
	Sep      Months        = "09"
	Oct      Months        = "10"
	Nov      Months        = "11"
	Dec      Months        = "12"
)

type (
	Roles         string
	Genders       string
	AccountStatus string
	Months        string
)
