### Fiber + GORM

Simple implementation REST API for Fiber with GORM

## Explaining test
- Use benchmarck tool in route /dogs
    - `autocannon http://localhost:3000/dogs`
- Populate database with 10 records
```curl
curl --location --request GET 'localhost:3000/dogs' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ralph",
    "breed":"Pastor",
    "age": 12,
    "isGoodBoy":true
}'
```

#### Benchmark
![bench](https://github.com/LeandroRezendeCoutinho/fiber-gorm/blob/main/img/fiber_gorm_bench.png)

#### Memory and CPU
![bench](https://github.com/LeandroRezendeCoutinho/fiber-gorm/blob/main/img/fiber_mem.png),

References:
- https://gofiber.io/
- https://gorm.io/
