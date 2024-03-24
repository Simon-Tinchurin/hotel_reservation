# Hotel reservation backend

## Project outline
- users -> book room from an hotel
- admins -> to check reservations/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver
Documentation
https://www.mongodb.com/docs/drivers/go/current/

Install mongodb client
go get go.mongodb.org/mongo-driver/mongo

### gofiber
Documentation
https://gofiber.io/

Installing gofiber
go get github.com/gofiber/fiber/v2

## Docker
## Installing mongodb as a Docker container
docker run --name mongodb -d mongo:latest -p 27017:27017

1. run 'go mod init hotel-reservation' in the terminal
2. run 'go get github.com/gofiber/fiber/v2' in the terminal
3. write a Makefile
3. run 'make build' in the terminal
// @ in the beginning of a command hides output
4. Install mongodb as a Docker container - run 'docker run --name mongodb -d --restart=unless-stopped -p 27017:27017 mongo:latest mongod --bind_ip_all' in the terminal
5. Install mongodb client - run 'go get go.mongodb.org/mongo-driver/mongo' in the terminal
