package cli

import (
	"github.com/filipedtristao/hexagonal-architecture/application"
	"fmt"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""
	
	switch action {
		case "create":
			product, err := service.Create(productName, productPrice)
			
			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with name %s and price %f and %s status was created", 
				product.GetId(), 
				product.GetName(), 
				product.GetPrice(),
				product.GetStatus(),
			)
		case "enable":
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}

			product, err = service.Enable(product)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with name %s and price %f was enabled",
				product.GetId(),
				product.GetName(),
				product.GetPrice(),
			)
		case "disable":
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}

			product, err = service.Disable(product)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with name %s and price %f was disabled",
				product.GetId(),
				product.GetName(),
				product.GetPrice(),
			)
		default:
			product, err := service.Get(productId)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
				product.GetId(),
				product.GetName(),
				product.GetPrice(),
				product.GetStatus(),
			)
	}

	return result, nil
}