### ARCHIVED: I no longer have the time or motivation to continue this project, it was pretty ambicious from the start, but I might resume sometime

# Noteer

ok hi this is my new project...

oh you wanna know what it is? well, its just a notes app with some cool features

what else do i say?

ummm... i guess you can look into the code and help me but its ~~probably better if u dont~~actually do, i need help

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

If you want to contribute to the project, having a dev instance is essential, as it's the only way we can get reproducible builds. 

```bash
# Clone the repository
git clone https://github.com/pisanvs/noteer
cd noteer

# Configure environment variables

cp .default.env .env
vim .env # edit to your liking

# and start the containers
docker-compose --file dev.docker-compose.yml up -d

# you should also add the nginx container ip to your /etc/hosts file for accesing locally, here's how i do it

# first make a copy of your hosts file to the hosts.ignore file in the noteer root dir
sudo cp /etc/hosts ./hosts.ignore

# add noteer entry to hosts file (make sure this entry matches the domain name in .env)
sudo sh -c 'echo "*\tnoteer.local" >> ./hosts.ignore'

# finally, run this command to replace the placeholder in the hosts.ignore file with the nginx container's ip 

sudo sed s/*/$(sudo docker inspect nginx | jq -M ".[].NetworkSettings.Networks.noteer_default.IPAddress" | tr -d "\"")/g hosts.ignore | sudo tee /etc/hosts
```

## License

ALL CODE IN THIS PROJECT, UNLESS **EXPLICITLY**[^1] STATED, IS LICENSED UNDER THE GNU GENERAL PUBLIC LICENSE VERSION 3.

[^1]: (Definition 1a) https://www.merriam-webster.com/dictionary/explicit 
