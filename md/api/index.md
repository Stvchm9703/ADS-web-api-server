# API back-end related

## Basic Data Structure

### Data Structure Requirement

base on the requirement, the object structure are as below : 

```python
Departments( DeptID, DeptName, Location )

Courses( CourseID, Title, Level )

Offer( DeptID, CourseID, Year, ClassSize, AvailablePlaces )

Students( StudentID, StuName, DOB )

Enrolled( StudentID, Year, CourseID, EnrolDate );
```

#### Data Dependecy

Consider the requirement and the relation of data structure.\
It show `Offer` have direct relationship to `Department` and `Course`, by the `DeptID` and `CourseID`. However, there are no direct relation bewteen `Department` and `Course` .

While the `Student` model have direct bonding with `Enrolled`, by the `StudentID`, while it also linked to `Year` , `CourseID` that having similar variables to `Courses` and `Offer`

Therefore, the model relationship can draft as below:

![draft1](./draft1.jpg)
![draft2](./draft2.jpg)


### Draft 1

Department Object
```js
// Department 
Department  = {
    _id : String<object_id>,
    DeptID: String,
    DeptName: String,
    Location: String,
    // Course Obj
    Courses: [{
        _id : String<object_id>,
        CourseID: String,
        Title: String,
        Level: Number,
        // Offer Obj
        Offers: [{ 
            _id : String<object_id>,
            Year: Number,
            ClassSize: Number,
            AvailablePlaces: Number
        }] 
    }], 
}
```

Student Object
```js
// Student Object
Student = {
    _id : String<object_id>,
    StudentID: String,
    StudentName: String,
    DOB: Date,
    // Enrolled Obj
    Enrolled: [{
    _id : String<object_id>,
    CourseID : String,
        Year : String,
        EnrollDate: Date
    }] 
}
```

### Draft 2

Department Object
```js
// Department 
Department  = {
    _id : String<object_id>,
    DeptID: String,
    DeptName: String,
    Location: String,
}
```

Course-Offer
```js
// Course Obj
Courses={
    _id : String<object_id>,
    DepartmentId : String<department.obj_id>,
    CourseID: String,
    Title: String,
    Level: Number,
    Offers:[ String<offer.obj_id>]
}
```

Offer-Enrolled
```js
Offers = { 
    _id : String<object_id>,
    Year: Number,
    ClassSize: Number,
    Enrolled : [{
        EnrollDate : Date,
        StudentID : String<student.obj_id>,
    }]
}

```
Student Object
```js
Student = {
    _id : String<object_id>,
    StudentID: String,
    StudentName: String,
    DOB: Date,
}
```
----

Draft 1 is based on the Department-Course-Offer and Student-Enrolled these relationship,\
 while Draft 2 is based on the Department, Course-Offer, Offer-Enrolled , Student as the collection

Although Draft 1's Data is partrialy replicated, but consider the future use-case in Web Application and mobile application, keeping in same document may less 