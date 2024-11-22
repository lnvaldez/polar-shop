/* Imports */
const mongoose = require("mongoose");
const validator = require("validator");
const bcrypt = require("bcrypt");

const userSchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
    },
    password: {
      type: String,
      required: true,
    },
    role: {
      type: String,
      default: "user",
    },
  },
  { timestamps: true }
);

//* Static
userSchema.statics.register = async function (name, password) {
  if (!name || !password) {
    throw Error("Fill out all required fields");
  }

  if (!validator.isStrongPassword(password)) {
    throw Error("Password is not strong enough");
  }

  const exists = await this.findOne({ name });

  if (exists) {
    throw Error("Name already in use");
  }

  const salt = await bcrypt.genSalt(10);
  const hash = await bcrypt.hash(password, salt);

  const user = await this.create({ name, password: hash });

  return user;
};

userSchema.statics.login = async function (name, password) {
  if (!name || !password) {
    throw Error("Fill out all required fields");
  }

  const user = await this.findOne({ name });

  if (!user) {
    throw Error(`User '${name}' does not exist`);
  }

  const match = await bcrypt.compare(password, user.password);

  if (!match) {
    throw Error("Incorrect password");
  }

  return user;
};

module.exports = mongoose.model("User", userSchema);
