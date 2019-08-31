# student

base url: `<host>:<port>/api/v1/<methods>/student`

## Object property table

| variables name | data types (json)                | data types (bson) | editable only? |
| -------------- | -------------------------------- | ----------------- | -------------- |
| `_id`          | `string`                         | `ObjectId`        | `no`           |
| `student_id`      | `string`                         | `string`          | `yes`          |
| `student_name`    | `string`                         | `string`          | `yes`          |
| `DOB`     | `string`(time foramt : RFC3339)         | `date`          | `yes`          |
| `created_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |
| `updated_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |

---

# API Usage

## List of student (GetList)

### URL :

`<host>:<port>/api/v1/l/student?<query>` 

Request Methods : `GET`

### Request query :
| query params | description            | data types                         |
| ------------ | ---------------------- | ---------------------------------- |
| `pn`         | `page_num`             | `int`                              |
| `pl`         | `page_limit`           | `int`                              |
|`sort` | `sort` | `json string` see also : request query|
| `qbind` | query bindings | `string` see also : request query|
| `student_id`    | student id          | `string`, see also : request query |
| `student_name`  | student name        | `string`, see also : request query |
| `dob`   | day of birth of student | `string`, see also : request query |
| `created_at` | date created at        | `string`, see also : request query |
| `updated_at` | date last updated at   | `string`, see also : request query |

### example : 
`localhost:8080/api/v1/l/student/?student_id={"eq":54420440}`

or

`localhost:8080/api/v1/l/student/?student_id=$eq:54420440`



server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "student_id": {
                $or : {
                   $eq : 54420440
                } ,
            },
        }
    }
}
```

output : 
``` json
{
    "page_limit": 500,
    "page_num": 1,
    "count": 1,
    "SortAr": null,
    "status": 0,
    "data": [
        {
            "_id": "5d5f7d2708386429880bad39",
            "student_id": "54420440",
            "student_name": "Chan Tai Man",
            "dob": "1987-06-04T00:00:00+09:00",
            "enrolled": [
                {
                    "_id": "5d5fbbce0838642cbcb221ff",
                    "year": 2019,
                    "course_id": "scope_23010",
                    "enroll_date": "2019-08-16T00:00:00+08:00",
                    "created_at": "2019-08-23T18:11:26.112+08:00",
                    "updated_at": "2019-08-23T18:11:26.112+08:00"
                }
            ],
            "created_at": "2019-08-23T13:44:07.403+08:00",
            "updated_at": "2019-08-23T16:58:52.575+08:00"
        }
    ]
}
```

--- 

## Get student info (GetOne)


### URL :

`<host>:<port>/api/v1/g/student/<id>`

Request Methods : `GET`

Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `_id`         | `object id of student` | `string`   | yes       |


### example : 
`localhost:8080/api/v1/student/g/5d5f7d2708386429880bad39`


server will get the query as : 
``` js
{ 
    $find : {
        "_id" : "5d5f7d2708386429880bad39"
    }
}
```

output : 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": {
        "_id": "5d5f7d2708386429880bad39",
        "student_id": "54420440",
        "student_name": "Chan Tai Man",
        "dob": "1987-06-04T00:00:00+09:00",
        "enrolled": [
            {
                "_id": "5d5fbbce0838642cbcb221ff",
                "year": 2019,
                "course_id": "scope_23010",
                "enroll_date": "2019-08-16T00:00:00+08:00",
                "created_at": "2019-08-23T18:11:26.112+08:00",
                "updated_at": "2019-08-23T18:11:26.112+08:00"
            }
        ],
        "created_at": "2019-08-23T13:44:07.403+08:00",
        "updated_at": "2019-08-23T16:58:52.575+08:00"
    }
}
```


--- 

## Create student info (Create)

### URL :
`<host>:<port>/api/v1/c/student` 

Request Methods : `POST`

Request body :
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `student_id`    | student id          | `string`   | yes       |
| `student_name`  | student name        | `string`   | yes       |
| `dob`   | day of birth of student | `string`      | yes       |



### example : 
`localhost:8080/api/v1/student/c`

request json body:
``` json
{
    "student_id": "54420440",
    "student_name": "Chan Tai Man",
    "dob": "1987-06-04T00:00:00+09:00",
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d5f7d2708386429880bad39",
        "student_id": "54420440",
        "student_name": "Chan Tai Man",
        "dob": "1987-06-04T00:00:00+09:00",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:47:02.411331+08:00"
    }
}
```

--- 

## Update Courses info (Update)

update the object data, by required fields

### URL : 

 `<host>:<port>/api/v1/u/student/<object_id>`

Request Methods : `POST`

Request body :
| query params | description         | data types | required?                       |
| ------------ | ------------------- | ---------- | ------------------------------- |
| `_id`        | object id of course | `string`   | yes                             |
| `student_id`    | student id           | `string`   | option, based on user necessary |
| `student_name`  | student name  | `string`   | option,based on user necessary  |
| `dob`   | day of birth of student  | `string`      | option,based on user necessary  |



### example : 
`localhost:8080/api/v1/u/student/5d5f7d2708386429880bad39`

and the request of update the dob of student only 

request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a",
    "dob":  "1987-06-05T00:00:00+09:00",
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
         "student_id": "54420440",
        "student_name": "Chan Tai Man",
        "dob": "1987-06-05T00:00:00+09:00",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:58:21.27+08:00"
    }
}
```

--- 

## Delete Courses info (Delete)

update the object data, by required fields

### URL :
`<host>:<port>/api/v1/delete/student` ,or\
 `<host>:<port>/api/v1/d/student`

Request Methods : `POST`

### Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of student | `string`   | yes       |

### example : 
`localhost:8080/api/v1/d/student/5d5f7d2708386429880bad39`
 
request json body:
``` json
{
    "_id" : "5d5f7d2708386429880bad39"
}
```

server return output (success) : 
``` json
{
    "status": 0,
    "data": true
}
```
