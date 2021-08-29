package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseContactsShouldParseDataSuccessfully(t *testing.T) {
	fakeData := [][]string{{
		"Murat", "905390001122",
	}}
	contacts := parseContacts(fakeData)

	assert.Equal(t, 1, len(contacts))
}

func TestParseOrdersShouldParseDataSuccessfully(t *testing.T) {
	fakeData := [][]string{{
		"Fake Message", "905390001122", "2021-10-10T12:12:12",
	}}
	orders := parseOrders(fakeData)

	assert.Equal(t, 1, len(orders))
}
