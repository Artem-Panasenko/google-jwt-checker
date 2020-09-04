package googlejwtchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expiredTestToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImJjNDk1MzBlMWZmOTA4M2RkNWVlYWEwNmJlMmNlNDM3ZjQ5YzkwNWUiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL25vdGlmaWNhdGlvbi1hcGktZGV2LmxpbmtzeW5lcmd5LmNvbS9ub3RpZmljYXRpb25zIiwiYXpwIjoiMTE1Mjg2NDA0NDMwODIyMTc4NDYxIiwiZW1haWwiOiJwdWJzdWItbm90aWZpY2F0aW9uLWFwaUBhcnRlbS1wdWJzdWItcHJvamVjdC5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJleHAiOjE1OTkyMTM3MTUsImlhdCI6MTU5OTIxMDExNSwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tIiwic3ViIjoiMTE1Mjg2NDA0NDMwODIyMTc4NDYxIn0.aYVTQflafyZbaurcRE4rNl98e8lBrtXsSB5P1gNdIBTT71_A_2uRxpTyAUXWwWqICqBF-KK20G-CttGx3haeRcNNnxor423rzApPVULp4HMUv-8aH-7q0RaC3f8CfgoNW2-BFgm5AHSqFaNkmWnuxcQqA1_Yje0LIA40yybZ7zJ3TcNuXSJe4k5by-jdXXX5F_IuF65uT8A1Owh_KKdTZXrQ3aY4_X3MdLLlTiUI1VFkVuKl8vg4SYU36QC5qNbvSsUsYVk98718xp7bweDvVS4uaglk9ETdHFyCEn_6gKO04oIwfAimMMp9Y3IHS2-K-HZJtk9FpqsbLgVSGDvyiw"

	wrongSigToken = expiredTestToken + "A"
)

func TestVerifier(t *testing.T) {
	v := Verifier{}
	token, err := v.Verify(wrongSigToken)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.Equal(t, "ParseWithClaims Error: crypto/rsa: verification error", err.Error())

	token, err = v.Verify(expiredTestToken)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "ParseWithClaims Error: token is expired by")
}
