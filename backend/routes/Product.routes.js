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
router.get("/", getAvailableProducts); // TODO
router.get("/:id", getProduct);

// PUT/PATCH
router.put("/:id", updateProduct); // TODO
router.patch("/:id", updateStock); // TODO
router.patch("/:id", updateProductImage); // TODO
router.patch("/:id", setProductAvailability); // TODO

// DELETE
router.delete("/:id", deleteProduct); // TODO

module.exports = router;
