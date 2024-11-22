/* Imports */
const express = require("express");
//* Controller Functions
const {
  addProduct,
  getAllProducts,
  getAvailableProducts,
  getProduct,
  updateTotalStock,
  increaseStockByOne,
  decreaseStockByOne,
  setProductAvailability,
  deleteProduct,
} = require("../controllers/Product.controller");

const router = express.Router();

// POST
router.post("/", addProduct);

// GET
router.get("/", getAllProducts);
router.get("/available", getAvailableProducts);
router.get("/:id", getProduct);

// PUT/PATCH
// TODO router.put("/:id", updateProduct);
router.patch("/stock/incr/:id", increaseStockByOne);
router.patch("/stock/dec/:id", decreaseStockByOne);
router.patch("/stock/:id", updateTotalStock);
// TODO router.patch("/:id", updateProductImage);
router.patch("/status/:id", setProductAvailability);

// DELETE
router.delete("/:id", deleteProduct);

module.exports = router;
