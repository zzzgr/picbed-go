cd ../front
yarn build
cd ../static
mv ../front/dist/* ./assets

# 在static目录执行
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...