# This is my own little Golang app for the purpose of developing b-tree searches and a W3 service for fun. It is not intended for anyone else to use for any reason. However, I am experimenting with performance improvements with pointers etc. So, do feel free to explore the code for that reason.

## Docker build
```
$ docker image build -t hraywilson/golang-vote-search-rest:1.2.1 .
moved to scratch image
$ docker image build -t hraywilson/scratch-vote-search-rest:1.0.0 .
including JSON and export
$ docker image build -t hraywilson/scratch-vote-search-rest-json:1.0.0 .
pointer based code base CRAZY FAST
$ docker image build -t hraywilson/scratch-vote-base-rest-pointers:3.0.0 .
```

## Docker container start
```
$ docker container run --name vote-search-rest-api --rm -d \
-v /lvm_one/landing_zone/data/vote_data/colorado/docker-datapath/:/vote_data \
-p 8080:8080 \
hraywilson/scratch-vote-base-rest-pointers:3.0.0
```

## Load up a bunch O data
```
$ FN=`ls -H  ../vote_data/arap* | awk -F"/" '{print $NF}'`; for f in $FN; do curl -s horus.byrds:8080/load_base_data/$f | jq; done; curl -s horus.byrds:8080/sort_base_data | jq
```