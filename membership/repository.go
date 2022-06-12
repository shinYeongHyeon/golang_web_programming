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

func (repository *Repository) Find(name string) (Membership, error) {
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

func (repository *Repository) issueId() string {
	return strconv.Itoa(len(repository.data) + 1)
}
