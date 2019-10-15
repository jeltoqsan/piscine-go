find . \( -type f , -type d \) | wc -l
# способ 2 find $PWD \( -type f -or -type d \) -name ".*" -prune -o -print | wc -l
# способ 3 ilifind . | wc -l