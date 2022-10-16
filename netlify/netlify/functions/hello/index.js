const { MongoClient } = require('mongodb');

exports.handler = async function (event, context) {
    const client = new MongoClient('mongodb+srv://app1:iamQsffYgOlnTKae@data-1.othux.mongodb.net/?retryWrites=true&w=majority');
    await client.connect();
    console.log('Connected successfully to server');
    const db = client.db('all');
    const collection = db.collection('test');

    let n = collection.countDocuments()

    return {
        statusCode: 200,
        body: JSON.stringify({ message: "Return from function!", "count": n }),
    };
};