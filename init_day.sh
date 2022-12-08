

dir_name="day_$(printf "%02d" $1)"
cp -R template $dir_name
cd $dir_name
wget --load-cookies=../cookies.txt "https://adventofcode.com/2022/day/$1/input"
touch main.go