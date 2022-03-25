var http = require('http');
http.createServer(function (request, response) {

    //编写返回头
    response.writeHead(200, {'Content-Type': 'text/plain;charset=UTF-8'});
    //关于content-type的，不解释，各位大佬自己去百度百度

    response.writeHead(200, {'Content-Type': 'text/html;charset=UTF-8'});
    response.write('<h1>后台已经接收到你的请求</h1>');
    //页面打印信息
//  结束响应，告诉客户端所有消息已经发送。当所有要返回的内容发送完毕时，该函数必须被调用一次。
//如何不调用该函数，客户端将永远处于等待状态。
    response.end();
}).listen(3000);
var date=new Date();
console.log("程序已启动,启动时间："+date.getTime());