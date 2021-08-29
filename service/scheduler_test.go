package service

import (
	mocks "awesomeProject/.mocks"
	"github.com/golang/mock/gomock"
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

func TestCreateInstantJobShouldReturnNilError(t *testing.T) {
	mockController := gomock.NewController(t)
	wpMockclient := mocks.NewMockChatClient(mockController)

	wpMockclient.EXPECT().SendMessage(gomock.Any(), gomock.Any()).Times(1)
	service := SchedulerService{Client: wpMockclient}
	contactPath := "./test-data/contact.csv"
	err := service.CreateInstantJob(contactPath)
	if err != nil {
		return
	}

	assert.Nil(t, err)
}

func TestReadCsvFromFileShouldReturnErrorWithWrongPath(t *testing.T) {
	_, err := readCsvFromFile("./qwe.csv")
	assert.Error(t, err)
}

func TestReadCsvFromFileShouldReturnSuccessfullyData(t *testing.T) {
	expectedData := [][]string{
		{"Murat", "0539"},
	}
	file, err := readCsvFromFile("./test-data/contact.csv")

	assert.Equal(t, expectedData, file)
	assert.Nil(t, err)
}

func TestCreateScheduleJobShouldReturnNilError(t *testing.T) {
	mockController := gomock.NewController(t)
	wpMockclient := mocks.NewMockChatClient(mockController)

	wpMockclient.EXPECT().SendMessage(gomock.Any(), gomock.Any()).Times(1)
	service := SchedulerService{Client: wpMockclient}
	ordersPath := "./test-data/orders.csv"
	err := service.CreateScheduleJob(ordersPath)
	if err != nil {
		return
	}

	assert.Nil(t, err)
}
