createdb:
	migrate -path ./migrations/ -database "postgresql://my_ecommerce_shop:password@localhost:5432/my_ecommerce_shop?sslmode=disable" -verbose up
	migrate -path ./migrations/ -database "postgresql://my_ecommerce_shop:password@localhost:5432/my_ecommerce_shop_test?sslmode=disable" -verbose up



