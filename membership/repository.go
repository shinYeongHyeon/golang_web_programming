package membership

import (
	"strconv"
)

// Repository : struct for membership data
type Repository struct {
	data map[string]Membership
}

var inMemoryRepository *Repository

// NewRepository : create repository
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
func (repository *Repository) FindByName(name string) Membership {
	var foundMembership Membership

	for _, membership := range repository.data {
		if membership.UserName == name {
			foundMembership = membership
			break
		}
	}

	return foundMembership
}

// Create : save for membership by name, membershipType
func (repository *Repository) Create(name string, membershipType string) Membership {
	id := repository.issueId()
	createdMembership := Membership{
		ID:             id,
		UserName:       name,
		MembershipType: membershipType,
	}

	repository.data[id] = createdMembership

	return createdMembership
}

// Update : update for membership by id, name, membershipType
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

// Delete : delete membership by id
func (repository *Repository) Delete(id string) {
	delete(repository.data, id)
}

func (repository *Repository) issueId() string {
	return strconv.Itoa(len(repository.data) + 1)
}
