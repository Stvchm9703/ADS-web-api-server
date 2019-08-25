
// Read pipeline
db.getCollection('Student').aggregate([
    { "$unwind": "$enrolled" },
    {
        "$project": {
            "_id": 1,
            "student_id": 1,
            "student_name": 1,
            "enroll_date": "$enrolled.enroll_date",
            "course_offer_id": {
                "$concat": ["$enrolled.course_id", "_", { "$toString": "$enrolled.year" }]
            }
        }
    },
    {
        "$lookup": {
            "from": "VCourseOffer",
            "localField": "course_offer_id",
            "foreignField": "course_offer_id",
            "as": "course_offer"
        }
    },
    { "$unwind": "$course_offer" }
])


// create command
db.runCommand({
    "create": "VStudentEnrolled",
    "viewOn": "Student",
    "pipeline":
        [
            { "$unwind": "$enrolled" },
            {
                "$project": {
                    "_id": 1,
                    "student_id": 1,
                    "student_name": 1,
                    "enroll_date": "$enrolled.enroll_date",
                    "course_offer_id": {
                        "$concat": ["$enrolled.course_id", "_", { "$toString": "$enrolled.year" }]
                    }
                }
            },
            {
                "$lookup": {
                    "from": "VCourseOffer",
                    "localField": "course_offer_id",
                    "foreignField": "course_offer_id",
                    "as": "course_offer"
                }
            },
            { "$unwind": "$course_offer" }
        ]

})