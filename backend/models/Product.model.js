const mongoose = require("mongoose");

const productSchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
    },
    price: {
      type: Number,
      required: true,
    },
    stock: {
      type: Number,
      default: 0,
    },
    image: {
      type: String,
      required: false,
    },
    available: {
      type: Boolean,
      default: true,
    },
    description: String,
  },
  { timestamps: true }
);

productSchema.pre("save", function (next) {
  if (this.stock == 0) {
    this.available = false;
  }
  next();
});

module.exports = mongoose.model("Product", productSchema);
