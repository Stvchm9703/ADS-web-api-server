# Department

base url: `<host>:<port>/api/v1/<methods>/dept`

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

## List of Department (GetList)

### URL :

`<host>:<port>/api/v1/l/dept?<query>` ,or\
 `<host>:<port>/api/v1/list/dept/?<query>`

Request Methods : `GET`

Request query :
| query params | description            | data types                             |
| ------------ | ---------------------- | -------------------------------------- |
| `pn`         | `page_num`             | `int`                                  |
| `pl`         | `page_limit`           | `int`                                  |
| `sort`       | `sort`                 | `json string` see also : request query |
| `qbind`      | query bindings         | `string` see also : request query      |
| `dept_id`    | department id          | `string`, see also : request query     |
| `dept_name`  | department name        | `string`, see also : request query     |
| `location`   | location of department | `string`, see also : request query     |
| `created_at` | date created at        | `string`, see also : request query     |
| `updated_at` | date last updated at   | `string`, see also : request query     |

### example : 
`localhost:8080/api/v1/l/dept/?level=$gte:[4,5]`

or: 

`localhost:8080/api/v1/l/dept/?level={"gte":[4,5]}`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "dept_id": {
                $or : {
                    $in : ["CS"]
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
            "_id": "5d5d822e0838640430d2fb77",
            "dept_id": "CS",
            "dept_name": "Computer Science",
            "location": "Blue Zone",
            "courses": [
                {
                    "_id": "5d5eb78608386435cc05f519",
                    "course_id": "scope_23010",
                    "title": "advanced developer rule",
                    "level": 4,
                    "created_at": "2019-08-22T23:40:54.734+08:00",
                    "updated_at": "2019-08-23T00:07:42.761+08:00"
                },
                {
                    "_id": "5d657cb70838643ac8890b93",
                    "course_id": "scope_29109",
                    "title": "developer rule",
                    "level": 5,
                    "created_at": "2019-08-28T02:55:51.499+08:00",
                    "updated_at": "2019-08-28T02:55:51.499+08:00"
                }
            ],
            "created_at": "2019-08-22T01:41:02.773+08:00",
            "updated_at": "2019-08-22T23:06:24.203+08:00"
        }
    ]
}
```

--- 

## Get Department info (GetOne)


### URL :

`<host>:<port>/api/v1/g/dept/<id>` ,or\
 `<host>:<port>/api/v1/get/dept/<id>`

Request Methods : `GET`

Request param :
| query params | description               | data types | required? |
| ------------ | ------------------------- | ---------- | --------- |
| `_id`        | `object id of department` | `string`   | yes       |


### example : 
`localhost:8080/api/v1/g/dept/5d5d822e0838640430d2fb77`


server will get the query as : 
``` js
{ 
    $find : {
        "_id" : "54ns5d5d822e0838640430d2fb77123idvb"
    }
}
```

output : 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": {
        "_id": "5d5d822e0838640430d2fb77",
        "dept_id": "CS",
        "dept_name": "Computer Science",
        "location": "Blue Zone",
        "courses": [
            {
                "_id": "5d5eb78608386435cc05f519",
                "course_id": "scope_23010",
                "title": "advanced developer rule",
                "level": 4,
                "offers": [
                    {
                        "_id": "5d5ef1fa08386444046cc5bf",
                        "year": 2013,
                        "class_size": 40,
                        "available_places": 38,
                        "created_at": "2019-08-23T03:50:18.046+08:00",
                        "updated_at": "2019-08-23T03:50:18.046+08:00"
                    },
                    {
                        "_id": "5d5f9b6a0838643650532103",
                        "year": 2019,
                        "class_size": 40,
                        "available_places": 36,
                        "created_at": "2019-08-23T15:53:14.315+08:00",
                        "updated_at": "2019-08-23T15:53:14.315+08:00"
                    },
                    {
                        "_id": "5d5febb4083864458c9dbff1",
                        "year": 2020,
                        "class_size": 40,
                        "available_places": 29,
                        "created_at": "2019-08-23T21:35:48.421+08:00",
                        "updated_at": "2019-08-24T08:05:09.572+08:00"
                    }
                ],
                "created_at": "2019-08-22T23:40:54.734+08:00",
                "updated_at": "2019-08-23T00:07:42.761+08:00"
            },
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
        ],
        "created_at": "2019-08-22T01:41:02.773+08:00",
        "updated_at": "2019-08-22T23:06:24.203+08:00"
    }
}
```


--- 

## Create Department info (Create)

### URL :

 `<host>:<port>/api/v1/c/dept`

Request Methods : `POST`

Request body :
| query params | description            | data types | required? |
| ------------ | ---------------------- | ---------- | --------- |
| `dept_id`    | department id          | `string`   | yes       |
| `dept_name`  | department name        | `string`   | yes       |
| `location`   | location of department | `string`   | yes       |



### example : 
`localhost:8080/api/v1/c/dept/`

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

## Update Department info (Update)

update the object data, by required fields

### URL : 

 `<host>:<port>/api/v1/u/dept/<_id>`

Request Methods : `POST`

### Request body :
| query params | description             | data types | required?                       |
| ------------ | ----------------------- | ---------- | ------------------------------- |
| `_id`        | object id of department | `string`   | yes                             |
| `dept_id`    | department id           | `string`   | option, based on user necessary |
| `dept_name`  | department name         | `string`   | option,based on user necessary  |
| `location`   | location of department  | `string`   | option,based on user necessary  |



### example : 
`localhost:8080/api/v1/u/dept/5d53f4b6df86c026f2e1c64a`

and the request of update the location of department only 

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

### URL :

 `<host>:<port>/api/v1/d/dept/<id>`

Request Methods : `POST`

### Request body :
| query params | description             | data types | required? |
| ------------ | ----------------------- | ---------- | --------- |
| `id`         | object id of department | `string`   | yes       |

example : `localhost:8080/api/v1/d/dept/5d53f4b6df86c026f2e1c64a`
 
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
