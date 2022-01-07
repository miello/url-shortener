FROM node:current-alpine

WORKDIR /usr/src/app

COPY ["package.json", "yarn.lock", "rollup.config.js", "tailwind.config.js", "./"]

COPY .env.production .env
COPY ./src ./src
COPY ./public ./public

RUN rm -rf public/build

RUN yarn
RUN yarn build

RUN yarn autoclean --init
RUN yarn autoclean --force

EXPOSE 5000
CMD yarn start --host