package membership

import (
	"errors"
)

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
	validateError := app.validateCreateRequestParameters(request)
	if validateError != nil {
		return CreateResponse{}, validateError
	}

	_, err := app.repository.FindByName(request.UserName)
	if err == nil {
		return CreateResponse{}, errors.New("can not create for duplicate name")
	}

	membership := app.repository.Create(request.UserName, request.MembershipType)

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

// Update : update for membership
func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	requestParameterErr := app.validateUpdateRequestParameters(request)
	if requestParameterErr != nil {
		return UpdateResponse{}, requestParameterErr
	}

	requestUserNameDuplicationErr := app.validateUpdateRequestDuplicateUserName(request)
	if requestUserNameDuplicationErr != nil {
		return UpdateResponse{}, requestUserNameDuplicationErr
	}

	membership := app.repository.Update(request.ID, request.UserName, request.MembershipType)

	if membership.ID == "" {
		return UpdateResponse{}, errors.New("can not update for not exist member id")
	}

	return UpdateResponse{membership.ID, membership.UserName, membership.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	if id == "" {
		return errors.New("can not delete for empty member id")
	}

	membership := app.repository.Find(id)

	if membership.ID == "" {
		return errors.New("can not delete for not exist id")
	}

	app.repository.Delete(id)

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

func (app *Application) validateCreateRequestParameters(request CreateRequest) error {
	if request.UserName == "" || request.MembershipType == "" {
		return errors.New("can not create for empty name or empty membershipType")
	}

	if !isAvailableMembershipType(request.MembershipType) {
		return errors.New("can not create for not available membershipType")
	}

	return nil
}

func (app *Application) validateUpdateRequestParameters(request UpdateRequest) error {
	if request.ID == "" || request.UserName == "" || request.MembershipType == "" {
		return errors.New("can not update for empty request")
	}

	if !isAvailableMembershipType(request.MembershipType) {
		return errors.New("can not update for not available membershipType")
	}
	return nil
}

func (app *Application) validateUpdateRequestDuplicateUserName(request UpdateRequest) error {
	foundMemberships := app.repository.FindAllByName(request.UserName)
	var found = false

	for _, foundMembership := range foundMemberships {
		if foundMembership.ID != request.ID {
			found = true
			break
		}
	}

	if found {
		return errors.New("can not update for duplication name")
	}

	return nil
}
