#Build app
FROM node:alpine3.14 as build-step

RUN mkdir -p /app

WORKDIR /app

COPY ./front/SSOPractice/package.json /app

RUN npm install

COPY ./front/SSOPractice /app
COPY ./configs/environment.prod.ts /app/src/environments/environment.prod.ts

ARG SCRIPT

RUN npm run $SCRIPT

#Deploy app
FROM nginx:alpine

COPY --from=build-step /app/dist/ssopractice /usr/share/nginx/html