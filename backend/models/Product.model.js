const mongoose = require("mongoose");

const productSchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
    },
    category: {
      type: String,
      required: true,
    },
    price: {
      type: Number,
      min: 0,
      required: true,
    },
    quantity: {
      type: Number,
      min: 0,
      default: 0,
    },
    image: {
      type: String,
      required: false,
    },
    isAvailable: {
      type: Boolean,
      default: false,
    },
    description: String,
  },
  { timestamps: true }
);

module.exports = mongoose.model("Product", productSchema);
