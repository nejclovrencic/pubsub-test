const express = require('express');

const router = express.Router({
  caseSensitive: true,
  strict: true,
});

router.post('/message', async (req, res) => {
  try {
    console.log(req.body);
    res.sendStatus(200);
  } catch (err) {
    console.log(err);
  }
});

module.exports = router;