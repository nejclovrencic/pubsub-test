# PUBSUB
This project contains two sub-projects - **publisher** and **subscriber**. They simulate a simple publish/subscribe mechanism using HTTP. Under the hood, both projects run express server, and do HTTP requests using node-fetch.

### Publisher
Publisher containts three endpoints:
```
1. /subscribe - accepts channel (string) and clientUrl (string), and saves this data to redis set
2. /unsubscribe - same as above, but removes data from redis set
3. /publish - accepts channel (string) and message (any), then fetches all client urls for specified channel from redis, and uses HTTP POST to send a request body to all client urls
```

Additionally, publisher will fetch RSS Meteo feed for Europe every 60 seconds, and publish data of each city.