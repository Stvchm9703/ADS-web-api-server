db.getCollection('Department').aggregate([ 
{
    "$match" : {
        courses : {
            "$elemMatch" : {
                "level" : {
                    "$in" : [1]
                },
                "_id" :{
                    "$in" : [ ObjectId("5d5679b108386430ac64ea85") ]
                },
            }
        }   
    }
},
{ "$unwind" : "$courses" }
])