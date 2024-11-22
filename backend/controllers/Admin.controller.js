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
