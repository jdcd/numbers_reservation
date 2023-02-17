# Number_reservation

This project is a test project for show some programing concepts

This project was build keeping in mind the most important
things about [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


### Run

remember set your env variables to connect to repository,
migration/mongo-init-js contains init script for your repository

    DB_COLLECTION="collection name"
    DB_NAME="db name";
    MONGO_URL="mongodb://<yourUser>:<yourPassword>@<yourHost>:"yourPort""

    go run cmd/numbers_reservation.go

docker-compose

    Run project
        
        docker-compose build
        docker-compose up



# Technologies

- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)
- [gingonic](https://github.com/gin-gonic/gin)
- [testify](https://github.com/stretchr/testify)
- [mongoDB](https://www.mongodb.com/)

