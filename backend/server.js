const express = require("express");

const PORT = process.env.EXPRESS_PORT || 8080;

const app = express();

app.listen(8080, () => {
  console.log(`Server running on port ${PORT}`);
});
