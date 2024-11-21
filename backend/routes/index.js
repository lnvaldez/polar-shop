/* Imports */
const express = require("express");
//* Route files
const productRoutes = require("./Product.routes");

const router = express.Router();

//* Grouped routes
router.use("/products", productRoutes);
// TODO router.use("/users", userRoutes);
// TODO router.use("/orders", orderRoutes);

module.exports = router;
