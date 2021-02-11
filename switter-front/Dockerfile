FROM node:lts-alpine
WORKDIR /var/www/switter-front/
COPY ./package.json .
COPY ./package-lock.json .
COPY ./babel.config.js .
COPY ./webpack.config.js .
RUN mkdir ./src
# RUN npm i webpack webpack-dev-server \
#     vue vue-loader vue-loader-plugin vue-template-compiler \
#     babel-loader @babel/core @babel/preset-env \
#     style-loader sass-loader css-loader \
#     node-sass \
#     html-webpack-plugin \
#     file-loader url-loader -D
# RUN npm cache verify
 RUN npm i

EXPOSE 8080
CMD ["npm", "run", "start"]