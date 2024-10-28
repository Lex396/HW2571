1.	git init - инициализируем Git в папке с проектом
2.	git remote add origin https://github.com/... - подулючаемся к удаленному репозиторию
3.	git branch -M master - создем основную ветку 
4.	git add . - добавляем все файлы для отслеживания
5.	git commit -m "start ver 1.0"" - фиксируем состояние
6.	git tag -a V1.0 -m "version 1.0" - ставим тэг
7.	git push -u origin master - передаем состояние в репозиторий
8.	git push origin V1.0 - передаем тэг в репозиторий
9.	git checkout -b develop master- создаем ветку разработки на основе основной ветки
10.	git push origin develop - передаем состояние в репозиторий
11.	git checkout -b feature/chg_input - создаем тематическу ветку для изменения типа вводимых данных на основе ветки develop
12.	меняем код программы
13.	git add main.go - добавляем файл для отслеживания
14.	git commit -m "changing the type of input data" - фиксируем состояние
15.	git push origin feature/chg_input - передаем состояние в репозиторий
16.	git checkout develop - переключаемся на ветку разработки
17.	git merge --no-ff feature/chg_input - мержим тематическую ветку в develop
18.	git push origin develop - передаем состояние в репозиторий
19.	git checkout -b release develop - создаем ветку релиза на основе ветки develop
20.	git push origin release - передаем состояние в репозиторий
21.	git tag -a V1.1 -m "version V1.1" - доавляем тэг
22.	git push origin V1.1 - передаем тэг в репозиторий
23.	git checkout master - переключаемся на ветку мастер
24.	git merge --no-ff release - мержим ветку релиза с основной
25.	git commit -m "ver V1.1" - фиксируем состояние
26.	git push origin master - передаем состояние в репозиторий
27.	git checkout -b hotfix master - создаем ветку исправлений на основе ветки master
28.	редактируем текст в коде
29.	git add main.go - добавляем файл для отслеживания
30.	git commit -m "text edits" - фиксируем состояние
31.	git push origin hotfix - передаем состояние в репозиторий
32.	git checkout master - переключаемся на ветку мастер
33.	git merge hotfix - мержим ветку hotfix с master
34.	git push origin master - передаем состояние в репозиторий
35.	git checkout develop - пеерключаемся на ветку develop
36.	git merge hotfix - мержим ветку hotfix с develop
37.	git commit -m "hotfix" - фиксируем состояние
38.	git push origin develop - передаем состояние в репозиторий
