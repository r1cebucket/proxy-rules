go run main.go

# compile singbox rule-set
mkdir -p ./data/sing-box
cd ./data/sing-box

file_name=$(basename $1)
file_name_base=$(basename $file_name '.tar.gz')
wget $1
tar -xzf $file_name
rm $file_name
mv ./$file_name_base/sing-box ./
rm -r $file_name_base

./sing-box rule-set compile ../rules/sing-box_direct.json
./sing-box rule-set compile ../rules/sing-box_proxy.json
./sing-box rule-set compile ../rules/sing-box_reject.json