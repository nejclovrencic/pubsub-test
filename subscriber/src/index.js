if (process.env.NODE_ENV !== 'production') {
    const dotenv = require('dotenv');
    dotenv.config();
}

const bodyParser = require('body-parser');
const express = require('express');
const pubSubRouter = require('./routers/PubSubRouter');
const fetch = require('node-fetch');

const port = parseInt(process.env.APP_PORT, 10) || 9000;

const app = express();

app.use(bodyParser.json());
app.use('/api/v1', pubSubRouter);

app.listen(port, (err) => {
    if (err) {
        console.log(err);
        throw err;
    }

    console.info(`Server running on port ${port}`);
});

const channel = process.argv[2];

const res = fetch(process.env.SERVER_URL, {
    method: 'post',
    body:    JSON.stringify({ channel, clientUrl: `http://localhost:${port}/api/v1/message`}),
    headers: { 'Content-Type': 'application/json' },
}).then((res) => res.json());



