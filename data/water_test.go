// load_test.go
package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	usgsNumber := "04250200"
	siteData := GetSiteData(usgsNumber)

	fmt.Printf("%+v\n", *siteData)
	fmt.Println()
	assert.NotNil(t, siteData)
	assert.Equal(t, usgsNumber, siteData.Number)
	assert.Equal(t, "SALMON RIVER AT PINEVILLE NY", siteData.Name)
	assert.Equal(t, 43.53119444, siteData.Latitude)
	assert.Equal(t, -76.0376944, siteData.Longitude)
	assert.Equal(t, 478, len(siteData.Times))
	assert.Equal(t, 478, len(siteData.Readings))

}
