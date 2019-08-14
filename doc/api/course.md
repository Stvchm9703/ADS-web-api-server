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


# API Usage

## List of Courses 

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

for example : `localhost:8080/api/v1/course/l?pn=1&pl=50&level={"gte":[0,1,2]}`

server will get as : 
``` javascript

{ 
    page_num : 1,
    page_limit: 50,
    level : {
        gte : [0, 1, 2]
    } 
}

```