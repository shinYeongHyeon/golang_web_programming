package membershipControllerDto

type MembershipUser struct {
	ID             string `json:"id"`
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}

type CreateRequest struct {
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}

type CreateResponse struct {
	Code           int             `json:"code"`
	Message        string          `json:"message"`
	MembershipUser *MembershipUser `json:"membershipUser,omitempty"`
}

type FindALlResponse struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	MembershipUsers []*MembershipUser `json:"membershipUsers,omitempty"`
}

type FindResponse struct {
	Code           int             `json:"code"`
	Message        string          `json:"message"`
	MembershipUser *MembershipUser `json:"membershipUser,omitempty"`
}

type UpdateRequest struct {
	UserName       string `json:"userName"`
	MembershipType string `json:"membershipType"`
}

type UpdateResponse struct {
	Code           int             `json:"code"`
	Message        string          `json:"message"`
	MembershipUser *MembershipUser `json:"membershipUser,omitempty"`
}

type DeleteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
