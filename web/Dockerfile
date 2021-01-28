FROM node:14

WORKDIR /app

ENV PORT 3000
ENV HOST 0.0.0.0

COPY package.json ./
RUN yarn install

COPY . .

EXPOSE 3000

RUN yarn build

CMD ["yarn", "start"]
