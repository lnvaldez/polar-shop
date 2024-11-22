/* Imports */
const express = require("express");
//* Controller functions
const {
  renderAllProductsPage,
  renderProductPage,
} = require("../controllers/Admin.controller");

const router = express.Router();

// GET
router.get("/products", renderAllProductsPage);
router.get("/products/:id", renderProductPage);

module.exports = router;
