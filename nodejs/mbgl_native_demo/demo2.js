let fs = require('fs')

var ttt = ""

var style = fs.readFile('./style_json/adcode.json', 'utf8', function (err, data) {
    if (err) {
        console.log(err)
    }
    ttt = JSON.parse(data)
})

console.log(ttt)
