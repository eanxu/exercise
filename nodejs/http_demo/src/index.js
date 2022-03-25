let http = require('http');
let util = require('util');
let url = require('url');
let zlib = require('zlib');
let fs = require('fs');
let path = require('path');
let querystring = require('querystring');
let ejs = require('ejs');
let template = fs.readFileSync(path.join(__dirname, 'template.html'), 'utf8');
let chalk = require('chalk');
let mime = require('mime');

let debug = require('debug')('hello');
let config = require('./config');

let stat = util.promisify(fs.stat);
let readdir = util.promisify(fs.readdir);

class Server {
    constructor(command) {
        this.config = {...config, ...command} // config和命令行的内容展示
        this.template = template;
    }

    // 主要的请求处理方法
    async handleRequest(req, res) {
        let {dir} = this.config; // 需要将请求的路径和dir拼接在一起
        //如 http://localhost:8080/index.html
        let {pathname} = url.parse(req.url);
        // 如果独到的是网站小图标，就直接输出
        if (pathname === '/favicon.ico') return res.end();
        pathname = decodeURIComponent(pathname); // 对文件夹名称进行转码处理
        // p是决定文件路径
        let p = path.join(dir, pathname);
        try {
            // 判断当前路径是文件 还是文件夹
            let statObj = await stat(p);
            if (statObj.isDirectory()) {
                // 读取当前访问的目录下的所有内容 readdir 数组 把数组渲染回页面
                res.setHeader('Content-Type', 'text/html;charset=utf8')
                let dirs = await readdir(p);
                dirs = dirs.map(item => ({
                    name: item,
                    // 因为点击第二层时 需要带上第一层的路径，所有拼接上就ok了
                    href: path.join(pathname, item)
                }))
                // 渲染template.html中需要填充的内容，name是当前文件目录，arr为当前文件夹下的目录数组
                let str = ejs.render(this.template, {
                    name: `Index of ${pathname}`,
                    arr: dirs
                });

                // 响应中返回填充内容
                res.end(str);

            } else {
                // 如果不是文件夹，则直接输出文件内容
                this.sendFile(req, res, statObj, p);
            }
        } catch (e) {
            debug(e); // 发送错误
            this.sendError(req, res);
        }
    }

    // 处理用户缓存
    cache(req, res, statObj, p) {
        // 设置缓存头
        res.setHeader('Cache-Control', 'no-cache');
        res.setHeader('Expires', new Date(Date.now() + 10 * 1000).getTime());
        // 设置etag和上次最新修改时间
        let eTag = statObj.ctime.getTime() + '-' + statObj.size;
        let lastModified = statObj.ctime.getTime();
        // 传给客户端
        res.setHeader('Etag', eTag);
        res.setHeader('Last-Modified', lastModified);
        // 客户端把上次设置的带过来
        let ifNoneMatch = req.headers['if-none-match'];
        let ifModifiedSince = req.headers['if-modified-since'];
        // 其中任意一个不生效缓存就不生效
        if (eTag !== ifNoneMatch && lastModified !== ifModifiedSince) {
            return false;
        }

        return true;
    }

    // 是否压缩
    gzip(req, res, statObj, p) {
        // 判断请求头是否设置了接收编码
        let encoding = req.headers['accept-encoding'];
        // 如果有则判断是否有gzip或者deflate
        if (encoding) {
            // gzip
            if (encoding.match(/\bgzip\b/)) {
                res.setHeader('Content-Encoding', 'gzip');
                return zlib.createGzip();
            }
            // deflate
            if (encoding.match(/\bdeflate\b/)) {
                res.setHeader('Content-Encoding', 'deflate');
                return zlib.createDeflate();
            }
            return false;
        }
        else {
            return false;
        }
    }

    // 判断是否有范围请求
    range(req, res, statObj, p) {
        let range = req.headers['range'];
        // 有范围请求时返回读流，断点续传
        if (range) {
            let [, start, end] = range.match(/bytes=(\d*)-(\d*)/);
            start = start ? Number(start) : 0;
            end = end ? Number(end) : statObj.size - 1;
            res.statusCode = 206;
            res.setHeader('Content-Range', `bytes ${start}-${end}/${statObj.size - 1}`);
            fs.createReadStream(p, {start, end}).pipe(res);
        }
        else {
            return false;
        }
    }

    // 发送文件
    sendFile(req, res, statObj, p) {
        if (this.cache(req, res, statObj, p)) {
            res.statusCode = 304;
            return res.end();
        }
        // 是范围请求就忽略
        if (this.range(req, res, statObj, p)) return;
        // 设置文件类型头，如果不设置，我们访问一个html文件可能会导致下载
        res.setHeader('Content-Type', mime.getType(p) + ';charset=utf8');
        // 如果是需要压缩则定义gzip转化流，讲文件压缩后输出
        let transform = this.gzip(req, res, statObj, p);
        if (transform) {
            return fs.createReadStream(p).pipe(transform).pipe(res);
        }
        // 如果不是不需要压缩则直接返回文件
        fs.createReadStream(p).pipe(res);
    }

    // 处理报错
    sendError(req, res){
        // 返回的状态码设置为404
        res.statusCode = 404;
        // 页面返回文字
        res.end(`404 Not Found`);
        this.start();
    }

    // 启动方法
    start() {
        let server = http.createServer(this.handleRequest.bind(this));
        server.listen(this.config.port, this.config.host, ()=> {
            console.log(`server start http://${this.config.host}:${chalk.green(this.config.port)}`);
        });
    }
}

module.exports = Server

