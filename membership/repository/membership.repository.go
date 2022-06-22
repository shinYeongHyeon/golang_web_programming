package membershipRepository

import (
	membershipPresentation "github.com/shinYeongHyeon/golang_web_programming/membership/presentation"
	"sort"
	"strconv"
)

type Repository struct {
	data map[string]membershipPresentation.Membership
}

var inMemoryRepository *Repository

func NewRepository(data map[string]membershipPresentation.Membership) *Repository {
	inMemoryRepository = &Repository{data: data}

	return inMemoryRepository
}

func (repository *Repository) FindAll() []membershipPresentation.Membership {
	var foundMemberships []membershipPresentation.Membership

	memberships := getOrderedSlices(repository)

	for _, membership := range memberships {
		foundMemberships = append(foundMemberships, membership)
	}

	return foundMemberships
}

func getOrderedSlices(repository *Repository) []membershipPresentation.Membership {
	var slices []membershipPresentation.Membership

	sortKeys := make([]string, 0, len(repository.data))
	for k := range repository.data {
		sortKeys = append(sortKeys, k)
	}

	sort.Strings(sortKeys)

	for _, k := range sortKeys {
		slices = append(slices, repository.data[k])
	}

	return slices
}

func (repository *Repository) Find(id string) membershipPresentation.Membership {
	var foundMembership membershipPresentation.Membership

	for _, membership := range repository.data {
		if membership.ID == id {
			foundMembership = membership
			break
		}
	}

	return foundMembership
}

func (repository *Repository) FindAllByName(name string) []membershipPresentation.Membership {
	var foundMemberships []membershipPresentation.Membership

	for _, membership := range repository.data {
		if membership.UserName == name {
			foundMemberships = append(foundMemberships, membership)
		}
	}

	return foundMemberships
}

func (repository *Repository) FindByName(name string) membershipPresentation.Membership {
	var foundMembership membershipPresentation.Membership

	for _, membership := range repository.data {
		if membership.UserName == name {
			foundMembership = membership
			break
		}
	}

	return foundMembership
}

func (repository *Repository) Create(name string, membershipType string) membershipPresentation.Membership {
	id := repository.issueId()
	createdMembership := membershipPresentation.Membership{
		ID:             id,
		UserName:       name,
		MembershipType: membershipType,
	}

	repository.data[id] = createdMembership

	return createdMembership
}

func (repository *Repository) Update(id string, name string, membershipType string) membershipPresentation.Membership {
	var toUpdateMembership membershipPresentation.Membership

	for _, membership := range repository.data {
		if membership.ID == id {
			toUpdateMembership = membershipPresentation.Membership{
				ID:             id,
				UserName:       name,
				MembershipType: membershipType,
			}
			repository.data[id] = toUpdateMembership
		}
	}

	return toUpdateMembership
}

func (repository *Repository) Delete(id string) {
	delete(repository.data, id)
}

func (repository *Repository) issueId() string {
	return strconv.Itoa(len(repository.data) + 1)
}
