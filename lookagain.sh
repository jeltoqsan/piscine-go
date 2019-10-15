find -name "*.sh" | cut -c 3-30 | sed 's/...$//'

# 1. поиск и отображение файлов с расширением .sh
# 2. cut -c 3-30 - извлекает символы с 3 по 30 (то есть пропускает 1 и 2 и 31 и дальше)
# 3. sed 's/...$//' - замена последних трех символов на пустой символ (пробел)
 
# 2 способ: find -name "*.sh" -printf "%f\n" | sed 's/...$//'
# 3 способ: find -name "*.sh" -print | sed 's/...$//' | sed 's/\(.\)\///' - можно без print
