package data

import (
	"fmt"
	"errors"
	"strconv"
)
const campaignTable = "campaigns"
type Campaign struct {
	id int
	guildId int64
	name string
	weatherSystem *WeatherSystem
}

func NewCampaign(id int, guildId int64, name string, ws *WeatherSystem) *Campaign {
	return &Campaign {
		id: id,
		guildId: guildId,
		name: name,
		weatherSystem: ws,
	}
}

func (c Campaign) Id() int {
	return c.id
}

func (c Campaign) GuildId() int64 {
	return c.guildId
}

func (c Campaign) Name() string {
	return c.name
}

func (c Campaign) WeatherSystem() *WeatherSystem {
	return c.weatherSystem
}

func (c Campaign) CampaignHasReportForDate(date string) bool {
	det := GetWeatherDetailsForCampaignDate(c.id, date)	
	if len(det) > 0 {
		return true
	}
	return false
}

func GetCampaignByGuild(gi string) ( *Campaign, error) {
	i, _ := strconv.ParseInt(gi, 10, 64);
	
	tableq :=  fmt.Sprintf("SELECT * FROM %s c WHERE c.guildId = ?", campaignTable);
	rows, err := runQuery(tableq, i)
	defer rows.Close()
	checkerr(err)

	campaigns := make([]Campaign, 0)
	weatherSystemIds := make([]int, 0)
	for rows.Next() {
		currentC := Campaign{}
		var currentWsId = 0 
		err = rows.Scan(&currentC.id, &currentC.guildId, &currentC.name, &currentWsId) 
		checkerr(err)
		campaigns = append(campaigns, currentC);
		weatherSystemIds = append(weatherSystemIds, currentWsId)
	}

	if (len(campaigns) > 1) {
		return nil, errors.New(fmt.Sprintf("This server has more than one campaign configured."));
	} else if (len(campaigns) < 1) {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve any campaings. Does this server have one registered?"));
	}
	ws, err := getWeatherSystemById(weatherSystemIds[0]);
	if err != nil {
		return nil, err
	}
	 return NewCampaign(campaigns[0].id, campaigns[0].guildId, campaigns[0].name, ws), nil;	


}

func GetCampaignById(id int) ( *Campaign, error) {
	tableq :=  fmt.Sprintf("SELECT * FROM %s c WHERE c.guildId = ?", campaignTable);
	rows, err := runQuery(tableq, id)
	defer rows.Close()
	checkerr(err)

	campaigns := make([]Campaign, 0)
	weatherSystemIds := make([]int, 0)
	for rows.Next() {
		currentC := Campaign{}
		var currentWsId = 0 
		err = rows.Scan(&currentC.id, &currentC.guildId, &currentC.name, &currentWsId) 
		checkerr(err)
		campaigns = append(campaigns, currentC);
		weatherSystemIds = append(weatherSystemIds, currentWsId)
	}

	if (len(campaigns) > 1) {
		return nil, errors.New(fmt.Sprintf("This server has more than one campaign configured."));
	} else if (len(campaigns) < 1) {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve any campaings. Does this server have one registered?"));
	}
	ws, err := getWeatherSystemById(weatherSystemIds[0]);
	if err != nil {
		return nil, err
	}
	 return NewCampaign(campaigns[0].id, campaigns[0].guildId, campaigns[0].name, ws), nil;	


}

func RegisterCampaign(GuildID, name, wt string) (string, error) {
	gi, _ := strconv.ParseInt(GuildID, 10, 64)
	if guildHasAtLeastOneCampaign(gi) {
		return "I am sorry, I can only support one campaign per server at this time.", nil
	}
	ws, wse := getWeatherSystemId(wt)
	if wse != nil {
		return "Error retrieving weather system information", wse
	}
	registerQ := fmt.Sprintf("INSERT INTO %s (guildId, name, weatherSystemId) VALUES (?, ?, ?);", campaignTable)
	stmt, err := db.Prepare(registerQ)
	checkerr(err)
	_, err = stmt.Exec(gi, name, ws);
	checkerr(err)
	defer stmt.Close()
	return "", nil
}

func guildHasAtLeastOneCampaign(gi int64) bool {
	var count int
	cq := fmt.Sprintf("SELECT COUNT(*) FROM %s c WHERE c.guildID = %d;", campaignTable, gi)
	err := db.QueryRow(cq).Scan(&count)
	checkerr(err)
	if count > 0 {
		return true;
	} 
	return false
}




