/* Imports */
const express = require("express");

const { register, login } = require("../controllers/User.controller");

const router = express.Router();

// POST
router.post("/register", register);
router.post("/login", login);

module.exports = router;
