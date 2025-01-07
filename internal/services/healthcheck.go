package services

import "github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"

type Healthcheck struct {
	HealthcheckRepository interfaces.IHealthcheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "service healthy", nil
}
