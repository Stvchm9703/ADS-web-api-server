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


#### Draft 1

Department Object
```js
// Department 
Department  = {
    DeptID: String,
    DeptName: String,
    Location: String,
    // Course Obj
    Courses: [{
        CourseID: String,
        Title: String,
        Level: Number,
        // Offer Obj
        Offers: [{ 
            _id: String,
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
    StudentID: String,
    StudentName: String,
    DOB: Date,
    // Enrolled Obj
    Enrolled: [{
        CourseID : String,
        Year : String,
        EnrolDate: Date
    }] 
}
```

#### Draft 2

Department Object
```js
// Department 
Department  = {
    DeptID: String,
    DeptName: String,
    Location: String,
}
```

Course-Offer-Enrolled
```js
// Course Obj
Courses={
    CourseID: String,
    Title: String,
    Level: Number,
    // Offer Obj
    Offers: [{ 
        _id: String,
        Year: Number,
        ClassSize: Number,
        Enrolled : [{
            EnrollDate : Date,
            StudentID : String,
        }]
    }] 
}

```
Student Object
```js
let Student = {
    StudentID: String,
    StudentName: String,
    DOB: Date,
}
```
----

Draft 1 is based on the Department-Course-Offer and Student-Enrolled these relationship,\
 while Draft 2 is based on the Department , Course-Offer-Enrolled , Student as the collection

Although Draft 1's Data is partrialy replicated, but consider the use-case of  