
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy package files & install dependencies
COPY package*.json ./
RUN npm install

# Copy source files
COPY . .

# Start the server
CMD ["node", "src/index.js"]
