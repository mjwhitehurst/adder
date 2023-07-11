# Stage 1: Build the ReactJS application
FROM node:14-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the frontend source code to the working directory
COPY . .

# Build the React application
RUN npm run build

# Stage 2: Create a lightweight container to serve the built React application
FROM nginx:alpine

# Copy the built files from the builder stage to the NGINX web server
COPY --from=builder /app/build /usr/share/nginx/html

# Expose the desired port
EXPOSE 8000

# Start the NGINX web server
CMD ["nginx", "-g", "daemon off;"]