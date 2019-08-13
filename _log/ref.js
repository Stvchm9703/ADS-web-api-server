var Order = [
{ "_id" : 1, "item" : "abc", "price" : NumberDecimal("12.00"), "quantity" : 2 },
{ "_id" : 2, "item" : "jkl", "price" : NumberDecimal("20.00"), "quantity" : 1 },
{ "_id" : 3, "item" : "abc", "price" : NumberDecimal("10.95"), "quantity" : 5 },
{ "_id" : 4, "item" : "xyz", "price" : NumberDecimal("5.95"), "quantity" : 5 },
{ "_id" : 5, "item" : "xyz", "price" : NumberDecimal("5.95"), "quantity" : 10 },
]


var inventory = [
{ "_id": 1, "sku": "abc", description: "product 1", "instock": 120 },
{ "_id": 2, "sku": "def", description: "product 2", "instock": 80 },
{ "_id": 3, "sku": "ijk", description: "product 3", "instock": 60 },
{ "_id": 4, "sku": "jkl", description: "product 4", "instock": 70 },
{ "_id": 5, "sku": "xyz", description: "product 5", "instock": 200 }
]

let proj_cmd = 
    {
        $project: {
            _id: 1,
            email: 1,
            userName: 1,
            userPhone: "$user_info.phone",
            role: "$user_role.role",
        }
    }

db.createView(
    "orderDetails",
    "orders",
    [
        {
            $lookup: {
                localField: "item", // (orders.item == inventory.sku) as inventory_docs 
                from: "inventory",
                    foreignField: "sku", // inventory.sku
                as: "inventory_docs"
            }
        },
        {
            $project: {
                "inventory_docs._id": 0, // off column 
                "inventory_docs.sku": 0
            }
        }
    ]
)


let result = [{
    "_id": 1,
    "item": "abc",
    "price": NumberDecimal("12.00"),
    "quantity": 2,
    "inventory_docs": [{ "description": "product 1", "instock": 120 }]
},
{
    "_id": 2,
    "item": "jkl",
    "price": NumberDecimal("20.00"),
    "quantity": 1,
    "inventory_docs": [{ "description": "product 4", "instock": 70 }]
}
,{
    "_id": 3,
    "item": "abc",
    "price": NumberDecimal("10.95"),
    "quantity": 5,
    "inventory_docs": [{ "description": "product 1", "instock": 120 }]
}
,{
    "_id": 4,
    "item": "xyz",
    "price": NumberDecimal("5.95"),
    "quantity": 5,
    "inventory_docs": [{ "description": "product 5", "instock": 200 }]
}
,{
    "_id": 5,
    "item": "xyz",
    "price": NumberDecimal("5.95"),
    "quantity": 10,
    "inventory_docs": [{ "description": "product 5", "instock": 200 }]
}]

{
    "ok": 0.0,
    "errmsg": "Bad projection specification, cannot include fields or add computed fields during an exclusion projection: { cs: 0.0, title: \"cs.title\", level: \"cs.level\", dept: 0.0, dept_name: \"dept.dept_name\", address: \"dept.address\" }",
    "code": 40179,
    "codeName": "Location40179"
}