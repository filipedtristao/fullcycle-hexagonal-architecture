package cli_test

import (
	mock_application "github.com/filipedtristao/hexagonal-architecture/application/mocks"
	"github.com/filipedtristao/hexagonal-architecture/adapters/cli"
	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
	"testing"
	"fmt"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "1"
	productName := "Product 1"
	productPrice := 10.0
	productStatus := "disabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf(
		"Product ID %s with name %s and price %f and %s status was created", 
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID %s with name %s and price %f was enabled",
		productId,
		productName,
		productPrice,
	)

	result, err = cli.Run(productServiceMock, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID %s with name %s and price %f was disabled",
		productId,
		productName,
		productPrice,
	)

	result, err = cli.Run(productServiceMock, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(productServiceMock, "", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}