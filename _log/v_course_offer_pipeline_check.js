db.getCollection('Department').aggregate([

    { "$unwind": "$courses" },
    { "$unwind": "$courses.offers" },
    {
        "$project": {
            "_id": "$courses.offers._id",
            "dept_obj_id": "$_id",
            "dept_id": "$dept_id",
            "dept_name": "$dept_name",
            "location": "$location",
            "course_obj_id": "$courses._id",
            "course_id": "$courses.course_id",
            "course_offer_id": {
                "$concat": ["$courses.course_id", "$courses.offers.year"]
            },
            "title": "$courses.title",
            "level": "$courses.level",
            "year": "$courses.offers.year",
            "class_size": "$courses.offers.class_size",
            "available_place": "$courses.offers.available_place",
            "num_of_stud": {
                "$subtract": ["$courses.offers.class_size", "$courses.offers.available_places"]
            }
        }
    }
])


// create 
db.run({
    "create": "VCourseOffer",
    "viewOn": "department",
    "pipeline": [

        { "$unwind": "$courses" },
        { "$unwind": "$courses.offers" },
        {
            "$project": {
                "_id": "$courses.offers._id",
                "dept_obj_id": "$_id",
                "dept_id": "$dept_id",
                "dept_name": "$dept_name",
                "location": "$location",
                "course_obj_id": "$courses._id",
                "course_id": "$courses.course_id",
                "course_offer_id": {
                    "$concat": ["$courses.course_id", "_", { "$toString": "$courses.offers.year" }]
                },
                "title": "$courses.title",
                "level": "$courses.level",
                "year": "$courses.offers.year",
                "class_size": "$courses.offers.class_size",
                "available_place": "$courses.offers.available_place",
                "num_of_stud": {
                    "$subtract": ["$courses.offers.class_size", "$courses.offers.available_places"]
                }
            }
        }
    ]
})