#! /usr/bin/env node
let Server = require('../src/index.js'); // 导入Server
let commander = require('commander'); // 导入命令行模块
let {version} = require('../package.json'); // 读取package.json的版本

// 配置命令行
commander
    .option('-p,--port <n>', 'config port') // 配置端口
    .option('-o,--host [value]', 'config hostname') // 配置主机名
    .option('-d,--dir [value]', 'config directory') // 配置访问目录
    .version(version, '-v,--version').parse(process.argv); // 展示版本

let server = new Server(commander);
server.start(); // 启动

let config =require('../src/config');

commander = {...config, ...commander}

let os = require('os');
// 执行模块
let {exec} = require('child_process')

// 判断操作系统平台，win32是windows，执行访问程序，会自动弹出默认浏览器喔
if (os.platform() === 'win32') {
    exec(`start http://${commander.host}:${commander.port}`);
}
else {
    exec(`open http://${commander.host}:${commander.port}`);
}
