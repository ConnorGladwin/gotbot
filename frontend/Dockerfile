FROM node:lts-alpine

WORKDIR /frontend

COPY package*.json

RUN npm i

COPY . .

RUN npm run build

EXPOSE 5173
CMD["dist"]
