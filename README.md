# SCTRACK

Created a todo list for a mini carrier demo. The app consists of a Web Service with several calls, supported by a 
Postgres database. The Postgres database has three tables to simulate a carrier and todo list.

## Render
This application is being hosted at Render, using the free tier. The following link points to the swagger page. 
If the app is not running, is because the free tier might have stopped it. Send me a message at 
alpha.aleph@gmail.com, and I'll be happy to get it started for you.  :-)

* Swagger link : https://sctrack.onrender.com/swagger/index.html#/

## Docker Hub
The application container will also be served in Docker Hub under:
```
docker pull alphaaleph/sctrack:v1.4
```

## Database
The database provides four tables for the demo:
* carrier - holds a carrier's name and telephone
* todos - holds the todos for each carrier
* journal - used to keep track of the carrier's goings
* action - list of actions associated with todos

## Connect
Current database connection:
```
DB_USER=sctrack_pg_user
DB_PASSWORD=7jAZPc5kmiqafsK37zc5XA6LFx8XE9H3
DB_NAME=sctrack_pg
DB_HOST=dpg-ck19417dorps73c4k42g-a.oregon-postgres.render.com
DB_PORT=5432
DB_DRIVER=postgres
SERVER_PORT=3030
```

## Data
There is a Python and SQL scripts under the data directory that initializes the database with dummy data. Update the 
connection_url in the Python script if you want to point to a different Postgres database. The SQL script has some 
dummy data included.
```
cd ./data
python ./init.py
```

## Building
Included with the project, there is a script to build the application from the command line. After getting a copy of 
the project, run "build.sh" from the command line.  

The script will do the following:

* Build the web service application locally
* Auto generate the swagger page
* Build a web service docker image

## Packages
Some of the packages used in the program
* [gorilla mux](https://github.com/gorilla/mux) - router and dispatcher
* [postgres](github.com/lib/pq) - a Go postgres driver
* [slog](golang.org/x/exp/slog) - Go's structured logging
* [swaggo](https://github.com/swaggo/swag) - swagger documentation

## Swagger
![plot](./assets/swagger_calls.png)
