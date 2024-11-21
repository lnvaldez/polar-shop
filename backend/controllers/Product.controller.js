const mongoose = require("mongoose");

const Product = require("../models");

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
    const products = await Products.find().sort({ createdAt: -1 });
    res.status(200).json(products);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};
