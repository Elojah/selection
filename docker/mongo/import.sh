#! /bin/bash

mongorestore -d main -c users dump/main/users.bson
mongorestore -d main -c tasks dump/main/tasks.bson
mongorestore -d tags -c tasks dump/tags/taskTags.bson
