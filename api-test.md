
# Model Testing 

| Model      | function | checked? |
| ---------- | -------- | -------- |
| Course     | -        | -        |
| -          | Create   | OK       |
| -          | Get      | ?        |
| -          | Get List | OK       |
| -          | Update   | ?        |
| -          | Delete   | ?        |
| Department | -        | -        |
| -          | Create   | ?        |
| -          | Get      | ?        |
| -          | Get List | ?        |
| -          | Update   | ?        |
| -          | Delete   | ?        |
| Enroll     | -        | -        |
| -          | Create   | ?        |
| -          | Get      | ?        |
| -          | Get List | ?        |
| -          | Update   | ?        |
| -          | Delete   | ?        |

``` js
db.inventory.find( { tags: ["red", "blank"] } )
```

req. params &map[
    level:[
        map[
            $in:[0 1 2]
        ]
    ]
]

{
    "level" : [{
        "$in" : [1,1,0,2]
    }]
}

req. params &map[
    level:map[
        $in:[1 1 0 2]
    ]
]

```json
{
    "level" : {
        "$in" : [1,1,0,2]
    }
}
```