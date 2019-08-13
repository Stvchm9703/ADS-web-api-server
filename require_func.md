## Query Functions Implementation (10 marks)

required
```js
Departments (
    DeptID, 
    DeptName, 
    Location
);

Courses (
    CourseID, 
    Title, 
    Level
);

Offer(
    DeptID, 
    CourseID, 
    Year, 
    ClassSize, 
    AvailablePlaces
);

Students (
    StudentID, 
    StuName, 
    DOB
);

Enrolled(
    StudentID, 
    Year, 
    CourseID, 
    EnrolDate
);
```

pre-define view 

```sql
CREATE VIEW v_course_offer as
    Select 
        b.ID,
        b.CourseID,
        b.Title,
        b.Level,
        a.DeptID,
        a.Year,
        a.ClassSize,
        a.AvailablePlaces,
        c.DeptName,
        c.Address
    FROM          
        Course  b
    INNER JOIN    
        Offer   a
        ON        a.CourseID = b.CourseID
    INNER JOIN
        Departments c
        ON        c.DeptID = a.DeptID
```


Moreover, your system should be able to answer the following queries:
- a) Find the titles of courses offered by the CS department in 2016.
  - `(Offer).(DeptID = $in:"CS" && Year = $eq:2016) => (Offer).(CourseID) == (Course).(CourseID) => (Course).(Title)`
    ```sql 
    SELECT        
        Title
    FROM          
        v_course_offer
    WHERE(
        (DeptID   in 'CS') AND
        (Year     in 2016)
    )
    ```


- b) List the information of courses offered by the CS or IS departments in 2016.
  - `(Offer).(DeptID = $in:["CS","IS"] && Year = $eq:2016) => (Offer).(CourseID) == (Course).(CourseID) => (Course)`
    ```sql
    SELECT
        *  
    FROM
        v_course_offer
    WHERE(
        (DeptID   IN ('CS' , 'IS')) AND
        (Year     IN 2016)
    );
    ```


- c) Find the information of the course which is the most popular course enrolled by students.
  - `(Enrolled).($count=(CourseID), $sort=($count:-1)) => (Enrolled).(CourseID) == (Course).(CourseID) => (Course)`
    ```sql
    SELECT 
        a.CourseID,
        COUNT(a.CourseID) AS `NumOfCourse`,
        -- a.Year,
        b.Title,
        b.Level
    FROM 
        `Enrolled`  a
    LEFT JOIN
        v_course_offer b
        ON(      
            a.CourseID = b.CourseID 
            -- AND a.Year = b.Year
        )
    GROUP BY 
        a.CourseID, 
        -- a.Year
    ORDER BY 
        NumOfCourse DESC
    ;
    ```
  
- d) List the numbers of students for each course, who have enrolled the course offered by the CS department in 2016.
  - `(Enrolled).($group: { _id:{CourseID,Year} , count:{$sum:1}  }).($sort(count)) => `
    ```sql
    SELECT 
        a.CourseID,
        COUNT(a.CourseID) AS `NumOfCourse`,
        a.Year,
        b.Title,
        b.Level,
        b.DeptID
    FROM 
        `Enrolled`  a
    LEFT JOIN
        v_course_offer b
        ON(      
            a.CourseID = b.CourseID 
            AND a.Year = b.Year
        )
    WHERE
        b.DeptID = 'CS'
        AND a.Year = 2016
    GROUP BY 
        a.CourseID, 
        a.Year
    ORDER BY 
        NumOfCourse DESC
    ;

    ```

- e) List the courses offered by the CS department that the student Chan Tai Man has enrolled in 2016.
    ```sql

    SELECT 
        c.StudentID,
        c.StuName,
        b.*
    FROM 
        `Enrolled`  a
    INNER JOIN
        Student c
        ON (
            a.StudentID = c.StudentID
        )
    LEFT JOIN
        `v_course_offer` b
        ON(      
            a.CourseID  = b.CourseID 
            a.Year      = b.Year
        )
    WHERE
        b.DeptID = 'CS'
        AND a.Year = 2016
        And c.StudName = 'Chan Tai Man'
    ;
    ```