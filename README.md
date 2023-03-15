![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/aaqibb13/guinevere) ![Lines of code](https://img.shields.io/tokei/lines/github/aaqibb13/guinevere) ![GitHub last commit](https://img.shields.io/github/last-commit/aaqibb13/guinevere)
# AranGO-base-setup
This respository contains database configuration for ArangoDB database for your Go project along with base-crud, creation of Collections, Graphs (coming soon) Custom Analyzers, Views, Query Executor and DB Transaction

## Direct Dependencies (Minimalistic)
1. **go-driver**: The go-driver is leveraged to provide the functionality a user might need.
2. **viper**: is used to read config from environment variables. 
3. **logrus**: is used for logging so that you can further make use of extended functionality of the logrus package in case you might want to set the formatting levels, make your logs more verbose. 
  
    **Note:** logrus can be omitted if you feel the need to.

   
## Included functionalities
The default incorporated functionalities that comes along are:
- [x] Creation of Databases and corresponding users
- [x] Creation of **Collections**,**Custom Analyzer** and **Views** at the startup (preferably)
- [x] Query Executor for Execution of raw AQL queries
- [x] QueryExecutor with cursor full count while executing raw AQL queries
- [x] Initializing a DB **Transaction**

# TODO
- [ ] Add tests for every function (Unit tests)
- [ ] Add creation of Graphs 
- [ ] Add different analyzer types (like norm, ngram, pipeline etc)
- [ ] Include Edit View functionality (ArangoDB does not support editing the views by default)
- [ ] Improve documentation and code structure
- [ ] Add scripts for backing up data

## How to use:
- Clone the repository at your preferred location:

          git clone https://github.com/aaqibb13/guinevere.git
      
- Run:

          go mod tidy
- You're ready to use the base setup.

## Running arangodb locally
- Make sure you have docker installed on your system (based on which machine you're on: `Linux`, `Mac` or `Windows`)
- cd into `guinevere/deployments` and run:
  
        docker-compose -f arango-docker-compose.yml up --build -d db
- You should be able to view the ArangoDB server running on `127.0.0.1:8529` which you can access using the credentials specified in the `arango-docker-compose.yml` file.
