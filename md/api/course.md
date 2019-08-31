# Course

base url: `<host>:<port>/api/v1/<method>/dept/<dept_obj_id>/course`

`<dept_obj_id>` : is the object id of department object


## Object property table

| variables name | data types (json)                | data types (bson) | editable only? |
| -------------- | -------------------------------- | ----------------- | -------------- |
| `_id`          | `string`                         | `ObjectId`        | `no`           |
| `course_id`    | `string`                         | `string`          | `yes`          |
| `title`        | `string`                         | `string`          | `yes`          |
| `level`        | `string`                         | `string`          | `yes`          |
| `offer`| `array object` | `array object` | no |
| `created_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |
| `updated_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |

---

# API Usage

## List of Courses (GetList)

list out the courses of the department

### URL :

common usage

`<host>:<port>/api/v1/list/dept/<dept_obj_id>/course?<query>` ,or\
 `<host>:<port>/api/v1/l/dept/<dept_obj_id>/course?<query>`

if you want to get all course without department object id, you can use:

`<host>:<port>/api/v1/list/dept/_/course/<query>` ,or\
 `<host>:<port>/api/v1/l/dept/_/course/?<query>`

Request Methods : `GET`

Request query :
| query params | description          | data types                         |
| ------------ | -------------------- | ---------------------------------- |
| `pn`         | `page_num`           | `int`                              |
| `pl`         | `page_limit`         | `int`                              |
| `sort` | `sort`|  `json string`, request query |
| `course_id`  | `course_id`          | `string`, see also : request query |
| `title`      | title of course      | `string`, see also : request query |
| `level`      | level of course      | `string`, see also : request query |
| `created_at` | date created at      | `string`, see also : request query |
| `updated_at` | date last updated at | `string`, see also : request query |

### example 

Get the course information list in department(object_id : 5d5d822e0838640430d2fb77)

`localhost:8080/api/v1/l/dept/5d5d822e0838640430d2fb77/course?level={"eq":5}`

or: 

`localhost:8080/api/v1/l/dept/5d5d822e0838640430d2fb77/course?level=$eq:5`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "_id" : "5d5d822e0838640430d2fb77",
            "course_id.level" : {
                $eq : 5
            },
        }
    }
}
```

output : only export the course in match conditions in 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": [
        {
            "_id": "5d657cb70838643ac8890b93",
            "course_id": "scope_29109",
            "title": "developer rule",
            "level": 5,
            "offers": [
                {
                    "_id": "5d6582070838643f3849f0f3",
                    "year": 2016,
                    "class_size": 40,
                    "available_places": 20,
                    "created_at": "2019-08-28T03:18:31.816+08:00",
                    "updated_at": "2019-08-28T03:18:45.56+08:00"
                }
            ],
            "created_at": "2019-08-28T02:55:51.499+08:00",
            "updated_at": "2019-08-28T02:55:51.499+08:00"
        }
    ]
}
```

Or, 
Get the course information list whatever in which department

`localhost:8080/api/v1/l/dept/_/course?level={"eq":5}`

or: 

`localhost:8080/api/v1/l/dept/_/course?level=$eq:5`

--- 

## Get Courses info (GetOne)


### URL : 
`<host>:<port>/api/v1/get/dept/<dept_obj_id>/course/<id>` ,or\
 `<host>:<port>/api/v1/g/dept/<dept_obj_id>/course/<id>`

Request Methods : `GET`

### Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `id`         | `object id of course` | `string`   | yes       |
| `dept_obj_id`      | `object id of department` | `string` | yes |

### example : 

Get the course information in department

`localhost:8080/api/v1/g/dept/5d5d822e0838640430d2fb77/course/5d657cb70838643ac8890b93`


server will get the query as : 
``` js
{ 
    $find : {
        "_id" : "5d5d822e0838640430d2fb77",
        "courses._id":"5d657cb70838643ac8890b93"
    }
}
```

output : 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": {
        "_id": "5d657cb70838643ac8890b93",
        "course_id": "scope_29109",
        "title": "developer rule",
        "level": 5,
        "offers": [
            {
                "_id": "5d6582070838643f3849f0f3",
                "year": 2016,
                "class_size": 40,
                "available_places": 20,
                "created_at": "2019-08-28T03:18:31.816+08:00",
                "updated_at": "2019-08-28T03:18:45.56+08:00"
            }
        ],
        "created_at": "2019-08-28T02:55:51.499+08:00",
        "updated_at": "2019-08-28T02:55:51.499+08:00"
    }
}
```


--- 

## Create Courses info (Create)


### URL :

`<host>:<port>/api/v1/c/dept/<dept_obj_id>/course` ,or\
 `<host>:<port>/api/v1/create/dept/<dept_obj_id>/course`

Request Methods : `POST`

### Request param:

| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `dept_obj_id`      | `object id of department` | `string` | yes |

### Request body :
| query params | description     | data types | required? |
| ------------ | --------------- | ---------- | --------- |
| `course_id`  | `course_id`     | `string`   | yes       |
| `title`      | title of course | `string`   | yes       |
| `level`      | level of course | `int`      | yes       |



for example : `localhost:8080/api/v1/c/dept/5d5d822e0838640430d2fb77/course`

request json body:
``` json
{
    "course_id": "scope_12904",
    "title" : "advance database system",
    "level": 3
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "course_id": "scope_12904",
        "title": "advance database system",
        "level": 3,
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:47:02.411331+08:00"
    }
}
```

--- 

## Update Courses info (Update)

update the object data, by required fields

### URL :
`<host>:<port>/api/v1/update/dept/<dept_obj_id>/course/<id>` ,or\
`<host>:<port>/api/v1/u/dept/<dept_obj_id>/course/<id>`

Request Methods : `POST`

### Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `id`         | `object id of course` | `string`   | yes       |
| `dept_obj_id`      | `object id of department` | `string` | yes |


### Request body :
| query params | description         | data types | required?                       |
| ------------ | ------------------- | ---------- | ------------------------------- |
| `_id`        | object id of course | `string`   | yes                             |
| `course_id`  | `course_id`         | `string`   | option, based on user necessary |
| `title`      | title of course     | `string`   | option,based on user necessary  |
| `level`      | level of course     | `int`      | option,based on user necessary  |



### example : 
`localhost:8080/api/v1/u/dept/5d5d822e0838640430d2fb77/course/5d657cb70838643ac8890b93`

and the request of update the level of course only 

request json body:
``` json
{
    "_id" : "5d657cb70838643ac8890b93",
    "level": 4
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d657cb70838643ac8890b93",
        "course_id": "scope_12904",
        "title": "advance database system",
        "level": 4,
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:58:21.27+08:00"
    }
}
```

--- 

## Delete Courses info (Delete)

update the object data, by required fields

### URL : 
`<host>:<port>/api/v1/d/dept/<dept_obj_id>/course/<id>` ,or\
 `<host>:<port>/api/v1/delete/dept/<dept_obj_id>/course/<id>`

Request Methods : `POST`

### Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `id`         | `object id of course` | `string`   | yes       |
| `dept_obj_id`      | `object id of department` | `string` | yes |


### Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of course | `string`   | yes       |




for example : `localhost:8080/api/v1/d/dept/5d5d822e0838640430d2fb77/course/5d657cb70838643ac8890b93`
 
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
