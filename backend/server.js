require("dotenv").config();

//* Requirements
const express = require("express");

/* Imports */
//* Config
const connectDB = require("./config/db.config");
//* Routes
const productRoutes = require("./routes");

const PORT = process.env.EXPRESS_PORT || 8080;

const app = express();

app.use(express.json());

app.use((req, res, next) => {
  console.log(req.path, req.method);
  next();
});

app.use("/products", productRoutes);

const startServer = async () => {
  try {
    await connectDB();
    app.listen(PORT, () => {
      console.log(`⭐ Express server running on port ${PORT}`);
    });
  } catch (error) {
    console.error("❌ Failed to start server:", error);
    process.exit(1);
  }
};

startServer();
