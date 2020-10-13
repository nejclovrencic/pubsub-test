const Redis = require('ioredis');
const fetch = require('node-fetch');
const Parser = require('rss-parser');

const parser = new Parser();
const redis = new Redis();

const subscribeToChannel = async (channel, clientUrl) => {
    return redis.sadd(channel.toLowerCase(), clientUrl);
}

const unsubscribeFromChannel = async (channel, clientUrl) => {
    return redis.srem(channel.toLowerCase(), clientUrl);
}

const publishMessageToChannel = async (channel, message) => {
    const clientUrls = await redis.smembers(channel);
    const sendRequest = async (url, body) => {
        const res = await fetch(url, {
            method: 'post',
            body:    JSON.stringify(body),
            headers: { 'Content-Type': 'application/json' },
        });

        return res.json();
    }

    await Promise.all(clientUrls.map((url) => sendRequest(url, { message }).catch(e => e)));
}

const fetchRssAndPublishMessage = async () => {
    const feed = await parser.parseURL('http://meteo.arso.gov.si/uploads/probase/www/observ/surface/text/en/observation_eu-capital_latest.rss');
    await Promise.all(feed.items.map((item) => {
        if (item.guid) {
            const city = item.guid.split('_')[0];
            publishMessageToChannel(city.toLowerCase(), item);
        }
        
    }));
}

module.exports = {
    subscribeToChannel,
    unsubscribeFromChannel,
    publishMessageToChannel,
    fetchRssAndPublishMessage,
}