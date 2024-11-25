const jwt = require("jsonwebtoken");

const auth = async (req, res, next) => {
  const token = req.session.token;

  if (!token) {
    return res.redirect("http://localhost:5001/login");
  }

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    req.user = { id: decoded._id, role: decoded.role };
    next();
  } catch (error) {
    console.error("JWT verification failed: ", error);
    req.session.destroy();
    return res.redirect("http://localhost:5001/login");
  }
};

module.exports = { auth };
