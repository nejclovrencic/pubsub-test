const express = require('express');
const { subscribeToChannel, unsubscribeFromChannel, publishMessageToChannel } = require('../services/PubSubService');
const { channelRequestValidator, publishRequestValidator } = require('./validation/ChannelValidation');
const { celebrate } = require('celebrate');

const router = express.Router({
  caseSensitive: true,
  strict: true,
});

router.post('/subscribe', celebrate({ body: channelRequestValidator }), async (req, res, next) => {
  try {
    const body = req.body;

    await subscribeToChannel(body.channel, body.clientUrl);
    res.status(200).json(body);
  } catch (err) {
    next(err);
  }
});

router.post('/unsubscribe', celebrate({ body: channelRequestValidator }), async (req, res, next) => {
  try {
    const body = req.body;
    await unsubscribeFromChannel(body.channel, body.clientUrl);
    res.status(200).json(body);
  } catch (err) {
    next(err);
  }
});

router.post('/publish', celebrate({ body: publishRequestValidator }), async (req, res, next) => {
  try {
    const body = req.body;
    await publishMessageToChannel(body.channel, body.message);
    res.status(200).json(body);
  } catch (err) {
    next(err);
  }
});

module.exports = router;