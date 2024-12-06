# The Gin for Go framework

## Contents
- Handlers
- Middleware - Logrus [https://github.com/sirupsen/logrus/](https://github.com/sirupsen/logrus/)

## Features:
- Route Grouping: [./routes/routes.go](./routes/routes.go)

- Gin Handlers: [./handlers/handlers.go](./handlers/handlers.go)
    ## Endpoints:
        - /getData
        - /postData
        ### Dealing with Query strings
        - /getQS   
        ### Dealing with query Paramaters
        - //getQueryParams/:name/:age

- MiddleWare :
    1. Authentication Middleware [./middleware/middleware.go](./middleware/middleware.go):
    ### Checks if the request made to the endpoint above /getData has Token: Auth and raises a status Unauthorised if False

    2. Logging :
    ### Used LogRus for logging [./middleware/middleware.go/LoggingMidd](./middleware/middleware.go/LoggingMidd)




