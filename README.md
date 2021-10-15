# Noteer

ok hi this is my new project...

oh you wanna know what it is? well, its just a notes app with some cool features

what else do i say?

ummm... i guess you can look into the code and help me but its probably better if u dont

## How to set up a local instance

If you want to run inside of a docker container (basically the only way unless you want to be infront of a terminal feeling helpless for the next 15 hours of your life):

```bash
# Clone the repository
git clone https://github.com/pisanvs/noteer
cd noteer

# Configure environment variables

cp .default.env .env
vim .env # edit to your liking

# start the containers
docker-compose up -d
```

## How to set up a dev instance

```bash
# Clone the repository
git clone https://github.com/pisanvs/noteer
cd noteer

# Configure environment variables

cp .default.env .env
vim .env # edit to your liking

# and start the containers
docker-compose --file dev.docker-compose.yml up -d
```

