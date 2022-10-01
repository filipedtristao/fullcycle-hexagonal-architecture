package application

type ProductService struct {
	ProductPersistence ProductPersistenceInterface
}

func NewProductService(productPersistence ProductPersistenceInterface) *ProductService {
	return &ProductService{productPersistence}
}

func (p *ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.ProductPersistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()

	if err != nil {
		return nil, err
	}

	result, err := p.ProductPersistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return nil, err
	}

	result, err := p.ProductPersistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return nil, err
	}

	result, err := p.ProductPersistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}