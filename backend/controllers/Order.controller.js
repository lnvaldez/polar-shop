const mongoose = require("mongoose");

const { Product } = require("../models");
const { Order } = require("../models");

const approveOrder = async (req, res) => {
  const { id } = req.params;
  const { productName, productQuantity } = req.body;

  if (!mongoose.Types.ObjectId.isValid(id)) {
    return res.status(404).json({ error: "Order not found" });
  }

  try {
    const product = await Product.findOneAndUpdate(
      { name: productName },
      { $inc: { quantity: -productQuantity } },
      { new: true }
    );

    if (!product) {
      return res.status(404).json({ error: "Product not found" });
    }

    const order = await Order.findByIdAndUpdate(
      id,
      { status: "Approved" },
      { new: true }
    );

    if (!order) {
      return res.status(404).json({ error: "Order not found" });
    }

    return res
      .status(200)
      .json({ message: "Order approved successfully", order, product });
  } catch (error) {
    console.error(error);
    return res.status(500).json({ error: "Internal server error" });
  }
};

module.exports = { approveOrder };
