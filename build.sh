rm -rf target
git clone https://github.com/Blizzard/s2client-proto target
mkdir target/build_res
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=target/build_res -I target/ target/s2clientprotocol/sc2api.proto
rm -rf SC2APIProtocol/*
cp target/build_res/s2clientprotocol/* SC2APIProtocol/
