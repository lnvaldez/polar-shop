require("dotenv").config();
const express = require("express");
const connectDB = require("./config/db.config");

const PORT = process.env.EXPRESS_PORT || 8080;

const app = express();

const startServer = async () => {
  try {
    await connectDB();
    app.listen(PORT, () => {
      console.log(`🚀 Express server running on port ${PORT}`);
    });
  } catch (error) {
    console.error("❌ Failed to start server:", error);
  }
};

startServer();
