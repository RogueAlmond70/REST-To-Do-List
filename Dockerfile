FROM mongo
# Create a user with the "readWriteAnyDatabase" role
CMD mongo --eval "db.getSiblingDB('admin').createUser({user: 'myuser', pwd: 'mypass', roles: [{role: 'readWriteAnyDatabase', db: 'admin'}]});"
#$ docker build -t my-mongo .
#$ docker run -d -p 27017:27017 --name my-mongo my-mongo
