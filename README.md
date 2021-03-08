
#### Start Database
```
    docker-compose up -d database
```

#### Wait for database started and start service
```
    docker-compose up balance
```

####Api
Default users id = 1, id = 2.
1. Get user balance - ```curl --location --request GET 'localhost:8001/users/1/balance'```
2. Update user balance - ```curl --location --request POST 'localhost:8001/users/balance' \
                         --header 'Content-Type: application/json' \
                         --data-raw '{"user_id": "1", "currency": "EUR", "amount": 1000, "time_placed" : "24-JAN-20 10:27:44", "type": "deposit"}
                         '```