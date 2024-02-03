# Use the official Node.js image
FROM node:20-alpine3.19

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files to the container
COPY . .

# Expose the port that the application will run on
EXPOSE 3000

# Command to run the application
CMD ["node", "app.js"]