/* Imports */
const express = require("express");

const { register, login, logout } = require("../controllers/User.controller");

const router = express.Router();

// POST
router.post("/register", register);
router.post("/login", login);

// GET
router.get("/logout", logout);

module.exports = router;
