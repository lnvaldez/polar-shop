const { User } = require("../models");
const jwt = require("jsonwebtoken");

const generateToken = (_id, role) => {
  return jwt.sign({ _id, role }, process.env.JWT_SECRET, { expiresIn: "3d" });
};

const register = async (req, res) => {
  const { name, email, password } = req.body;

  try {
    const user = await User.register(name, email, password);

    res.status(200).json({ name });
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

const login = async (req, res) => {
  const { email, password } = req.body;

  try {
    const user = await User.login(email, password);

    const token = generateToken(user._id, user.role);

    req.session.token = token;

    res.redirect("/admin/dashboard");
  } catch (error) {
    res.status(400).json({ error: error.message });
  }
};

module.exports = { register, login };
