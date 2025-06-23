# Use an official Node.js runtime as a parent image
FROM node:18-alpine3.21

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application's source code
COPY . .

# Vite uses port 5173 by default, but we'll map it to 3000
EXPOSE 3000

# The command to run the dev server
# The "-- --host" is crucial to expose the server outside the container
CMD ["npm", "run", "dev", "--", "--port", "3000", "--host"]
