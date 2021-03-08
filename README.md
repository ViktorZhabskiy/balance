Default users id = 1, id = 2.

#### Start Database
```
    docker-compose up database -d
```

#### Start service
```
    docker-compose up
```

####Api
Get user balance - ```curl --location --request GET 'localhost:8001/users/1/balance'```
Update user balance - ```curl --location --request POST 'localhost:8001/users/balance' \
                         --header 'Content-Type: application/json' \
                         --data-raw '{"user_id": "1", "currency": "EUR", "amount": 1000, "time_placed" : "24-JAN-20 10:27:44", "type": "withdrawal"}
                         '```