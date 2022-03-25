exports.keys = "123123";

// add middleware robot
exports.middleware = [
    'robot'
];

// robot's configurations
exports.robot = {
    ua: [
        /Baiduspider/i,
    ]
};

// //跨域配置
// config.security = {
//     csrf: {
//         enable: false, // 前后端分离，post请求不方便携带_csrf
//         ignoreJSON: true
//     },
//     domainWhiteList: ['http://www.baidu.com', 'http://localhost:8080'], //配置白名单
// };

exports.cors = {
    origin: '*',
    allowMethods: 'GET,HEAD,PUT,POST,DELETE,PATCH'
};

