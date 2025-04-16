FROM node:lts-alpine

RUN npm install -g http-server
WORKDIR /app
COPY Frontend_BD2/package*.json ./
RUN npm install
COPY Frontend_BD2 .
RUN npm run build

EXPOSE 8081
CMD [ "http-server", "dist", "-p=8081" ]
