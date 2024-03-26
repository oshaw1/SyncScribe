# Use the official Node.js image as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package.json ./
COPY package-lock.json ./

# Install the frontend dependencies
RUN npm install

# Install react-scripts globally
RUN npm install -g react-scripts

# Copy the frontend source code to the container
COPY . .

# Build the React application
RUN npm run build

# Use a lightweight web server to serve the built files
FROM nginx:alpine
COPY --from=0 /app/build /usr/share/nginx/html

# Expose the port on which the frontend will be served
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]