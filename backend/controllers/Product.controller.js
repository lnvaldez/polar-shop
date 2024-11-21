const mongoose = require("mongoose");

const { Product } = require("../models");

/* CRUD */

//* ..Create
const addProduct = async (req, res) => {
  const { name, price, description } = req.body;

  let emptyFields = [];

  if (!name) emptyFields.push("name");
  if (!price) emptyFields.push("price");

  if (emptyFields.length > 0) {
    return res
      .status(400)
      .json({ error: "Fill all required fields ", emptyFields });
  }

  try {
    const product = await Product.create({ name, price, description });
    res.status(200).json(product);
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

//* ..Read
const getAllProducts = async (req, res) => {
  try {
    const products = await Product.find().sort({ createdAt: -1 });
    res.status(200).json(products);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

const getProduct = async (req, res) => {
  const { id } = req.params;

  if (!mongoose.Types.ObjectId.isValid(id)) {
    return res.status(404).json({ error: "Product not found" });
  }

  const product = await Product.findById({ _id: id });

  if (!product) {
    return res.status(404).json({ error: "Product not found" });
  }

  res.status(200).json(product);
};

//* ..Update

//* ..Delete
const deleteProduct = async (req, res) => {
  const { id } = req.params;

  if (!mongoose.Types.ObjectId.isValid(id)) {
    return res.status(404).json({ error: "Product not found" });
  }

  const product = await Product.findOneAndDelete({ _id: id });

  if (!product) {
    return res.status(404).json({ error: "Product not found" });
  }

  res.status(204);
};

module.exports = {
  addProduct,
  getAllProducts,
  getProduct,
  deleteProduct,
};
