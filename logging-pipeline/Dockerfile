
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy package files & install dependencies
COPY package*.json ./
RUN npm install

# Copy source files
COPY . .

# Copy wait-for-logstash script and make it executable
COPY wait-for-logstash.sh /wait-for-logstash.sh
RUN chmod +x /wait-for-logstash.sh

# Start using the wait script
CMD ["/wait-for-logstash.sh"]
