package membershipPresentation

type Membership struct {
	ID             string
	UserName       string
	MembershipType string
}

var AvailableMemberships = [3]string{"toss", "naver", "payco"}
