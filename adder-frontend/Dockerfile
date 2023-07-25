# Specify the base image
FROM node:14

# Set the working directory in the container to /app
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies in the container
RUN npm install

# Copy the rest of your app's source code from your host to your image filesystem.
COPY . .

# Specify the port
EXPOSE 3000

# The command to run your application when the docker container is run
CMD [ "node", "app.js" ]