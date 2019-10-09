# View Course Offer

base url: `<host>:<port>/api/v1/vpj/<methods>/dept`



## Object property table


Object is base on the Department-Course-Offer relationship to shorten in same level

| variables name    | data types (json)         | data types (bson) | editable only? |
| ----------------- | ------------------------- | ----------------- | -------------- |
| `_id`             | `string`                  | `ObjectId`        | `no`           |
| `dept_obj_id`     | `object id of department` | `string`          | no             |
| `dept_id`         | `string`                  | `string`          | no             |
| `dept_name`       | `string`                  | `string`          | no             |
| `location`        | `string`                  | `string`          | no             |
| `course_obj_id`   | `string`                  | `string`          | no             |
| `course_id`       | `string`                  | `string`          | no             |
| `title`           | `string`                  | `string`          | no             |
| `level`           | `int`                     | `int`             | no             |
| `course_offer_id` | `string`                  | `string`          | no             |
| `year`            | `int`                     | `int`             | no             |
| `class_size`      | `int`                     | `int`             | no             |
| `num_of_stud`     | `int`                     | `int`             | no             |


---

# API Usage

## List of student (GetList)

### URL :

`<host>:<port>/api/v1/vpj/l/student?<query>` 

Request Methods : `GET`

### Request query :
| query params      | description                                                       | data types                             |
| ----------------- | ----------------------------------------------------------------- | -------------------------------------- |
| `pn`              | `page_num`                                                        | `int`                                  |
| `pl`              | `page_limit`                                                      | `int`                                  |
| `sort`            | `sort`                                                            | `json string` see also : request query |
| `qbind`           | query bindings                                                    | `string` see also : request query      |
|                   |                                                                   |                                        |
| `_id`             | object id of the offer object                                     | `string` see also : request query      |
| `dept_obj_id`     | object id of department                                           | `string`                               |
| `dept_id`         | department id                                                     | `string`                               |
| `dept_name`       | department name                                                   | `string`                               |
| `location`        | office location                                                   | `string`                               |
| `course_obj_id`   | object id of course                                               | `string`                               |
| `course_id`       | course id                                                         | `string`                               |
| `title`           | course title                                                      | `string`                               |
| `level`           | course level                                                      | `int`                                  |
| `course_offer_id` | joined key: course_id + year , for joining the `vStudentEnrolled` | `string`                               |
| `year`            | year of Course offer                                              | `int`                                  |
| `class_size`      | Course class size                                                 | `int`                                  |
| `num_of_stud`     | the number of student that joined the course                      | `int`                                  |

### example : 

`localhost:8080/api/v1/vpj/l/dept/?year={"eq":2019}`

or

`localhost:8080/api/v1/l/dept/?year={"eq":2019}`



server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "year":{
                $eq : 2019
            }   
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
            "_id": "5d5f9b6a0838643650532103",
            "dept_obj_id": "5d5d822e0838640430d2fb77",
            "dept_id": "CS",
            "dept_name": "Computer Science",
            "location": "Blue Zone",
            "course_obj_id": "5d5eb78608386435cc05f519",
            "course_id": "scope_23010",
            "title": "advanced developer rule",
            "level": 4,
            "course_offer_id": "scope_23010_2019",
            "year": 2019,
            "class_size": 40,
            "num_of_stud": 4
        }
    ]
}
```
