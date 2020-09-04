package googlejwtchecker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCerts(t *testing.T) {
	certs, err := getCerts()
	assert.Nil(t, err, "Unexpected error")

	cacheAge := certs.Expiry.Sub(time.Now()).Seconds()
	assert.GreaterOrEqual(t, cacheAge, float64(7200), "max-age not found")

	key := certs.Keys["bc49530e1ff9083dd5eeaa06be2ce437f49c905e"]
	assert.NotNil(t, key, "bc49530e1ff9083dd5eeaa06be2ce437f49c905e should exists")
}

func TestGetCertsCache(t *testing.T) {
	certs = &Certs{
		Expiry: time.Now(),
	}
	certs, err := getCerts() // trigger update
	assert.Nil(t, err)

	key := certs.Keys["bc49530e1ff9083dd5eeaa06be2ce437f49c905e"]
	assert.NotNil(t, key, "bc49530e1ff9083dd5eeaa06be2ce437f49c905e should exists")
}
