const redis = require('redis');


const client = redis.createClient({
    db: 0,
    host: '127.0.0.1', // 默认 host
    port: '6379' // 默认端口
});

// 1 键值对
// client.set('color', 'green', redis.print);
// client.get('color', function (err, value) {
//     if (err) throw err;
//     console.log('Got: ' + value)
//     client.quit();
// })

// 2.哈希表
// client.hmset('kitty', {
//     'age': '2-year-old',
//     'sex': 'male'
// }, redis.print);
//
// client.hget('kitty', 'age', function (err, value) {
//     if (err) throw err;
//     console.log('kitty is ' + value);
// });
//
// client.hkeys('kitty', function (err, keys) {
//     if (err) throw err;
//     keys.forEach(function (key, i) {
//         console.log(key, i);
//     });
//     client.quit();
// });

// 3.链表
// Redis链表类似JS数组，lpush向链表中添加值，lrange获取参数start和end范围内的链表元素, 参数end为-1，表明到链表中最后一个元素。
// 注意：随着链表长度的增长，数据获取也会逐渐变慢（大O表示法中的O(n)）
// client.lpush('tasks', 'Paint the house red.', redis.print);
// client.lpush('tasks', 'Paint the house green.', redis.print);
// client.lrange('tasks', 0, -1, function (err, items) {
//     if (err) throw err;
//     items.forEach(function (item, i) {
//         console.log(' ' + item);
//     });
//     client.quit();
// });

// 4.集合
// 类似JS中的Set，集合中的元素必须是唯一的，其性能: 大O表示法中的O(1)
// client.sadd('ip', '192.168.3.7', redis.print);
// client.sadd('ip', '192.168.3.7', redis.print);
// client.sadd('ip', '192.168.3.9', redis.print);
// client.smembers('ip', function(err, members) {
//     if (err) throw err;
//     console.log(members);
//     client.quit();
// });

// 5.信道
// Redis超越了数据存储的传统职责，它还提供了信道，信道是数据传递机制，提供了发布/预定功能。
var clientA = redis.createClient(6379, '127.0.0.1')
var clientB = redis.createClient(6379, '127.0.0.1')

clientA.on('message', function (channel, message) {
    console.log('Client A got message from channel %s: %s', channel, message);
});
clientA.on('subscribe', function (channel, count) {
    clientB.publish('main_chat_room', 'Hello world!');
});
clientA.subscribe('main_chat_room');