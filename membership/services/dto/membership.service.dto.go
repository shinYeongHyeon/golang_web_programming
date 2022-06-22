package membershipServicesDto

type MembershipUser struct {
	ID             string
	UserName       string
	MembershipType string
}

type CreateRequest struct {
	UserName       string
	MembershipType string
}

type CreateResponse struct {
	Code           string
	MembershipUser *MembershipUser
}

type FindAllRequest struct {
	Offset string
	Limit  string
}

type FindAllResponse struct {
	MembershipUsers []*MembershipUser
}

type FindOneResponse struct {
	Code           string
	MembershipUser *MembershipUser
}

type UpdateRequest struct {
	ID             string
	UserName       string
	MembershipType string
}

type UpdateResponse struct {
	Code           string
	MembershipUser *MembershipUser
}

type DeleteResponse struct {
	Code string
}
