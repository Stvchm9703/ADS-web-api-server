db.getCollection('Department').update({
            "courses._id": ObjectId("5d5679b108386430ac64ea85")
    },
    {
        "$push": {
            "courses.$.offers": {
                "year": 2017,
                "class_size": 40,
                "available_places": 40
            }
        },
        "$set":{
            "courses.$.updated_at" : Date.now()
        }  
        
    }
)