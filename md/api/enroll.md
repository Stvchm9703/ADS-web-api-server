# Enroll

base url: `<host>:<port>/api/v1/<method>/student/<student_id>/enrolled`

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

### URL :
`<host>:<port>/api/v1/l/student/<student_obj_id>/enrolled/?<query>` 

or, get all enroll record without student object id :

`<host>:<port>/api/v1/l/student/_/enrolled/?<query>` 

Request Methods : `GET`

### Request param : 
| query params     | description          | data types | required? |
| ---------------- | -------------------- | ---------- | --------- |
| `student_obj_id` | object id of student | `string`   | option    |


### Request query :
| query params  | description            | data types                         |
| ------------- | ---------------------- | ---------------------------------- |
| `pn`          | `page_num`             | `int`                              |
| `pl`          | `page_limit`           | `int`                              |
| `year`        | student enrolling year | `int`                              |
| `course_id`   | course id              | `string`                           |
| `enroll_date` | date enrolled          | `string`, see also : request query |
| `created_at`  | date created at        | `string`, see also : request query |
| `updated_at`  | date last updated at   | `string`, see also : request query |

### example : 
`localhost:8080/api/v1/l/student/5d5f7d2708386429880bad39/enrolled`



server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "_id" : "5d5f7d2708386429880bad39"
        }
    }
}
```

output : 
``` js
{
    "SortAr": null,
    "status": 0,
    "data": [
        {
            "_id": "5d5fbbce0838642cbcb221ff",
            "year": 2019,
            "course_id": "scope_23010",
            "enroll_date": "2019-08-16T00:00:00+08:00",
            "created_at": "2019-08-23T18:11:26.112+08:00",
            "updated_at": "2019-08-23T18:11:26.112+08:00"
        }
    ]
}
```

--- 

## Get Offer info (GetOne)

### URL : 

`<host>:<port>/api/v1/g/student/<student_obj_id>/enrolled/<id>` 

Request Methods : `GET`

### Request param :
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `id`         | `object id of offer`   | `string`   | yes       |
| `student_obj_id`        | `object id of student` | `string`   | yes       |


### example : 
`localhost:8080/api/v1/g/student/5d5f7d2708386429880bad39/enrolled/5d5fbbce0838642cbcb221ff`

server will get the query as : 
``` js
{ 
    $find : {
        "_id" : "5d5f7d2708386429880bad39",
        "enrolled._id" : "5d5fbbce0838642cbcb221ff"
    }
}
```

output : 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": [
        {
            "_id": "5d5fbbce0838642cbcb221ff",
            "year": 2019,
            "course_id": "scope_23010",
            "enroll_date": "2019-08-16T00:00:00+08:00",
            "created_at": "2019-08-23T18:11:26.112+08:00",
            "updated_at": "2019-08-23T18:11:26.112+08:00"
        }
    ]
}
```


--- 

## Create Enroll record (Create)

### URL : 

`<host>:<port>/api/v1/c/student/<student_obj_id>/enrolled/` 

Request Methods : `POST`

### Request param : 
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `student_obj_id`        | `object id of student` | `string`   | yes       |


### Request body :
| query params  | description            | data types                         | required? |
| ------------- | ---------------------- | ---------------------------------- | --------- |
| `year`        | student enrolling year | `int`                              | `yes`     |
| `course_id`   | course id              | `string`                           | `yes`     |
| `enroll_date` | date enrolled          | `string`, see also : request query | `yes`     |



### example : 
`localhost:8080/api/v1/c/student/5d5f7d2708386429880bad39/enrolled/`

request json body:
``` json
{
   "year": 2019,
    "course_id": "scope_23010",
    "enroll_date": "2019-08-16T00:00:00+08:00",
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
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

### URL :
`<host>:<port>/api/v1/u/student/<student_obj_id>/enrolled/<id>` 

Request Methods : `POST`


### Request param : 
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `student_obj_id`        | `object id of student` | `string`   | yes       |
| `id`        | object id of offer | `string`   | yes       |


### Request body :
| query params  | description            | data types                         | required?                      |
| ------------- | ---------------------- | ---------------------------------- | ------------------------------ |
| `_id`         | object id of course    | `string`                           | yes                            |
| `year`        | student enrolling year | `int`                              | option,based on user necessary |
| `course_id`   | course id              | `string`                           | option,based on user necessary |
| `enroll_date` | date enrolled          | `string`, see also : request query | option,based on user necessary |



### example : 
`localhost:8080/api/v1/u/student/5d5f7d2708386429880bad39/enrolled/5d5fbbce0838642cbcb221ff`

and the request of update the year of enroll record only 

request json body:
``` json
{
    "_id" : "5d5fbbce0838642cbcb221ff",
    "year": 2018
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d5fbbce0838642cbcb221ff",
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

delete the object data, by required fields

### URL :

 `<host>:<port>/api/v1/d/student/<student_obj_id>/enrolled/<id>` 


Request Methods : `POST`

### Request param : 
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `student_obj_id`        | `object id of student` | `string`   | yes       |
| `id`        | object id of offer | `string`   | yes       |
### Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of offer | `string`   | yes       |

### example : 
`localhost:8080/api/v1/d/student/5d5f7d2708386429880bad39/enrolled/5d5fbbce0838642cbcb221ff`
 
request json body:
``` json
{
    "_id" : "5d5fbbce0838642cbcb221ff"
}
```

server return output (success) : 
``` json
{
    "status": 0,
    "data": true
}
```
