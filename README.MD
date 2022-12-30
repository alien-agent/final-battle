# Лабораторная работа №5 - PDA
## Участники
- Пашкевич Денис [@alien-agent](github.com/alien-agent)
- Сайкин Антон [@kekega](github.com/kekega)
- Стафиефский Стас [@stasstaf](github.com/stasstaf)

## Установка
Требуется Go 1.19, но скорее всего будет работать и на более ранних версиях.
```shell
git clone github.com/alien-agent/final-battle
cd final-battle
go build
```

## Параметры запуска
Программа по-умолчанию отрисовывает PDA в формате DOT, и сохраняет 
результат в файл **pda.dot** в одной папке с программой.

Поддерживаются следующие параметры командной строки:
<pre>
-m, --mode=[stdin|file] режим работы программы (чтение из входного потока или файла)
    --pda-filename      путь до файла с описанием PDA
    --language-filename путь до файла с параметрами грамматики (опционально)
    --words-filename    путь до файла со словами для обработки (обязательно для mode=="file")
</pre>
В случае использования флага `--language-filename`, в каждой строке файла настройки
указываются имя и значение параметра грамматики через пробел. См. пример.

## Примеры
### Описание PDA
<pre>
States -> {q0, q1}
FinalStates -> {q0}
InputAlphabet -> {(, )}
StackAlphabet -> Z0, B
InitialState -> q0
InitialStackSymbol -> Z0
Transition -> q0 / q1 / ( / Z0 / B, Z0
Transition -> q1 / q1 / ( / B / B, B
Transition -> q1 / q1 / ) / B /
Transition -> q1 / q0 / ε / Z0 / Z0
</pre>
### Параметризация грамматики (default value)
```
Arrow ->
Delimiter ,
TransitionDelimiter /
UniversalQuantifier ∀
Epsilon ε
```

### Условные обозначения
Условные обозначения для графа, полученного в результате работы программы
![Alt text](Images/legend.png)![Alt text](Images/states_legend.png)
