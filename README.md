
# userbase-api

## Description
A basic userbase API backend servic implemented in Golang with a MongoDB Atlas cloud DB. Swagger is used to generate the documentation [(snapshot)](docs/Swagger%20Docs.mhtml). 

## Setup
First, to connect to the db, obtain `mongoCert.pem` and place it in `/cert`.

There are two ways to run the server.
- Run `./main.exe` (or main_Linux/main_MacOS depending on your OS)
- With go 1.17, `go run server/main.go`

## Usage
- Open [localhost](http://localhost:10000/users) in a web browser (or Postman)
- Open [Swagger UI](http://localhost:10000/swagger/index.html#) to view docs and test functionality

## Directory Structure
```
/cert: stores the MongoDB Atlas Certificate required
/docs: stores a snapshot of Swagger documentation
/server: source code for API
    /dal: data model & driver code for 3rd party services e.g. mongoDB
    /docs: generated Swagger docs
    /handlers: handlers for common middleware & for each API e.g. users
    /service: business logic for each endpoint e.g. users
    /utils: simple utils functions, custom errors, etc.
```

---
## Task 
Build a RESTful API that can get/create/update/delete user data from a persistence database
### User Model
    -  "id"
    -  "name"
    -  "dob"
    -  "address"
    -  "description"
    -  "createdAt"
### Functionality
    The API should follow typical RESTful API design pattern.
    The data should be saved in the DB.
    Provide proper unit test.
    Provide proper API document.
### Tech stack
    Use any framework.
    Use any DB. NoSQL DB is preferred.
### Bonus
    Write clear documentation on how it's designed and how to run the code.
    Write good in-code comments.
    Write good commit messages.

### Advanced requirements
- These are used for some further challenges. You can safely skip them if you are not asked to do any, but feel free to try out.
- Provide a complete user auth (authentication/authorization/etc.) strategy, such as OAuth.
- Provide a complete logging (when/how/etc.) strategy.
- Imagine we have a new requirement right now that the user instances need to link to each other, i.e., a list of "followers/following" or "friends". Can you find out how you would design the model structure and what API you would build for querying or modifying it?
- Related to the requirement above, suppose the address of the user now includes a geographic coordinate(i.e., latitude and longitude), can you build an API that,
- given a user name
- return the nearby friends