# Race Manager
A league scoring and points manager for NR2003

### Getting Started
Install sqlite3, then run the following:
```BASH
git clone --recurse-submodules https://github.com/kylequamme/race-manager.git  
cd race-manager  
mkdir sqlite  
sqlite3 sqlite/racemanager.db  
go run .
```
To populate the db with tables and example data, uncomment lines 23-47 of api.go and restart.