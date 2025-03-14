FROM node:18

WORKDIR /app

COPY package.json ./

RUN npm install --include=dev

COPY . .

EXPOSE $CLIENT_PORT
CMD sh -c "npm run dev -- --host 0.0.0.0 --port $CLIENT_PORT"
