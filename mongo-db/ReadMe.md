**Supported APIs**
- Get all users: http://localhost:8080/users
- Get one user: http://localhost:8080/users/{user_id}
- Delete user: http://localhost:8080/users/delete/{user_id}
- Insert new user -> POST: http://localhost:8080/createUser

**Mongo DB Installation**
- URL for Mongo DB installation: https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/

- Run below commands to install the MongoDB
> `brew tap mongodb/brew`<br/>
> `brew update`<br/>
> `brew install mongodb-community@7.0`<br/>

- To Start the Service run below command
> `brew services start mongodb-community@7.0`
- To Stop the Service run below command
> `brew services stop mongodb-community@7.0`

**Mongo DB Compass (GUI) Installation**
> URL to download and install -> https://www.mongodb.com/try/download/compass <br/>
> URL for intel MAC -> https://downloads.mongodb.com/compass/mongodb-compass-1.42.0-darwin-x64.dmg <br/>
> URL for Apple Silicon MAC -> https://downloads.mongodb.com/compass/mongodb-compass-1.42.0-darwin-arm64.dmg <br/>

- It by defaults connects to:
> mongodb://localhost:27017

**Important Links**
> mongodb-golang driver https://github.com/mongodb/mongo-go-driver

**External Packages**
> go get -u github.com/gorilla/mux  
> go get go.mongodb.org/mongo-driver/mongo
