/* Imports */
const express = require("express");
//* Controller Functions
const {
  addProduct,
  getAllProducts,
  getProduct,
  increaseProductStockByOne,
  setProductAsAvailable,
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
router.patch("/stock/:id", increaseProductStockByOne);
// TODO router.patch("/:id", updateProductImage);
router.patch("/status/:id", setProductAvailability);

// DELETE
router.delete("/:id", deleteProduct);

module.exports = router;
