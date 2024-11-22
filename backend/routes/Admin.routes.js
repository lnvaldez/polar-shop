/* Imports */
const express = require("express");
const { Product } = require("../models");

const router = express.Router();

// GET
router.get("/products", async (req, res) => {
  const products = await Product.find();
  res.render("admin/products", { products });
});

module.exports = router;
