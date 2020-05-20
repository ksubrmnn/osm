package vault

import (
	"sync"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/open-service-mesh/osm/pkg/certificate"
)

// CertManager implements certificate.Manager and contains a Hashi Vault client instance.
type CertManager struct {
	// How long will newly issued certificates be valid for
	validity time.Duration

	// The Certificate Authority root certificate to be used by this certificate manager
	ca certificate.Certificater

	// The channel announcing to the rest of the system when a certificate has changed
	announcements chan interface{}

	// Cache for all the certificates issued
	cache     *map[certificate.CommonName]certificate.Certificater
	cacheLock sync.Mutex

	// Hashicorp Vault client
	client *api.Client
}