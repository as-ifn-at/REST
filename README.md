# Glofox


### To run the application
> source configuration/env.sh && go run cmd/glofox/main.go 


curl --location 'localhost:8080/classes/v1/create' \
--header 'Content-Type: application/json' \
--data '{
    "class_name":"Pilates",
    "start_date":"01-12-2024",
    "end_date":"20-12-2024",
    "capacity": 20
}'

curl --location 'localhost:8080/classes/v1/Pilates'

