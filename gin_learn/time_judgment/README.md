http://ip:8010/Log/addChargLog
POST --- JSON
{
    "Request": str,
    "Resp": str
    "collectionName": str  // 必选
}

return
{
    "error": "",
    "success": true
}

http://ip:8010/Log/findChargLog
POST --- JSON
{
    "collectionName": str,  // 必选
    "$or": [
        {"resp": "ok"}  // 属性需小写
    ],
    "Sort": [    // 非必须 排序
        "resp"  // -resp 倒排序(大到小)
        ]
    "Limit":  5 // 非必须
}

return
{
    {
        "_id": "xxxxxxx",
        "resp": str,
        "requests": str
    }
}
