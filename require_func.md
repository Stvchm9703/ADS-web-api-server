## Query Functions Implementation (10 marks)

required
```js
Departments ( DeptID, DeptName, Location );
Courses ( CourseID, Title, Level );

Offer( DeptID, CourseID, Year, ClassSize, AvailablePlaces );

Students ( StudentID, StuName, DOB );

Enrolled( StudentID, Year, CourseID, EnrolDate );
```
Consider the function requirement, and data dependance, to design the document \
Plus, depend the API usage 

`Department > Course > Offer` \
and \
`Student > Enrolled -> course`
```js
let dept = {
    _id: "object_id",
    DeptID: "CS",
    DeptName: "Computer Science",
    Location: "red zone",
    Courses: [{
        CourseID: "scope_29010",
        Title: "scope introduction",
        Level: 3,
        Offers: [{
            _id: "object_id",
            Year: 2017,
            ClassSize: 40,
            AvailablePlaces: 40
        }] // Offer Obj
    }], // Course Obj
}

let Student = {
    _id: "object_id",
    StudentID: "54403124",
    StudentName: "Chan Tai Man",
    DOB: Date,
    Enrolled: [{
        CourseID : String,
        Year : String,
        EnrolDate: Date
    }] // Enrolled Obj
}
```


Moreover, your system should be able to answer the following queries:
- a) Find the titles of courses offered by the CS department in 2016.

** add `search/sort` at `dept/_id/course` OK

``` js
db.getCollection("dept").find({ DeptID : { $eq : "CS", }, "Courses.Offers.Year" : { $eq : 2016 } }, { "Courses.title" : 1, "_id" : 0 })
db.getCollection("dept").aggregate([ { $find:{ DeptID : { $eq : "CS", }, "Courses.Offers.Year" : { $eq : 2016 } } }, { $project : { "Courses.title" : 1, "_id" : 0 } } ])
```
- b) List the information of courses offered by the CS or IS departments in 2016.

page usage : `dept/?_id/cousre/` or `index page`

/api/v1/l/dept/{req Query}
 
``` js
db.getCollection("dept").aggregate([ { $find : { DeptID : { $in : ["CS", "IS"], }, "Courses.Offers.Year" : { $eq : 2016 } } }, { $project : { "_id" : 0, "Courses" : 1 } } ])
```

- c) Find the information of the course which is the most popular course enrolled by students.
  - `(Enrolled).($count=(CourseID), $sort=($count:-1)) => (Enrolled).(CourseID) == (Course).(CourseID) => (Course)`
    - Index Page

```js 
db.getCollection("dept").find({}, { "Courses" : 1, "_id" : 0 }).sort({ "Courses.Offers.AvailablePlaces" : 1 })

db.getCollection("dept").aggregate([ {$project : { "Courses" : 1, "_id" : 0 }}, {$sort:{ "Courses.Offers.AvailablePlaces" : 1 }} ])

```
  
- d) List the numbers of students for each course, who have enrolled the course offered by the CS department in 2016.
    - Project : "offer.ClassSize - Offer.Available"
```js
db.getCollection("dept").aggregate([
    { $find : { "dept_id" : "CS", "Courses.Offers.Year" : 2016 } },
    { $project : { _id : "Courses.CourseID", Courses : 1, NumOfStudent : { $subtract: [ "$Courses.Offers.ClassSize" , "$Courses.Offers.AvailablePlaces" ] } } }
])

```
- e) List the courses offered by the CS department that the student Chan Tai Man has enrolled in 2016.
  - Search Filter / Sort List in "student/Enrolled"


```js
db.getCollection("student").aggregate([
    { $find : { "StudentName" : "Chan Tai Man" } }, 
    { $project : { "_id" : 1, "StudentId" : 1, "StudentName" : 1, "CourseOfferID" : $as { }, "EnrolledDate" : "$Enrolled.EnrollDate", } },
    { $lookup : { "localField" : "Enrolled.CourseOfferID", "from":"department", "foreignField" : "department.Courses.Offers._id" "" } },
    { $project : { "Enrolled.CourseOfferID":1 } }
])
```