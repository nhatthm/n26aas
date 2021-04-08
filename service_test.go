package n26aas_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nhatthm/n26api"
	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/n26aas"
)

func TestNew(t *testing.T) {
	t.Parallel()

	deviceID := uuid.New()

	n26aas.New(deviceID,
		n26api.WithUsername("foo"),
		n26api.WithPassword("bar"),
		func(c *n26api.Client) {
			assert.Equal(t, deviceID, c.DeviceID())
		},
	)
}

func TestService_TransactionsFinder(t *testing.T) {
	t.Parallel()

	s := n26aas.New(uuid.New())

	assert.Equal(t, s, s.TransactionsFinder())
}
