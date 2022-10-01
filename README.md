An example project using hexagonal/ports and adapters architecture

generating mocks
```
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

running the tests
```
go test ./...
```

creating a product with the CLI
```
go run main.go cli -a=create -n="Product" -p=99.99
```

searching a product with the CLI
```
go run main.go cli -i={producId}
```

starting the http server
```
go run main.go http
```

