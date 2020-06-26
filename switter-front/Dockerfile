# build stage
FROM node:lts-alpine as build-stage
WORKDIR /var/www/switter-front/
COPY ./src/switter-front/package.json .

#COPY ./src/switter-front/package.json .
#WORKDIR /app/switter-front/
#RUN npm install -g @vue/cli @vue/cli-service 
RUN npm install --save vue-axios axios
RUN npm install
COPY ./src/switter-front/ /var/www/switter-front/
RUN npm run build

# production-stage
FROM nginx:stable-alpine as production-stage
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY --from=build-stage /var/www/switter-front/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]