# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY ./src/switter-front/package*.json ./
RUN npm install
COPY ./src/switter-front/ .
RUN npm run build

# production-stage
FROM nginx:stable-alpine as production-stage
COPY ./etc/nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]