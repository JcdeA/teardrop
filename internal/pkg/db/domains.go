package db

import (
	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/domain"
)

func GetDomain(host string) (*ent.Domain, error) {
	domain, err := DBClient.Domain.Query().Where(domain.DomainEQ(host)).First(Ctx)
	if err != nil {
		return domain, err
	}

	return domain, nil

}

func DomainExists(host string) (*ent.Domain, error) {
	domain, err := DBClient.Domain.Query().Where(domain.DomainEQ(host)).First(Ctx)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:

		}
	}

	return domain, nil

}
