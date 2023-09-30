FROM node:latest
WORKDIR /app

# copy dependencies
COPY package.json package-lock.json ./
RUN npm install
COPY . .

# build the client
RUN npm run build

# expose the port and start the client
EXPOSE 3000
CMD ["npm", "start"]
