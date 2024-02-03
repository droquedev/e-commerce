const express = require('express');
const mongoose = require('mongoose');

const app = express();
const port = 3000;

// Connect to MongoDB
mongoose.connect(process.env.MONGO_CONNECTION_STRING, { useNewUrlParser: true, useUnifiedTopology: true });
const db = mongoose.connection;
db.on('error', console.error.bind(console, 'MongoDB connection error:'));

// User model
const User = mongoose.model('User', { name: String, email: String });

// Express middleware
app.use(express.json());

// Routes
app.get('/users', async (req, res) => {
  const users = await User.find();
  res.json(users);
});

app.post('/users', async (req, res) => {
  const { name, email } = req.body;
  const user = new User({ name, email });
  await user.save();
  res.json(user);
});

// Add other CRUD routes as needed

// Start server
app.listen(port, () => {
  console.log(`Users service listening at http://localhost:${port}`);
});