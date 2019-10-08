## Save Port Services
The project contains 2 services, the "ClientAPI" service
implements a REST interface:

*  **POST**: accepts a file containing information about seaports in JSON format, processes it and transfers ready-made data to the PortDomainService service, which in turn receives data and writes it to the database.

### brief technical description

The project is built on the principles of microservice architecture and uses a docker. Each microservice as well as the database are stored in separate containers.
The services are written in the golang programming language and use the official docker container golang. Data transfer between services occurs using GRPC. The database is selected by MongoDB and is also located in the docker container.

### Start

1. Install [Docker](https://www.docker.com/) if it is still not installed.

2. Run `$ docker-compose up --build` 

### Endpoints
* #### GET
    to get all ports `GET: http://localhost:9090` (set to limit 100, todo make pagination)
* #### GET by Id
    `GET: http://localhost:9090/{id}` where `{id}` is port identification. Geting singl port
* #### POST
    Sending json file to save data: `POST: http://localhost:9090/` Add json file to form-data
* #### DELETE
    `DELETE: http://localhost:9090/{id}` where `{id}` is port identification. Removing port from db

