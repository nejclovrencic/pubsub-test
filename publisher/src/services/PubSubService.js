const Redis = require('ioredis');
const redis = new Redis();
const fetch = require('node-fetch');


const subscribeToChannel = async (channel, clientUrl) => {
    return redis.sadd(channel, clientUrl);
}

const unsubscribeFromChannel = async (channel, clientUrl) => {
    return redis.srem(channel, clientUrl);
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


module.exports = {
    subscribeToChannel,
    unsubscribeFromChannel,
    publishMessageToChannel,
}