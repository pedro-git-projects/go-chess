# Use the official Node.js image as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install 

# Copy the rest of the application source code
COPY . .

# Build the frontend application
RUN npm run build

# Expose the port on which the application will run
EXPOSE 5173

# Set the command to run the application
CMD ["npm", "start"]

