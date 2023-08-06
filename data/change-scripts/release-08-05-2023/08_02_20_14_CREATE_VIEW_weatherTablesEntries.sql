CREATE VIEW weatherTableEntriesAndDetails AS SELECT
wtbl.weatherSystemId,
wtbl.name,       
wtbl.diceType,   
wtbl.diceNumber, 
wt.weatherName,   
wte.diceResult,
wd.location, 
wd.details
from weatherTableEntries wte   
join weatherTypes wt on wte.weatherTypeId  = wt.id 
join weatherTables wtbl on wte.tableId  = wtbl.id
join weatherDetails wd  on wd.weatherTypeId  = wte.weatherTypeId