const Service = require('egg').Service;
var fs = require('fs');
var request = require('request');
var mbgl = require('@mapbox/mapbox-gl-native');
var sharp = require('sharp');

class TmsService extends Service {
    async xyz(x, y, z) {
        return new Promise(resolve => {
            var options = {
                request: function (req, callback) {
                    request({
                        url: req.url,
                        encoding: null,
                        gzip: true
                    }, function (err, res, body) {
                        if (err) {
                            callback(err);
                        } else if (res.statusCode === 200) {
                            const response = {};

                            if (res.headers.modified) {
                                response.modified = new Date(res.headers.modified);
                            }
                            if (res.headers.expires) {
                                response.expires = new Date(res.headers.expires);
                            }
                            if (res.headers.etag) {
                                response.etag = res.headers.etag;
                            }

                            response.data = body;

                            callback(null, response);
                        } else {
                            callback(new Error(JSON.parse(body).message));
                        }
                    });
                },
                ratio: 1
            };

            var map = new mbgl.Map(options);

            var obj = JSON.parse('{"version": 8,"sources": {"mapbox": {"type": "vector"}},"layers": [{"id": "adcode","type": "fill","source": "mapbox","source-layer": "adcode","paint": {"fill-color": "#d33046","fill-outline-color": "#000000"}}]}')

            let url = "http://localhost:8080/test/orb/"
            url = url.concat(z, "/", x, "/", y)
            obj['sources']['mapbox']['tiles'] = new Array(url)

            map.load(obj);

            map.render({}, function (err, buffer) {
                if (err) throw err;
                map.release();

                let image = sharp(buffer, {
                    raw: {
                        width: 512,
                        height: 512,
                        channels: 4
                    }
                });
                return image.png().toBuffer().then(resolve)
            })
        });
    }
}

module.exports = TmsService
