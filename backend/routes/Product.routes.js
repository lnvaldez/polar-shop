/* Imports */
const express = require("express");
//* Controller Functions
const {
  addProduct,
  getAllProducts,
  getProduct,
  deleteProduct,
} = require("../controllers/Product.controller");

const router = express.Router();

// POST
router.post("/", addProduct);

// GET
router.get("/", getAllProducts);
// TODO router.get("/", getAvailableProducts);
router.get("/:id", getProduct);

// PUT/PATCH
// TODO router.put("/:id", updateProduct);
// TODO router.patch("/:id", updateStock);
// TODO router.patch("/:id", updateProductImage);
// TODO router.patch("/:id", setProductAvailability);

// DELETE
router.delete("/:id", deleteProduct);

module.exports = router;
