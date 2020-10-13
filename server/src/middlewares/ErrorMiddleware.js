const { isCelebrateError } = require('celebrate');

handleError = (err) => {
    if (isCelebrateError(err)) {
        const details = err.details.get('body') || err.details.get('params');
        
        let errorDetails;
        if (details && Array.isArray(details)) {
            errorDetails = details.map((item) => {
                let { message } = item;
                const { path } = item;
                message = message.replace(new RegExp('"', 'g'), '');
          
                return { message, path };
            });
        }
        
        return { error: 'INVALID_PARAMS', message: errorDetails || 'Unknown error', status: 400 };
    } else {
        return { error: 'INTERNAL_SERVER_ERROR', message: 'Unknown error', status: 500 };
    }
}

errorMiddleware = (err, _req, res, _next) => {
    const error = handleError(err);
    res.status(error.status);
    res.json(error);
}


module.exports = { 
    errorMiddleware,
    handleError,
 };