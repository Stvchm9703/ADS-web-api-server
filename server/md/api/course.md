# Course

base url: `<host>:<port>/api/v1/course`

## Object property table

| variables name | data types (json)                | data types (bson) | editable only? |
| -------------- | -------------------------------- | ----------------- | -------------- |
| `_id`          | `string`                         | `ObjectId`        | `no`           |
| `course_id`    | `string`                         | `string`          | `yes`          |
| `title`        | `string`                         | `string`          | `yes`          |
| `level`        | `string`                         | `string`          | `yes`          |
| `created_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |
| `updated_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |

---

# API Usage

## List of Courses (GetList)

URL : \
`<host>:<port>/api/v1/course/list?<query>` ,or\
 `<host>:<port>/api/v1/course/l?<query>`

Request Methods : `GET`

Request query :
| query params | description          | data types                         |
| ------------ | -------------------- | ---------------------------------- |
| `pn`         | `page_num`           | `int`                              |
| `pl`         | `page_limit`         | `int`                              |
| `course_id`  | `course_id`          | `string`, see also : request query |
| `title`      | title of course      | `string`, see also : request query |
| `level`      | level of course      | `string`, see also : request query |
| `created_at` | date created at      | `string`, see also : request query |
| `updated_at` | date last updated at | `string`, see also : request query |

for example : `localhost:8080/api/v1/course/l?pn=1&pl=50&level={"gte":[0,1,2],"eq":1}&title={"in":"database"}`

or: 

`localhost:8080/api/v1/course/l?pn=1&pl=50&level=$gte:[0,1,2]_eq:1&title=$in:database`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $or:{
            "level" : {
                $or : {
                    $gte : [0, 1, 2],
                    $eq : 1
                } ,
            },
            "title" :{
                $or :{
                    $in : ["database"]
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
`<host>:<port>/api/v1/course/get/<id>` ,or\
 `<host>:<port>/api/v1/course/g/<id>`

Request Methods : `GET`

Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `id`         | `object id of course` | `string`   | yes       |


for example : `localhost:8080/api/v1/course/g/54ns123idvb`


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
        "course_id": "scope_12904",
        "title": "advance database system",
        "level": 3,
        "created_at": "2019-08-14T19:47:02.411+08:00",
        "updated_at": "2019-08-14T19:47:02.411+08:00"
    }
}
```


--- 

## Create Courses info (Create)


URL : \
`<host>:<port>/api/v1/course/create` ,or\
 `<host>:<port>/api/v1/course/c`

Request Methods : `POST`

Request body :
| query params | description     | data types | required? |
| ------------ | --------------- | ---------- | --------- |
| `course_id`  | `course_id`     | `string`   | yes       |
| `title`      | title of course | `string`   | yes       |
| `level`      | level of course | `int`      | yes       |



for example : `localhost:8080/api/v1/course/c`

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

URL : \
`<host>:<port>/api/v1/course/update` ,or\
 `<host>:<port>/api/v1/course/u`

Request Methods : `POST`

Request body :
| query params | description         | data types | required?                       |
| ------------ | ------------------- | ---------- | ------------------------------- |
| `_id`        | object id of course | `string`   | yes                             |
| `course_id`  | `course_id`         | `string`   | option, based on user necessary |
| `title`      | title of course     | `string`   | option,based on user necessary  |
| `level`      | level of course     | `int`      | option,based on user necessary  |



for example : `localhost:8080/api/v1/course/u`

and the request of update the level of course only 

request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a",
    "level": 4
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
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

URL : \
`<host>:<port>/api/v1/course/delete` ,or\
 `<host>:<port>/api/v1/course/d`

Request Methods : `POST`

Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of course | `string`   | yes       |




for example : `localhost:8080/api/v1/course/d`
 
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
