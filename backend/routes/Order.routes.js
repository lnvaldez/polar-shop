const express = require("express");

const { approveOrder } = require("../controllers/Order.controller");

const router = express.Router();

router.post("/approve/:id", approveOrder);

module.exports = router;
