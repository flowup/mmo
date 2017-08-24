var swaggermerge = require('swagger-merge');
var fs = require("fs");

var info = {
    version: "0.0.1",
    title: "Swagger",
    description: "All services\n"
};

var schemes = ['http', 'https'];

swaggermerge.on('warn', function (msg) {
    console.log(msg)
});

var swaggers = [];

process.argv.forEach(function (val, index, array) {
    if (index > 1) {
        swaggers[index - 2] = require('/tmp/' + val + '/proto.swagger.json')
    }
});

merged = swaggermerge.merge(swaggers, info, '/', 'host', schemes)
fs.writeFile('/tmp/swagger.json', JSON.stringify(merged), 'utf8', function(err) {
    if (err) {
        return console.error(err);
    }
});
