
# Golang Server Init
Template to initiate the api server using golang and mysql.

## Features:
    1. Easy to Clone and use.
    2. Used SOLID Pattern to develop SAAS
    3. Used Factory Pattern to initiate Object.
    4. Singleton Pattern to initiate database object to reduce complexity
    5. Quick solution to add any new api service.
    6. Docker enabled
    7. Dependency Injection mechenism
    8. No additional software Required
    9. Easily initiate to run multiple services using single docker-compose file
    10. Latest tools and software
    11. Integrating Independent Database for Development

## Required Software for Development
    1. Visual Studio Code

## How to add New Service
    1. Create new service folder inside services with required service name
    2. Add you service logic inside this
        1. handler.go : contains handler function
        2. init.go : contain function to initiate service
        3. request.go : contain request model and method
        4. response.go : contain response model and method
        5. service.go : contain business logic
    3. Initiate it inside "services/init.go"

## Required Software To Run
    1. Windows : Docker Desktop "https://docs.docker.com/desktop/install/windows-install/"
    2. Linux : Docker
        sudo apt update
        sudo apt install docker-ce
        sudo systemctl start docker
        sudo systemctl enable docker
        docker --version

## Run Application
    Run "docker-compose up -d"


