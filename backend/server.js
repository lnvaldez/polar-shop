require("dotenv").config();

//* Requirements
const express = require("express");
const cors = require("cors");
const path = require("path");

/* Imports */
//* Config
const connectDB = require("./config/db.config");
//* Routes
const routes = require("./routes");

const PORT = process.env.EXPRESS_PORT || 8080;

const app = express();

//* App set
app.set("view engine", "pug");
app.set("views", path.join(__dirname, "views"));

//* App use
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(
  cors({
    origin: "http://localhost:5001",
    methods: ["POST"],
  })
);

app.use(express.static(path.join(__dirname, "public")));

app.use((req, res, next) => {
  console.log(req.path, req.method);
  next();
});

app.use("/", routes);

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
