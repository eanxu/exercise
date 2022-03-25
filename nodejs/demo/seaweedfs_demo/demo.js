'use strict';

const request = require('request');
const fs = require('fs');


function getFileFromSeaweedfs() {
    // 从seaweedfs中获取数据
    let src = "http://192.168.0.219:8888/test_grid/1/GF1B_PMS_E119.6_N29.7_20210220_L3D1228005925__f4df3d12-12d3-4d20-8685-5c7b391af008/L10/R423/C852"
    let dest = "852.png"
    request({ url: src, encoding: null}, (error, response, body) => {
        console.log(response.statusCode)
        console.log(response.headers)
        fs.writeFile(dest, body, function (err) {
            if (err) {
                console.log(err);
            } else {
                console.log('ok.');
            }
        });
    })
}

function saveFileToSeaweedfs() {
    // 向seaweedfs中存储数据
    fs.readFile('852.png', {flag:'r',encoding: null}, (error, data) => {
        if(error) {
            console.log(error)
        } else {
            var url = 'http://192.168.0.219:8888/testMvt/852.png'
            request.post({url:url, formData: {file: data}}, function (error, response, body) {
                console.log(response.statusCode)
            })
        }
    })
}

function deleteFileInSeaweedfs() {
    // 删除seaweedfs中存储的数据
    let url = 'http://192.168.0.219:8888/testMvt/852.png'
    request.delete(url, function (error, response, body) {
        console.log(response.statusCode)
    })
}

// getFileFromSeaweedfs()

// saveFileToSeaweedfs()

deleteFileInSeaweedfs()