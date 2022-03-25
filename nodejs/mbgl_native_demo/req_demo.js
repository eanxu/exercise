var request = require('request');

request('http://localhost:8080/test/mvt/6/53/26', function (error, response, body) {
    if (!error && response.statusCode == 200) {
        console.log(body) // 请求成功的处理逻辑
    }
});
