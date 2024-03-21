package utility

type Role string

const (
	Coach  Role = "Coach"
	Normal Role = "Normal"
)

func GetRoleString(role Role) string {
	roleMap := map[Role]string{
		Coach:  "Coach",
		Normal: "Normal",
	}
	return roleMap[role]
}
