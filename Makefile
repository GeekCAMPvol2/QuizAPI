mongoRun: docker run -it --name quiz-mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=mongo mongo
mongoStart: docker start mongo
test: go test -v -cover ./...
.PHONY: mongoRun mongoStart test