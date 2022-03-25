let path = require('path')

console.log("dfasdf: ", path.dirname(path.dirname(__dirname)))

module.exports = {
    port: 8080, // 默认端口
    host: 'localhost', // 默认主机名
    dir:  path.dirname(path.dirname(__dirname)) // 默认读取目录
}
