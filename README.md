# Test Walmart Search

<img src="https://github.com/alemelomeza/test-walmart-search/workflows/Lint%20Go%20Code/badge.svg?branch=main"/> <img src="https://github.com/alemelomeza/test-walmart-search/workflows/Test/badge.svg?branch=main"/> <img src="https://github.com/alemelomeza/test-walmart-search/workflows/Build%20and%20Publish%20Docker%20Image/badge.svg?branch=main"/> <img src="https://github.com/alemelomeza/test-walmart-search/workflows/Deploy/badge.svg?branch=main"/>

Test Walmart Search is a Web application that contains a search engine and results section to list the products found from a database in MongoDB. In the event that the search is a **pálindrome**, the products must be returned with a **50%** discount already applied to the price. Developed with the Golang language, including automated tests and packaged in a docker image.

## Technologies

<img src="https://img.shields.io/badge/html5%20-%23E34F26.svg?&style=for-the-badge&logo=html5&logoColor=white"/> <img src="https://img.shields.io/badge/css3%20-%231572B6.svg?&style=for-the-badge&logo=css3&logoColor=white"/> <img src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"/> <img src ="https://img.shields.io/badge/MongoDB-%234ea94b.svg?&style=for-the-badge&logo=mongodb&logoColor=white"/> <img src="https://img.shields.io/badge/github%20-%23121011.svg?&style=for-the-badge&logo=github&logoColor=white"/> <img src="https://img.shields.io/badge/github%20actions%20-%232671E5.svg?&style=for-the-badge&logo=github%20actions&logoColor=white"/> <img src="https://img.shields.io/badge/heroku%20-%23430098.svg?&style=for-the-badge&logo=heroku&logoColor=white"/> <img src="https://img.shields.io/badge/docker%20-%230db7ed.svg?&style=for-the-badge&logo=docker&logoColor=white"/>

## Use

Download source code:

```
git clone https://github.com/alemelomeza/test-walmart-search.git
```

Download products database:

```
git clone https://github.com/walmartdigital/products-db.git
```

Run products database:

```
cd products-db \
    && make database-up \
    && cd ..
# Stop database
# make database-down
```

Build image:

```
cd test-walmart-search \
    && docker build -t test-walmart-search .
```

Or download from [docker hub](https://hub.docker.com/repository/docker/alxmlo/test-walmart-search):

```
docker pull alxmlo/test-walmart-search:latest
```

Configure Dotenv file:

```
cp .env.example .env
```

Edit file `.env`:

```
DBURL=
DBNAME=
USERNAME=
PASSWORD=
```

Run container:

```
docker run \
-e PORT=80
-v $PWD/.env:/app/.env \
-p 8080:80 \
-d
test-walmart-search
```

Visit [http://localhost:8080](http://localhost:8080).

## Tests

```
docker run \
-it
-e ENVIRONMENT=test
--entrypoint go test -v -race ./...
text-walmart-search
```