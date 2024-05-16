
var db = db.getSiblingDB('shortUrl');
db.createCollection('urls');


db.urls.createIndex({ url: 1 }, { unique: true });
db.urls.createIndex({ key: 1 }, { unique: true });