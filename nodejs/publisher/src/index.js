if (process.env.NODE_ENV !== 'production') {
    const dotenv = require('dotenv');
    dotenv.config();
}

const bodyParser = require('body-parser');
const express = require('express');
const healthRouter = require('./routers/HealthRouter');
const pubSubRouter = require('./routers/PubSubRouter');
const { errorMiddleware } = require('./middlewares/ErrorMiddleware');
const { fetchRssAndPublishMessage } = require('./services/PubSubService');

const port = parseInt(process.env.APP_PORT, 10) || 3000;

const app = express();

app.use(bodyParser.json());
app.use('/health', healthRouter);
app.use('/api/v1', pubSubRouter);
app.use(errorMiddleware);

app.listen(port, (err) => {
    if (err) {
        console.log(err);
        throw err;
    }

    console.info(`Server running on port ${port}`);
    
    fetchRssAndPublishMessage();
    
    const interval = setInterval(fetchRssAndPublishMessage, 60 * 1000);
    
    process.on('SIGINT', () => {
        clearInterval(interval);
        process.exit();
    });
});


