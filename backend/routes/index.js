/* Imports */
const express = require("express");
//* Route files
const productRoutes = require("./Product.routes");
const userRoutes = require("./User.routes");
const orderRoutes = require("./Order.routes");
const adminRoutes = require("./Admin.routes");

const router = express.Router();

//* Grouped routes
router.use("/products", productRoutes);
router.use("/users", userRoutes);
router.use("/admin", adminRoutes);
router.use("/orders", orderRoutes);

module.exports = router;
