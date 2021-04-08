package n26aas

import (
	"github.com/google/uuid"
	"github.com/nhatthm/n26api"
	"github.com/nhatthm/n26api/pkg/transaction"
)

// Service is a N26 client wrapper.
type Service struct {
	*n26api.Client
}

// TransactionsFinder provides N26 client as a transaction.Finder.
func (s *Service) TransactionsFinder() transaction.Finder {
	return s
}

// New initiates a new N26 client as a service.
func New(deviceID uuid.UUID, options ...n26api.Option) *Service {
	options = append(options, nil)
	copy(options[1:], options)
	options[0] = n26api.WithDeviceID(deviceID)

	return &Service{
		Client: n26api.NewClient(options...),
	}
}
