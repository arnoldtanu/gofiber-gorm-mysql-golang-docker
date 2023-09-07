# Disbursement API
API endpoint for disbursement of user balance from mysql database.

### Getting Started

This app can be run through docker, here is how to run it:

##### 1. Clone this project
##### 2. Download docker
- Download docker desktop from https://www.docker.com/products/docker-desktop/ and install it.

##### 3. Build and run from docker
- Enter project's root folder, and run command inside terminal / cmd:

```
docker-compose up --build
```
- Docker compose will download and build necesary image (golang and mysql).

##### 4. Wait until MySQL and Golang server is ready
- wait until mysql and go are ready to use which is indicated by the appearance of a message from fiber

##### 5. Send JSON payload
- Submit JSON payload to localhost:3000/disbursement
- example:
```
{
    "id":2,
    "amount":1000,
    "passkey":"secret"
}
```

## Build With
- [Go-lang](https://go.dev/)
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/index.html)
- [Fiber](https://gofiber.io/)