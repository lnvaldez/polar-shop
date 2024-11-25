/* Imports */
const express = require("express");
//* Controller functions
const {
  renderAdminDashboard,
  renderAllProductsPage,
  renderProductPage,
  renderAddProductPage,
  renderUsersPage,
  renderOrdersPage,
  renderSettingsPage,
} = require("../controllers/Admin.controller");

const router = express.Router();

//* Dashboard
router.get("/dashboard", renderAdminDashboard);

//* Products
router.get("/products", renderAllProductsPage);
router.get("/products/:id", renderProductPage);
router.get("/add/product", renderAddProductPage);

//* Users
router.get("/users", renderUsersPage);

//* Orders
router.get("/orders", renderOrdersPage);

//* Settings
router.get("/settings", renderSettingsPage);

module.exports = router;
