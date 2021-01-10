# build stage
FROM node:lts-alpine
WORKDIR /var/www/switter-front/
COPY ./package.json .
COPY ./package-lock.json .
COPY ./public/ .
COPY ./src/assets ./public
RUN mkdir ./src
COPY ./babel.config.js .
COPY ./vue.config.js .
RUN npm cache verify
RUN npm install

EXPOSE 8080
CMD ["npm", "run", "serve"]