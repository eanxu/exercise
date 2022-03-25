var fs = require('fs');
var path = require('path');
var mbgl = require('@mapbox/mapbox-gl-native');
var sharp = require('sharp');

var options = {
    request: function(req, callback) {
        fs.readFile(req.url, function(err, data) {
            callback(err, { data: data });
        });
    },
    ratio: 1
};

var map = new mbgl.Map(options);

map.load(require('./style_json/example.json'));

map.render({zoom: 0}, function(err, buffer) {
    if (err) throw err;

    map.release();

    var image = sharp(buffer, {
        raw: {
            width: 512,
            height: 512,
            channels: 4
        }
    });

    // Convert raw image buffer to PNG
    image.toFile('./image.png', function(err) {
        if (err) throw err;
    });
});
