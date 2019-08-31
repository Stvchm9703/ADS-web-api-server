# Offer

base url: `<host>:<port>/api/v1/<method>/course/<course_id>/offer`

## Object property table

| variables name     | data types (json)                | data types (bson) | editable only? |
| ------------------ | -------------------------------- | ----------------- | -------------- |
| `_id`              | `string`                         | `ObjectId`        | `no`           |
| `year`             | `int`                            | `int`             | `yes`          |
| `class_size`       | `int`                            | `int`             | `yes`          |
| `available_places` | `int`                            | `int`             | `yes`          |
| `created_at`       | `string` (time foramt : RFC3339) | `date`            | `no`           |
| `updated_at`       | `string` (time foramt : RFC3339) | `date`            | `no`           |

---

# API Usage

## List of Course Offers (GetList)

get the list of course offer

### URL :
`<host>:<port>/api/v1/list/course/<course_id>/offer` 

or\
 `<host>:<port>/api/v1/l/course/<course_id>/offer`

Request Methods : `GET`

### Request query :
| query params       | description                             | data types                         |
| ------------------ | --------------------------------------- | ---------------------------------- |
| `pn`               | `page_num`                              | `int`                              |
| `pl`               | `page_limit`                            | `int`                              |
| `year`             | department name                         | `int`, see also : request query    |
| `available_places` | available places of course in that year | `int`, see also : request query    |
| `class_size`       | class size of that year                 | `int`, see also : request query    |
| `created_at`       | date created at                         | `string`, see also : request query |
| `updated_at`       | date last updated at                    | `string`, see also : request query |

### example : 
`localhost:8080/api/v1/l/course/5d5eb78608386435cc05f519/offer?year={"eq":2013}`

or: 

`localhost:8080/api/v1/l/course/5d5eb78608386435cc05f519/offer?year=$eq:2013`

server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $or:{
            "courses._id":"5d5eb78608386435cc05f519"
            "courses.offers.year" :{
                $eq : 2013
            },
        }
    }
}
```

output : 
``` json
{
    "page_limit": 50,
    "page_num": 1,
    "status": 0,
    "data": [
        {
            "_id": "5d5ef1fa08386444046cc5bf",
            "year": 2013,
            "class_size": 40,
            "available_places": 38,
            "created_at": "2019-08-23T03:50:18.046+08:00",
            "updated_at": "2019-08-23T03:50:18.046+08:00"
        },
    ]
}
```

--- 

## Get Course Offer info (GetOne)

### URL : 
`<host>:<port>/api/v1/g/course/<course_obj_id>/offer/<id>` 

,or\
 `<host>:<port>/api/v1/get/course/<course_obj_id>/offer/<id>`

Request Methods : `GET`

### Request param :
| query params    | description         | data types | required? |
| --------------- | ------------------- | ---------- | --------- |
| `_id`           | object id of offer  | `string`   | yes       |
| `course_obj_id` | object id of course | `string`   | yes       |

### example : 
`localhost:8080/api/v1/g/course/5d657cb70838643ac8890b93/offer/5d5ef1fa08386444046cc5bf`


server will get the query as : 
``` js
{ 
    $find : {
        "courses._id" : "5d657cb70838643ac8890b93",
        "courses.offers._id":"5d5ef1fa08386444046cc5bf"
    }
}
```

output : 
``` json
{
    "SortAr": null,
    "status": 0,
    "data": {
        "_id": "5d5ef1fa08386444046cc5bf",
        "year": 2013,
        "class_size": 40,
        "available_places": 38,
        "created_at": "2019-08-23T03:50:18.046+08:00",
        "updated_at": "2019-08-23T03:50:18.046+08:00"
    }
}
```


--- 

## Create Courses info (Create)

Creat the Offer Object under Course 

### URL :

`<host>:<port>/api/v1/create/course/<course_obj_id>/offer` 

,or


 `<host>:<port>/api/v1/c/course/<course_obj_id>/offer`

Request Methods : `POST`
 
### Request param:
| query params | description           | data types | required? |
| ------------ | --------------------- | ---------- | --------- |
| `course_obj_id`      | `object id of course` | `string` | yes |

### Request body :
| query params       | description               | data types | required? |
| ------------------ | ------------------------- | ---------- | --------- |
| `year`             | year of Offer             | `int`      | `yes`     |
| `class_size`       | class size of offer       | `int`      | `yes`     |
| `available_places` | available places of offer | `int`      | `yes`     |


for example : `localhost:8080/api/v1/c/course/5d657cb70838643ac8890b93/offer`

request json body:
``` json
{
    "year": 2014,
    "class_size": 40,
    "available_places":40,
}
```

server return output (success) : 
- server will auto generate the object id (`_id`) , created time (`created_at`) and last updated time (`updated_at`)
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
        "year": 2014,
        "class_size": 40,
        "available_places":40,
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:47:02.411331+08:00"
    }
}
```

--- 

## Update Offer info (Update)

update the object data, by required fields

### URL :
`<host>:<port>/api/v1/update/course/<course_obj_id>/offer/<id>` 

,or

 `<host>:<port>/api/v1/u/course/<course_obj_id>/offer/<id>`

Request Methods : `POST`

### Request param :
| query params    | description         | data types | required? |
| --------------- | ------------------- | ---------- | --------- |
| `_id`           | object id of offer  | `string`   | yes       |
| `course_obj_id` | object id of course | `string`   | yes       |


### Request body :
| query params | description            | data types | required?                       |
| ------------ | ---------------------- | ---------- | ------------------------------- |
| `_id`        | object id of course    | `string`   | yes                             |
| `year`             | year of Offer             | `int`      | option,based on user necessary    |
| `class_size`       | class size of offer       | `int`      | option,based on user necessary    |
| `available_places` | available places of offer | `int`      | option,based on user necessary    |

### example :

`localhost:8080/api/v1/u/course/5d657cb70838643ac8890b93/offer/5d53f4b6df86c026f2e1c64a`


request json body:
``` json
{
    "_id" : "5d53f4b6df86c026f2e1c64a",
    "year":  2015
}
```

server return output (success) : 
- server will auto update the last updated time
``` json
{
    "status": 0,
    "data": {
        "_id": "5d53f4b6df86c026f2e1c64a",
         "year": 2015,
        "class_size": 40,
        "available_places": 38,
        "created_at": "2019-08-14T19:47:02.411331+08:00",
        "updated_at": "2019-08-14T19:58:21.27+08:00"
    }
}
```

--- 

## Delete Offer info (Delete)

delete the object data, by required fields

### URL :
`<host>:<port>/api/v1/delete/course/<course_obj_id>/offer/<id>` 

,or

 `<host>:<port>/api/v1/d/course/<course_obj_id>/offer/<id>`


Request Methods : `POST`

### Request param :
| query params    | description         | data types | required? |
| --------------- | ------------------- | ---------- | --------- |
| `_id`           | object id of offer  | `string`   | yes       |
| `course_obj_id` | object id of course | `string`   | yes       |


### Request body :
| query params | description         | data types | required? |
| ------------ | ------------------- | ---------- | --------- |
| `_id`        | object id of course | `string`   | yes       |

for example : `localhost:8080/api/v1/d/course/5d657cb70838643ac8890b93/offer/5d5ef1fa08386444046cc5bf`
 
request json body:
``` json
{
    "_id" : "5d5ef1fa08386444046cc5bf"
}
```

server return output (success) : 
``` json
{
    "status": 0,
    "data": true
}
```
