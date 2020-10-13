const { Joi } = require('celebrate');

const channelRequestValidator = Joi.object().keys({
  channel: Joi.string().required(),
  clientUrl: Joi.string().required(),
});

const publishRequestValidator = Joi.object().keys({
  channel: Joi.string().required(),
  message: Joi.any().required(),
});

module.exports = {
  channelRequestValidator,
  publishRequestValidator,
}
