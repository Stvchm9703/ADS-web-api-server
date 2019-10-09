# View Student Enrolled Course

base url: `<host>:<port>/api/v1/vpj/<methods>/student`

## Object property table

| variables name | data types (json) | data types (bson) | editable only? |
| -------------- | ----------------- | ----------------- | -------------- |
| `_id`          | `string`          | `ObjectId`        | `no`           |
| `student_id`   | `string`          | `string`          | `no`           |
| `student_name` | `string`          | `string`          | `no`           |
| `enroll_date`  | `string`          | `string`          | `no`           |
| `course_offer` | `object`          | `object`          | `no`           |
 

 > `course_offer` :see also View Course Offer
---

# API Usage

## List of student enrolled courses (GetList)

### URL :

`<host>:<port>/api/v1/vpj/l/student?<query>` 

Request Methods : `GET`

### Request query :
| query params      | description                  | data types                             |
| ----------------- | ---------------------------- | -------------------------------------- |
| `pn`              | `page_num`                   | `int`                                  |
| `pl`              | `page_limit`                 | `int`                                  |
| `sort`            | `sort`                       | `json string` see also : request query |
| `qbind`           | query bindings               | `string` see also : request query      |
|                   |                              |                                        |
| `student_id`      | student id                   | `string`, see also : request query     |
| `student_name`    | student name                 | `string`, see also : request query     |
| `enroll_date`     | Course Enrolled Date         | `string`, see also : request query     |
| `course_offer...` | Course Offer Object property | `string`, see also: request query      |

### example : 

`localhost:8080/api/v1/vpj/l/student/?student_id={"eq":54420440}&course_offer.year={"eq":2019}`

or

`localhost:8080/api/v1/l/student/?student_id=$eq:54420440&course_offer.year={"eq":2019}`



server will get the query as : 
``` js
{ 
    page_num : 1,
    page_limit: 50,
    $find : {
        $and:{
            "student_id": {
                $eq : "54420440"
            },
            "course_offer.year":{
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
            "_id": "5d5f7d2708386429880bad39",
            "student_id": "54420440",
            "student_name": "Chan Tai Man",
            "enroll_date": "2019-08-16T00:00:00+08:00",
            "course_offer": {
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
        }
    ]
}
```
