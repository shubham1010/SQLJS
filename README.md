# SQL.js and localStorage to store and manipulate data on client side 


## Getting Started

1. Clone the project
2. cd SQLJS/
3. ./run.sh, please change permission of file in case it arise any error, follow "chmod 754 run.sh" command without quote.
4. If you are using Firefox browser, it will take you to index.html page otherwise other than Firefox browser please open client/index.html file in any browser.

### Prerequisite
```
Docker =>  As I have used Golang for API, I thought that without bothering about installation and dependencies of Golang instead put that in container to make it run easily. 
```
### Installing

1) [DOCKER](https://www.docker.com/)
  For installation please refer [Here](https://docs.docker.com/get-docker/)

2) [Docker-Compose](https://docs.docker.com/compose/)
  For installation please refer [Here](https://docs.docker.com/compose/install/)

## Built With Stack

* [HTML5]()- The langauge which is basic building block of webpage is used for to structure elements for webpage.
* [CSS]()- The language is used to give design/style for the inserted elements to look pretty good and well structure.
* [Javascript]()- The language is used for accessing DOM elements, fetch API request and load to the DOM content.
* [Sql.js]()- Used to store records of students.
* [Golang]()- API calls.
* [Docker]()- To take care of dependencies.


## Current Features

```
1) You can display records from .json to browser.
2) You can change any record by just clicking into columns of table.
3) You can save record instantly into Sql.js first and then localStorage.
4) State of data is maintained using Sql.js and localStorage.
5) It fetch records only if it is not stored in localStorage.
6) Currently localStorage's data is going to remove after 60seconds, we can update it to whatever we want.
```


## Upcoming Features

```
1) Create New button, to create new record on webpage itself.
2) Upload button, to add/insert updated into server side's .json file.
3) Error handling to give respective response for user actions, currently I'm just printing it to console.log().
```

### Please Note
This is a built code, I shall give you source code once I complete all upcoming features, thank you.


## Versioning
Version 1.0
