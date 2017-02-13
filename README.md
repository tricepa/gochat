# gochat
A Golang and React chat application backed by MySQL

## Setup
### Prerequisites
* [MySQL installed] (https://dev.mysql.com/)
* [Golang installed] (https://golang.org/doc/install)
* [yarn] (https://yarnpkg.com/en/docs/install) or [npm] (https://docs.npmjs.com/getting-started/installing-node) installed
* [webpack] (https://github.com/webpack/webpack) installed

### Configure database
* Log in to MySQL and create a database 'chat' by running `create database chat`
* Import the chat database structure by running `mysql -u <username> -p chat < config/chat_database.sql`, replacing `<username>` with your MySQL username
* Export the database data source name used for connecting by running `export CHAT_DB_DSN="<username>:<password>@/chat?charset=utf8&parseTime=true"`, replacing `<username>` and `<password>` with your MySQL credentials

### Install dependencies
* run `go install` 
* run `yarn install` or `npm install`
* run `webpack`

### Start server
`go run server.go`

## Features
* Persistent database stores username and messages
* Supports multiple connections and broadcasted messages through websockets 
* On login, shows new (unseen by that user) messages by default. The user may then choose to display all messages.

## Resources
* [Setting up for React] (https://scotch.io/tutorials/setup-a-react-environment-using-webpack-and-babel)
* [Gin] (https://github.com/gin-gonic/gin)
* [Melody] (https://github.com/olahol/melody/)
* [go-sql-driver] (https://github.com/go-sql-driver/mysql)
* [Go React chat example] (https://github.com/abhishekpillai/go-react-chat-app)

## Looking ahead
### Application
* Incorporate user authentication
* Allow user to select number of lines to display in chatroom
* Add mute feature
### Implementation
* Add config file for more customizable usage
* Add testing
