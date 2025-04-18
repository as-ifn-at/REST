# Glofox


### To run the application
> source configuration/env.sh && go run cmd/glofox/main.go 


### create classes

curl --location 'localhost:8080/classes/v1/create' \
--header 'Content-Type: application/json' \
--data '{
    "class_name":"Pilates",
    "start_date":"01-12-2024",
    "end_date":"20-12-2024",
    "capacity": 20
}'

### Query class
curl --location 'localhost:8080/classes/v1/Pilates'

### Do class booking

curl --location 'localhost:8080/bookings/v1/book' \
--header 'Content-Type: application/json' \
--data '{
    "name":"Asif",
    "date":"02-12-2025",
    "class_name":"Pilates"
}'

### Query booking

curl --location 'localhost:8080/bookings/v1/Asif'