package membership

import (
	"errors"
	"strconv"
)

type Repository struct {
	data map[string]Membership
}

var inMemoryRepository *Repository

func NewRepository(data map[string]Membership) *Repository {
	inMemoryRepository = &Repository{data: data}

	return inMemoryRepository
}

// Find : find membership by id
func (repository *Repository) Find(id string) Membership {
	var foundMembership Membership

	for _, membership := range repository.data {
		if membership.ID == id {
			foundMembership = membership
			break
		}
	}

	return foundMembership
}

// FindAllByName : find all membership by name
func (repository *Repository) FindAllByName(name string) []Membership {
	var foundMemberships []Membership

	for _, membership := range repository.data {
		if membership.UserName == name {
			foundMemberships = append(foundMemberships, membership)
		}
	}

	return foundMemberships
}

// FindByName : find membership by name
func (repository *Repository) FindByName(name string) (Membership, error) {
	var foundMembership Membership
	found := false

	for _, membership := range repository.data {
		if membership.UserName == name {
			found = true
			foundMembership = membership
		}
	}

	if !found {
		return foundMembership, errors.New("Can not find " + name + " user")
	}

	return foundMembership, nil
}

func (repository *Repository) Create(name string, membership string) Membership {
	id := repository.issueId()
	createdMembership := Membership{
		ID:             id,
		UserName:       name,
		MembershipType: membership,
	}

	repository.data[id] = createdMembership

	return createdMembership
}

func (repository *Repository) Update(id string, name string, membershipType string) Membership {
	var toUpdateMembership Membership

	for _, membership := range repository.data {
		if membership.ID == id {
			toUpdateMembership = Membership{
				ID:             id,
				UserName:       name,
				MembershipType: membershipType,
			}
			repository.data[id] = toUpdateMembership
		}
	}

	return toUpdateMembership
}

func (repository *Repository) issueId() string {
	return strconv.Itoa(len(repository.data) + 1)
}
