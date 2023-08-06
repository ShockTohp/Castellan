CREATE VIEW detailedWeatherHexReports AS SELECT 
wm.weatherDate as weatherDate, 
wm.hexId as hexId, 
wm.campaignId as campaignId,
wt.weatherName as weatherName, 
wd.location as location, 
wd.details as details
from weatherMarkers wm 
join weatherHexes wh on wh.Id = wm.hexId
join weatherTypes wt on wt.Id = wh.weatherTypeId
join weatherDetails wd on wd.weatherTypeId = wh.weatherTypeId ;