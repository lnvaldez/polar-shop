const mongoose = require("mongoose");

const { Product } = require("../models");
const { User } = require("../models");

//* ..Render views
const renderAdminDashboard = async (req, res) => {
  try {
    const products = await Product.find();
    const totalProducts = products.length;
    res.render("admin/dashboard", { totalProducts });
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderAllProductsPage = async (req, res) => {
  try {
    const products = await Product.find().sort({
      isAvailable: -1,
      createdAt: -1,
    });
    res.render("admin/productList", { products });
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderProductPage = async (req, res) => {
  const id = req.params.id;

  try {
    const product = await Product.findById({ _id: id });
    res.render("admin/productDetail", { product });
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderAddProductPage = async (req, res) => {
  try {
    res.render("admin/addProduct");
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderUsersPage = async (req, res) => {
  try {
    const users = await User.find().sort({ createdAt: -1 });
    res.render("admin/users", { users });
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderOrdersPage = async (req, res) => {
  try {
    res.render("admin/orders");
  } catch (error) {
    console.error({ error: error.message });
    res.render("error");
  }
};

const renderSettingsPage = async (req, res) => {
  try {
    res.render("admin/settings");
  } catch (error) {
    console.error({ error: erro.message });
    res.render("error");
  }
};

module.exports = {
  renderAdminDashboard,
  renderAllProductsPage,
  renderProductPage,
  renderAddProductPage,
  renderUsersPage,
  renderOrdersPage,
  renderSettingsPage,
};
