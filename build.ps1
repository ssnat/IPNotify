$version = $args[0]

$dir = "./build/v$version"

$linuxArm64Dir = $dir + "/IPNotify-linux-arm64-v" + $version
$linuxAmd64Dir = $dir + "/IPNotify-linux-amd64-v" + $version
$windowsArm64Dir = $dir + "/IPNotify-windows-arm64-v" + $version
$windowsAmd64Dir = $dir + "/IPNotify-windows-amd64-v" + $version


$linuxArm64Path = $linuxArm64Dir + "/IPNotify"
$linuxAmd64Path = $linuxAmd64Dir + "/IPNotify"
$windowsArm64Path = $windowsArm64Dir + "/IPNotify.exe"
$windowsAmd64Path = $windowsAmd64Dir + "/IPNotify.exe"

echo $linuxArm64Path
$env:GOOS="linux"
$env:GOARCH="arm64"
go build -o $linuxArm64Path .

echo $linuxAmd64Path
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o $linuxAmd64Path .

echo $windowsArm64Path
$env:GOOS="windows"
$env:GOARCH="arm64"
go build -o $windowsArm64Path .

echo $windowsAmd64Path
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o $windowsAmd64Path .

cp ./config.yaml $linuxArm64Dir
cp ./config.yaml $linuxAmd64Dir
cp ./config.yaml $windowsAmd64Dir
cp ./config.yaml $windowsArm64Dir

echo "Done"