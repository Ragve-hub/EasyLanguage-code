Run = true;

function main()
	-- Цикл, поддерживающий работу скрипта
	while Run do sleep(100); end;
end;


-- Вызывается терминалом QUIK в момент остановки скрипта
function OnStop()
	-- Выключает флаг, чтобы остановить цикл while внутри main
	Run = false;
	
end;

--Pipe

local function SendTeleMessage(msg, pipe_name)
   local tele_pipe = io.open("\\\\.\\PIPE\\some_pipe", "w+b") -- открываем именованный канал
   if not tele_pipe then
       return false
   end
   tele_pipe:write(msg) -- записываем сообщение в канал
   tele_pipe:close() -- закрываем канал
end

--- Функция вызывается терминалом QUIK при получении изменения стакана котировок
local my_class = "SPBFUT";
local my_sec = "RIU2";
Subscribe_Level_II_Quotes(my_class, my_sec);

function OnQuote(class, sec)
   --
    if class == my_class and sec == my_sec then
         ql2 = getQuoteLevel2(class, sec);
      -- Представляет снимок СТАКАНА в виде СТРОКИ


QuoteStr = os.date("%Y%m%d")..";"
QuoteStr = QuoteStr..os.date("%H%M%S",os.time())..";"

         QuoteStr = QuoteStr..tonumber(getParamEx(class,  sec, "QTY").param_value)..";"  --Количество в последней сделке
         QuoteStr = QuoteStr..tonumber(getParamEx(class,  sec, "LAST").param_value)..";" --Цена последней сделки
 QuoteStr = QuoteStr..getParamEx(class,  sec, "OFFER").param_value..";"   --Лучшая цена предложения
 QuoteStr = QuoteStr..getParamEx(class,  sec, "OFFERDEPTH").param_value..";"   --Предложение по лучшей цене
  QuoteStr = QuoteStr..getParamEx(class,  sec, "BID").param_value..";"            --Лучшая цена спроса
 QuoteStr = QuoteStr..getParamEx(class,  sec, "BIDDEPTH").param_value..";"       --Спрос по лучшей цене



   -- Записывает строку 
        QuoteStr = QuoteStr.."\n"
        --message(QuoteStr)

SendTeleMessage(QuoteStr, 'some_pipe')
    end;
end;


