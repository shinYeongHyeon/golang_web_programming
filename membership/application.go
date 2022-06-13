package membership

import "errors"

type Application struct {
	repository Repository
}

var availableMemberships = [3]string{"toss", "naver", "payco"}

// NewApplication : Application struct 생성
func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

// Create : create for membership
func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	validateError := app.validateRequestParameters(request)
	if validateError != nil {
		return CreateResponse{}, validateError
	}

	_, err := app.repository.Find(request.UserName)
	if err == nil {
		return CreateResponse{}, errors.New("can not create for duplicate name")
	}

	membership := app.repository.Create(request.UserName, request.MembershipType)

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	return UpdateResponse{}, nil
}

func (app *Application) Delete(id string) error {
	return nil
}

func isAvailableMembershipType(membershipType string) bool {
	isAvailable := false
	for _, value := range availableMemberships {
		if value == membershipType {
			isAvailable = true
			break
		}
	}

	return isAvailable
}

func (app *Application) validateRequestParameters(request CreateRequest) error {
	if request.UserName == "" || request.MembershipType == "" {
		return errors.New("can not create for empty name or empty membershipType")
	}

	if !isAvailableMembershipType(request.MembershipType) {
		return errors.New("can not create for not available membershipType")
	}

	return nil
}
