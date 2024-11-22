const User = require("../models");
const jwt = require("jsonwebtoken");

const generateToken = (_id) => {
  return jwt.sign({ _id }, process.env.JWT_SECRET, { expiresIn: "3d" });
};

const login = async (req, res) => {
  const { name, password } = req.body;

  try {
    const user = await User.login(name, password);

    const token = generateToken(user._id);

    res.status(200).json({ name, token });
  } catch (error) {
    res.statuso(400).json({ error: error.message });
  }
};
