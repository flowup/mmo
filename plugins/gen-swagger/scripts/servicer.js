var fs = require("fs");prefix = process.argv[2];

jsonpath = '/tmp/' + prefix + '/proto.swagger.json';

var swagger = require(jsonpath);

for (var key in swagger.paths) {
    if (swagger.paths.hasOwnProperty(key)) {
        swagger.paths['/' + prefix + key] = swagger.paths[key];
        delete swagger.paths[key];
    }
}

fs.writeFile(jsonpath, JSON.stringify(swagger), 'utf8', function(err) {
    if (err) {
        return console.error(err);
    }
});
