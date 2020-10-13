const { handleError } = require('./../src/middlewares/ErrorMiddleware');
const { CelebrateError, Segments, Joi } = require('celebrate');

test('returns INVALID PARAMS, when error originates from celebrate', () => {
    const result = Joi.string().valid('foo').validate('bar', { abortEarly: false });
    const err = new CelebrateError(undefined, { celebrated: true });
    err.details.set(Segments.BODY, result.error);
   
    const expected = { error: 'INVALID_PARAMS', message: 'Unknown error', status: 400 };
    const actual = handleError(err);
    
    expect(expected).toStrictEqual(actual);
});

test('returns 500 error, when error is unknown', () => {
    const err = new Error();
   
    const expected = { error: 'INTERNAL_SERVER_ERROR', message: 'Unknown error', status: 500 };
    const actual = handleError(err);
    
    expect(expected).toStrictEqual(actual);
});