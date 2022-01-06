FROM node:current-alpine

ENV NODE_ENV production
WORKDIR /usr/src/app

COPY ["package.json", "yarn.lock", "rollup.config.js", "tailwind.config.js", "./"]

COPY .env.production .env
COPY ./src ./src
COPY ./public ./public

RUN yarn

EXPOSE 5000
CMD yarn start --host