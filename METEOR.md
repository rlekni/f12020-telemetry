Bundle on Intel machine (https://github.com/meteor/meteor-feature-requests/issues/130)
# Install git
sudo apt-get install git

# Install meteor
curl https://install.meteor.com/ | sh

# Go to app
cd <your-meteor-app-here>

# Bundle your meteor app
meteor build <Destination path>

# Run mongo
# Install MongoDB
curl -s https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb [ arch=arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
sudo apt update
sudo apt upgrade
sudo apt install mongodb-org

# Start MongoDB
sudo systemctl enable mongod
sudo systemctl start mongod

# Check MongoDB
sudo systemctl status mongod

# Install NodeJS
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo apt-get install gcc g++ make

Go to meteor bundle
Install bundle dependencies
cd <bundle>/programs/server/
npm install

export MONGO_URL='mongodb://localhost:27017/app-name'
export ROOT_URL='http://ubuntu.local:3000/'
export PORT=3000
Run bundled app
cd <bundle>
node main

sample docker meteor https://github.com/disney/meteor-base
https://github.com/tozd/docker-meteor
https://guide.meteor.com/deployment.html
https://dockerize.io/guides/docker-meteor-guide
https://github.com/zodern/meteor-docker
https://www.meteor.com/tutorials/vue/creating-an-app
https://guide.meteor.com/vue.html