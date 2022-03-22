
* <p>[<a href="http://www.jurikres.com/catalog1/ms_cfb.htm#top">1</a>]      Composite Fractal Behavior, Jurik Research </p>


График РТС Дневной таймфрейм


Логика следующая:
В качестве примера возьмем уровень 10. Рост индикатора тренд, падение индикатора флет.


<img src="https://raw.githubusercontent.com/Ragve-hub/scribble/gh-pages/images/cfb_chart.jpg" alt="">


## Правила

### Вход Лонг
Если JR_Fractal24 (h+l,5) меньше 10 то ставим ордер на покупку на highest(high,5)

### Выход
Выход EOD в конце вечерней сессии или по трейлингу lowest (low,5)

## Бэктестинг
Тест на 1 контракт. Фьючерс РТС, Таймфрейм 1 час.
Интрадей . 2005- 2020
Recovery 14, win 48%, Trades 2626, Avg 180, MaxDD -13,2%

<img src="https://raw.githubusercontent.com/Ragve-hub/scribble/gh-pages/images/equity_cfb.jpg" alt="">
