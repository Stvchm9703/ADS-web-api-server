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
