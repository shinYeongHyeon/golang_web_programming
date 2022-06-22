package membershipServices

import (
	membershipPresentation "github.com/shinYeongHyeon/golang_web_programming/membership/presentation"
	membershipRepository "github.com/shinYeongHyeon/golang_web_programming/membership/repository"
	membershipServicesDto "github.com/shinYeongHyeon/golang_web_programming/membership/services/dto"
	"strconv"
)

var SuccessCreate = "OK"
var DuplicateUserName = "USER_NAME_DUPLICATED"
var NotAvailableMembershipType = "NOT_AVAILABLE_MEMBERSHIP_TYPE"
var SuccessFound = "OK"
var NotFoundMembership = "NOT_FOUND"
var SuccessUpdate = "OK"
var SuccessDelete = "OK"

type Service struct {
	repository membershipRepository.Repository
}

func NewService(repository membershipRepository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(req membershipServicesDto.CreateRequest) membershipServicesDto.CreateResponse {
	if isAvailable := isAvailableMembershipType(req.MembershipType); !isAvailable {
		return membershipServicesDto.CreateResponse{Code: NotAvailableMembershipType}
	}

	if foundMembership := s.repository.FindByName(req.UserName); foundMembership.ID != "" {
		return membershipServicesDto.CreateResponse{Code: DuplicateUserName}
	}

	membership := s.repository.Create(req.UserName, req.MembershipType)

	return membershipServicesDto.CreateResponse{Code: SuccessCreate, MembershipUser: generateMembershipUserDto(membership)}
}

func (s *Service) FindAll(req membershipServicesDto.FindAllRequest) membershipServicesDto.FindAllResponse {
	memberships := getMemberships(s)
	if req.Offset == "" || req.Limit == "" {
		return membershipServicesDto.FindAllResponse{MembershipUsers: memberships}
	}

	offset, _ := strconv.Atoi(req.Offset)
	limit, _ := strconv.Atoi(req.Limit)
	memberships = getFilteredMemberships(offset, limit, memberships)

	return membershipServicesDto.FindAllResponse{MembershipUsers: memberships}
}

func getFilteredMemberships(offset int, limit int, memberships []*membershipServicesDto.MembershipUser) []*membershipServicesDto.MembershipUser {
	current := 0
	var filteredMemberships []*membershipServicesDto.MembershipUser

	for index, membership := range memberships {
		if current >= limit {
			break
		}

		if index >= offset {
			filteredMemberships = append(filteredMemberships, membership)
			current = current + 1
		}
	}

	return filteredMemberships
}

func getMemberships(s *Service) []*membershipServicesDto.MembershipUser {
	var memberships []*membershipServicesDto.MembershipUser
	foundMemberships := s.repository.FindAll()
	for _, membership := range foundMemberships {
		memberships = append(memberships, generateMembershipUserDto(membership))
	}

	return memberships
}

func (s *Service) FineOne(id string) membershipServicesDto.FindOneResponse {
	membership := s.repository.Find(id)

	if membership.ID == "" {
		return membershipServicesDto.FindOneResponse{
			Code:           NotFoundMembership,
			MembershipUser: nil,
		}
	}

	return membershipServicesDto.FindOneResponse{
		Code:           SuccessFound,
		MembershipUser: generateMembershipUserDto(membership),
	}
}

func (s *Service) UpdateOne(req membershipServicesDto.UpdateRequest) membershipServicesDto.UpdateResponse {
	if foundMembership := s.repository.Find(req.ID); foundMembership.ID == "" {
		return membershipServicesDto.UpdateResponse{Code: NotFoundMembership}
	}

	if foundMembership := s.repository.FindByName(req.UserName); foundMembership.ID != "" {
		return membershipServicesDto.UpdateResponse{Code: DuplicateUserName}
	}

	if isAvailable := isAvailableMembershipType(req.MembershipType); !isAvailable {
		return membershipServicesDto.UpdateResponse{Code: NotAvailableMembershipType}
	}

	membership := s.repository.Update(req.ID, req.UserName, req.MembershipType)

	return membershipServicesDto.UpdateResponse{Code: SuccessUpdate, MembershipUser: generateMembershipUserDto(membership)}
}

func isAvailableMembershipType(membershipType string) bool {
	isAvailable := false
	for _, value := range membershipPresentation.AvailableMemberships {
		if value == membershipType {
			isAvailable = true
			break
		}
	}

	return isAvailable
}

func (s *Service) DeleteOne(id string) membershipServicesDto.DeleteResponse {
	if foundMembership := s.repository.Find(id); foundMembership.ID == "" {
		return membershipServicesDto.DeleteResponse{Code: NotFoundMembership}
	}

	s.repository.Delete(id)

	return membershipServicesDto.DeleteResponse{Code: SuccessDelete}
}

func generateMembershipUserDto(membership membershipPresentation.Membership) *membershipServicesDto.MembershipUser {
	return &membershipServicesDto.MembershipUser{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}
}
