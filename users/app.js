const express = require('express');
const mongoose = require('mongoose');
const { connect } = require('node-nats-streaming');
const {execSync} = require('child_process');

execSync('sleep 5');

const app = express();
const port = 3000;
// Connect to MongoDB

const natsConnection = connect("test-cluster", "users-service", { url: 'nats://nats-streaming:4222' });

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

  const data = { userId: user._id, name, email };
  const jsonData = JSON.stringify(data);
  natsConnection.publish('user.created', jsonData, (err, guid) => {
    if (err) {
      console.error(`Error publishing to NATS Streaming: ${err}`);
    } else {
      console.log(`Published message with guid: ${guid}`);
    }
  });
  
  res.json(user);
});

// Add other CRUD routes as needed

// Start server
app.listen(port, async () => {
  console.log(`Users service listening at http://localhost:${port}`);
});