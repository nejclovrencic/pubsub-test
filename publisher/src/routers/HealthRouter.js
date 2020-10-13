const express = require('express');

const router = express.Router({
  caseSensitive: true,
  strict: true,
});

router.post('/liveness', async (_req, res) => {
  res.sendStatus(200);
});

module.exports = router;