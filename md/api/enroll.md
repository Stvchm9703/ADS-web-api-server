# Enroll

base url: `<host>:<port>/api/v1/enroll`

## Object property table

| variables name | data types (json)                | data types (bson) | editable? |
| -------------- | -------------------------------- | ----------------- | --------- |
| `_id`          | `string`                         | `ObjectId`        | `no`      |
| `student_id`   | `string`                         | `string`          | `yes`     |
| `year`         | `int`                            | `int`             | `yes`     |
| `course_id`    | `string`                         | `string`          | `yes`     |
| `enroll_date`  | `string` (time foramt : RFC3339) | `date`            | `yes`     |
| `created_at`   | `string` (time foramt : RFC3339) | `date`            | `no`      |
| `updated_at`   | `string` (time foramt : RFC3339) | `date`            | `no`      |

---

# API Usage

## List of Enroll (GetList)

URL : \
`<host>:<port>/api/v1/enroll/list?<query>` ,or\
 `<host>:<port>/api/v1/enroll/l?<query>`

Request Methods : `GET`

Request query :
| query params  | description            | data types                         |
| ------------- | ---------------------- | ---------------------------------- |
| `pn`          | `page_num`             | `int`                              |
| `pl`          | `page_limit`           | `int`                              |
| `student_id`  | student id             | `string`                           |
| `year`        | student enrolling year | `int`                              |
| `course_id`   | course id              | `string`                           |
| `enroll_date` | date enrolled          | `string`, see also : request query |
| `created_at`  | date created at        | `string`, see also : request query |
| `updated_at`  | date last updated at   | `string`, see also : request query |

for example : `localhost:8080/api/v1/enroll/l?pn=1&pl=50&year={"eq":2017}&student_id={"eq":"54420560"}`

or: 

`localhost:8080/api/v1/enroll/l?pn=1&pl=50&`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $or:{
            "year" :  {
                $or : {
                    $eq : 2017
                }
            },
            "student_id" :{
                $or :{
                    $eq : "54420560"
                }
            },
        }
    }
}
```

output : 
``` js
{
    "page_limit": 50,
    "page_num": 1,
    "status": 0,
    "data": [
        {

        }
    ]
}
```

--- 

## Get Courses info (GetOne)


URL : \
`<host>:<port>/api/v1/enroll/get/<id>` ,or\
 `<host>:<port>/api/v1/enroll/g/<id>`

Request Methods : `GET`

Request param :
| query params | description               | data types | required? |
| ------------ | ------------------------- | ---------- | --------- |
| `_id`        | `object id of department` | `string`   | yes       |


for example : `localhost:8080/api/v1/enroll/g/54ns123idvb`


server will get the query as : 
``` js
{ 
    $find : {
        "_id" : "54ns123idvb"
    }
}
```

output : 
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "student_id" : "54420560",
        "year" : 2019,
        "course_id" : "scope_12904",
        "enroll_date" :"2019-08-01T19:47:02.411+08:00",
        "created_at": "2019-08-14T19:47:02.411+08:00",
        "updated_at": "2019-08-14T19:47:02.411+08:00"
    }
}
```


--- 

## Create Enroll record (Create)


URL : \
`<host>:<port>/api/v1/enroll/create` ,or\
 `<host>:<port>/api/v1/enroll/c`

Request Methods : `POST`

Request body :
| query params  | description            | data types                         | required? |
| ------------- | ---------------------- | ---------------------------------- | --------- |
| `student_id`  | student id             | `string`                           | `yes`     |
| `year`        | student enrolling year | `int`                              | `yes`     |
| `course_id`   | course id              | `string`                           | `yes`     |
| `enroll_date` | date enrolled          | `string`, see also : request query | `yes`     |



for example : `localhost:8080/api/v1/enroll/c`

request json body:
``` json
{
   "student_id" : "54420560",
    "year" : 2019,
    "course_id" : "scope_12904",
    "enroll_date" :"2019-08-01T19:47:02.411+08:00"
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "student_id" : "54420560",
        "year" : 2019,
        "course_id" : "scope_12904",
        "enroll_date" :"2019-08-01T19:47:02.411+08:00",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:47:02.411331+08:00"
    }
}
```

--- 

## Update Enroll record (Update)

update the object data, by required fields

URL : \
`<host>:<port>/api/v1/enroll/update` ,or\
 `<host>:<port>/api/v1/enroll/u`

Request Methods : `POST`

Request body :
| query params  | description            | data types                         | required?                      |
| ------------- | ---------------------- | ---------------------------------- | ------------------------------ |
| `_id`         | object id of course    | `string`                           | yes                            |
| `student_id`  | student id             | `string`                           | option,based on user necessary |
| `year`        | student enrolling year | `int`                              | option,based on user necessary |
| `course_id`   | course id              | `string`                           | option,based on user necessary |
| `enroll_date` | date enrolled          | `string`, see also : request query | option,based on user necessary |



for example : `localhost:8080/api/v1/enroll/u`

and the request of update the year of enroll record only 

request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a",
    "year": 2018
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "student_id" : "54420560",
        "year" : 2018,
        "course_id" : "scope_12904",
        "enroll_date" :"2019-08-01T19:47:02.411+08:00",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:58:21.27+08:00"
    }
}
```

--- 

## Delete Enroll  (Delete)

update the object data, by required fields

URL : \
`<host>:<port>/api/v1/enroll/delete` ,or\
 `<host>:<port>/api/v1/enroll/d`

Request Methods : `POST`

Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of course | `string`   | yes       |

for example : `localhost:8080/api/v1/enroll/d`
 
request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a"
}
```

server return output (success) : 
``` json
{
    "status": 0,
    "data": true
}
```
