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
2) You can change any record by just clicking on columns of table.
3) You can add new records by clicking on + symbol.
4) You are allow to click only if excepted condition are matches.
5) You can also modify the newly added record(s).
6) Type of roll number is set to be an Integer, so for every insertion you have to enter roll number of type integer.
7) You can save record instantly into Sql.js first and then localStorage.
8) State of data is maintained using Sql.js and localStorage.
9) It fetch records only if it is not stored in localStorage.
10) Currently localStorage's data is going to remove after 60seconds, we can update it to whatever we want.
11) You can send data to server by clicking on upload button which saves your updated data permanently.
12) Errors are handled on client side.
```


## Upcoming Features

```
1) Error handling for updating roll number column.
```

### Please Note
This is a built code. I have invited you to source code on gitlab, please accept invitation. Thank you. 


## Versioning
Version 1.1
