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
const increaseStockByOne = async (req, res) => {};
const decreaseStockByOne = async (req, res) => {};
const updateTotalProductStock = async (req, res) => {}; // TODO Replaces current value

const updateProductStock = async (req, res) => {
  const { id } = req.params;
  const { stock } = req.body;

  if (!mongoose.Types.ObjectId.isValid(id)) {
    return res.status(404).json({ error: "Product not found" });
  }

  const p = await Product.findById(id);

  if (stock < 0 && p.stock == 0) {
    return res.status(405).json({ error: "Stock can't be less than zero" });
  }

  try {
    let update;

    if (stock !== undefined) {
      update = { $inc: { stock: stock } };
    } else {
      update = { $inc: { stock: 1 } };
    }

    const product = await Product.findOneAndUpdate({ _id: id }, update, {
      new: true,
    });

    if (!product) {
      return res.status(404).json({ error: "Product not found" });
    }

    return res.status(200).json(product);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

const setProductAvailability = async (req, res) => {
  const { id } = req.params;

  if (!mongoose.Types.ObjectId.isValid(id)) {
    return res.status(404).json({ error: "Product not found" });
  }

  const product = await Product.findOneAndUpdate(
    { _id: id },
    [{ $set: { available: { $eq: [false, "$available"] } } }],
    { new: true }
  );

  if (!product) {
    return res.status(404).json({ error: "Product not found" });
  }

  res.status(200).json(product);
};

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
  updateProductStock,
  setProductAvailability,
  deleteProduct,
};
