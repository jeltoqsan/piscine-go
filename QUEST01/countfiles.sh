find $PWD \( -type f -or -type d \) -name ".*" -prune -o -print | wc -l # type f,d - поиск файлов и директорий
# способ 2 find . | wc -l