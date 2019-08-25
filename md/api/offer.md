# Offer

base url: `<host>:<port>/api/v1/offer`

## Object property table

| variables name | data types (json)                | data types (bson) | editable only? |
| -------------- | -------------------------------- | ----------------- | -------------- |
| `_id`          | `string`                         | `ObjectId`        | `no`           |
| `dept_id`      | `string`                         | `string`          | `yes`          |
| `dept_name`    | `string`                         | `string`          | `yes`          |
| `location`     | `string`                         | `string`          | `yes`          |
| `created_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |
| `updated_at`   | `string` (time foramt : RFC3339) | `date`            | `no`           |

---

# API Usage

## List of Courses (GetList)

URL : \
`<host>:<port>/api/v1/dept/list?<query>` ,or\
 `<host>:<port>/api/v1/dept/l?<query>`

Request Methods : `GET`

Request query :
| query params | description            | data types                         |
| ------------ | ---------------------- | ---------------------------------- |
| `pn`         | `page_num`             | `int`                              |
| `pl`         | `page_limit`           | `int`                              |
| `dept_id`    | department id          | `string`, see also : request query |
| `dept_name`  | department name        | `string`, see also : request query |
| `location`   | location of department | `string`, see also : request query |
| `created_at` | date created at        | `string`, see also : request query |
| `updated_at` | date last updated at   | `string`, see also : request query |

for example : `localhost:8080/api/v1/dept/l?pn=1&pl=50&location={"gte":[0,1,2],"eq":1}&dept_name={"in":"database"}`

or: 

`localhost:8080/api/v1/dept/l?pn=1&pl=50&location=$gte:[0,1,2]_eq:1&dept_name=$in:database`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $or:{
            "location" : {
                $or : {
                    $eq : "CS"
                } ,
            },
            "dept_name" :{
                $or :{
                    $in : ["Computer"]
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
`<host>:<port>/api/v1/dept/get/<id>` ,or\
 `<host>:<port>/api/v1/dept/g/<id>`

Request Methods : `GET`

Request param :
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `_id`         | `object id of department` | `string`   | yes       |


for example : `localhost:8080/api/v1/dept/g/54ns123idvb`


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
        "dept_id": "CS",
        "dept_name": "Computer Science",
        "location": "red zone",
        "created_at": "2019-08-14T19:47:02.411+08:00",
        "updated_at": "2019-08-14T19:47:02.411+08:00"
    }
}
```


--- 

## Create Courses info (Create)


URL : \
`<host>:<port>/api/v1/dept/create` ,or\
 `<host>:<port>/api/v1/dept/c`

Request Methods : `POST`

Request body :
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `dept_id`    | department id          | `string`   | yes       |
| `dept_name`  | department name        | `string`   | yes       |
| `location`   | location of department | `string`      | yes       |



for example : `localhost:8080/api/v1/dept/c`

request json body:
``` json
{
    "dept_id": "CS",
    "dept_name" : "Computer Science",
    "location": "red zone"
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "dept_id": "CS",
        "dept_name" : "Computer Science",
        "location": "red zone",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:47:02.411331+08:00"
    }
}
```

--- 

## Update Courses info (Update)

update the object data, by required fields

URL : \
`<host>:<port>/api/v1/dept/update` ,or\
 `<host>:<port>/api/v1/dept/u`

Request Methods : `POST`

Request body :
| query params | description         | data types | required?                       |
| ------------ | ------------------- | ---------- | ------------------------------- |
| `_id`        | object id of course | `string`   | yes                             |
| `dept_id`    | department id           | `string`   | option, based on user necessary |
| `dept_name`  | department name  | `string`   | option,based on user necessary  |
| `location`   | location of department  | `string`      | option,based on user necessary  |



for example : `localhost:8080/api/v1/dept/u`

and the request of update the location of course only 

request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a",
    "location":  "green zone"
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
         "dept_id": "CS",
        "dept_name" : "Computer Science",
        "location": "green zone",
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:58:21.27+08:00"
    }
}
```

--- 

## Delete Courses info (Delete)

update the object data, by required fields

URL : \
`<host>:<port>/api/v1/dept/delete` ,or\
 `<host>:<port>/api/v1/dept/d`

Request Methods : `POST`

Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of course | `string`   | yes       |

for example : `localhost:8080/api/v1/dept/d`
 
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
