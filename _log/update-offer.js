db.runCommand({
    "update": "Department",
    "updates": [{
        "q" : {
            "courses._id": ObjectId("5d5679b108386430ac64ea85"),
            "courses.offers._id": ObjectId("5d58ebd008386425e847d8bd"),
        },
        "u": {
            "$set": {
                "courses.$[ele].offers.$[elem]": {
                    "_id": ObjectId("5d58ebd008386425e847d8bd"),
                    "year": 2013,
                    "class_size": 40,
                    "available_places": 37,
                    "created_at": ISODate("2019-08-18T06:10:24.286Z"),
                    "updated_at": ISODate("2019-08-18T06:10:24.286Z")
                }
            }
        },
        "arrayFilters": [
            { "ele._id": { "$eq": ObjectId("5d5679b108386430ac64ea85") } },
            { "elem._id": { "$eq": ObjectId("5d58ebd008386425e847d8bd") } },
        ]
    }]
})