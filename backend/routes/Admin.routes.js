/* Imports */
const express = require("express");
//* Controller functions
const {
  renderAdminDashboard,
  renderAllProductsPage,
  renderProductPage,
} = require("../controllers/Admin.controller");

const router = express.Router();

// GET
router.get("/dashboard", renderAdminDashboard);
router.get("/products", renderAllProductsPage);
router.get("/products/:id", renderProductPage);

module.exports = router;
