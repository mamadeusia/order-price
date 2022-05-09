# ORDER - PRICE SERVICE 

## Description

In this project we get two request from user . the request handlers are in `server/http/` folder and the main feature are ability to create user and create and update order . 

for database we use redis. for managing inputs to the redis i prefer to create `Redisable` interface in `data` folder ,for persist data in redis i user setH modes for better performance . 

before persisting data in redis we have to get the price from price service , so i used grpc for fetching data from price service and for notifing price service about new order for calculating new price i have to choice between rabbitMQ as a message broker or kafka for loging system , because i have no knowledge about kafka and i'm reading about that i prefer to use rabbit but i think it is better to implement the connection between two services with kafka . 

for config you can set configs in `config.yaml` . the docker-compose is currently in writing . 


## How to improve? 

As i mentioned before , i think it's better to use kafka and distributed log functionality for communication layer between services . 

Folder structure of two services are interwind , for core layer i think i need to refactor , for data layer before persist data in redis in `createOrder` function in data layer i call the `GetCurrentPrice` from server package that get data from grpc , i think i need to call `GetCurrentPrice` before calling `createOrder` to write code more decoupled and have better orthogonality between packages .  
