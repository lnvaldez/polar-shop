const mongoose = require("mongoose");

const { Product } = require("../models");

//* ..Render views
const renderAllProductsPage = async (req, res) => {
  try {
    const products = await Product.find().sort({ createdAt: -1 });
    res.render("admin/products", { products });
  } catch (error) {
    console.error({ error: error.message });
    res.render("/error");
  }
};

const renderProductPage = async (req, res) => {
  const id = req.params.id;

  try {
    const product = await Product.findById(id);
    res.render(`admin/products/${id}`, { product });
  } catch (error) {
    console.error({ error: error.message });
    res.render("/error");
  }
};
