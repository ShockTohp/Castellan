PRAGMA foreign_keys = ON;
CREATE TABLE campaigns (
	id INTEGER,
	guildId INTEGER,
	name TEXT,
	weatherSystemId INTEGER not null REFERENCES weatherSystems (id),
	PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE resolutionTypes (
	id INTEGER,
	name TEXT,
	PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE weatherSystems (
	id INTEGER,
	systemName TEXT,
	resolutionTypeId INTEGER not null REFERENCES resolutionTypes (id),
	startingHex INTEGER,
	PRIMARY KEY (id AUTOINCREMENT)
	);

CREATE TABLE weatherTypes (
	id INTEGER,
	weatherSystemId Integer,
	weatherName TEXT,
	FOREIGN KEY (weatherSystemId) REFERENCES weatherSystems (id),
	PRIMARY KEY (id AUTOINCREMENT)
	);

CREATE TABLE weatherHexes (
	id INTEGER,
	weatherSystemId INTEGER not null references weatherSystems(id),
	weatherTypeId INTEGER not null REFERENCES weatherTypes(id),
	one INTEGER,
	two INTEGER,
	three INTEGER,
	four INTEGER,
	five INTEGER,
	six INTEGER,
	PRIMARY KEY (id AUTOINCREMENT)
);


CREATE TABLE weatherMarkers (
	id INTEGER,
	campaignId INTEGER,
	weatherSystemId INTEGER,
    hexId INTEGER,
    weatherDate TEXT,
    lastDiceRoll INTEGER,
	PRIMARY KEY (id AUTOINCREMENT),
	FOREIGN KEY (weatherSystemId) REFERENCES weatherSystems (id),
	FOREIGN KEY (campaignId) REFERENCES campaigns (id)
);

CREATE TABLE weatherTables (
	id INTEGER,
	name TEXT,
	weatherSystemId INTEGER not null REFERENCES weatherSystems (id),
    diceType INTEGER,
    diceNumber TEXT,
	PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE weatherTableEntries (
	id INTEGER,
	diceResult INTEGER,
	tableId INTEGER not null REFERENCES weatherTables (id),
	weatherTypeId not null REFERENCES weatherTypes (id),
	PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE weatherDetails (
	id INTEGER,
	weatherTypeId INTEGER not null REFERENCES weatherTypes (id),
	location TEXT,
	details TEXT,
	PRIMARY KEY (id AUTOINCREMENT)
);


INSERT INTO resolutionTypes  (name)
VALUES ("Hex Flower");

INSERT INTO resolutionTypes  (name)
VALUES ("Dice");

INSERT INTO weatherSystems (systemName, resolutionTypeId, startingHex)
VALUES ("Yoon Suin", 1, 13);

INSERT INTO weatherSystems (systemName, resolutionTypeId)
VALUES ("OSE", 2);

INSERT INTO weatherTypes (id, weatherSystemId,  weatherName)
VALUES (1, 1, "Clear Skies");

INSERT INTO weatherTypes (id, weatherSystemId, weatherName)
VALUES (2, 1, "Heat Wave");

INSERT INTO weatherTypes (id, weatherSystemId, weatherName)
VALUES (3, 1, "Cold Front");

INSERT INTO weatherTypes  (id, weatherSystemId, weatherName)
VALUES (4, 1, "Strong Winds");

INSERT INTO weatherTypes (id, weatherSystemId, weatherName)
VALUES (5, 1, "Tropical Storm");

INSERT INTO weatherTypes (id, weatherSystemId, weatherName)
VALUES (6, 1, "Typhoon");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "No Wind");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Faint Breeze");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Gentle Breeze");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Moderate Breeze");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Fresh Breeze");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Strong Breeze");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "High Wind");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Near Gale");

INSERT INTO weatherTypes(weatherSystemId, weatherName)
VALUES (2, "Gale or Storm");

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (1, 1, 4, 34, 10, 4, 3, 2, 16);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (2, 1, 5, 33, 1, 3, 6, 5, 2);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (3, 1, 1, 1, 4, 8, 7, 6, 2);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (4, 1, 1, 35, 17, 8, 8, 3, 1);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (5, 1, 5, 5, 2, 6, 11, 10, 5);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (6, 1, 5, 2, 3, 7, 12, 11, 5);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (7, 1, 4, 3, 8, 14, 13, 19, 6);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (8, 1, 5, 4, 9, 15, 14, 7, 3);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (9, 1, 4, 36, 9, 16, 15, 8, 4);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (10, 1, 6, 10, 5, 11, 17, 10, 10);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (11, 1, 5, 5, 6, 12, 18, 17, 10);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (12, 1, 4, 6, 7, 13, 19, 18, 11);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (13, 1, 1, 7, 14, 21, 20, 19, 12);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (14, 1, 1, 8, 15, 22, 21, 13, 7);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (15, 1, 5, 9, 16, 23, 22, 14, 8);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (16, 1, 3, 16, 16, 16, 23, 15, 9);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (17, 1, 5, 10, 11, 18, 24, 4, 17);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (18, 1, 4, 11, 12, 19, 25, 24, 17);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (19, 1, 4, 12, 13, 20, 26, 25, 18);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (20, 1, 1, 13, 21, 28, 27, 26, 19);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (21, 1, 5, 14, 22, 29, 28, 20, 13);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (22, 1, 1, 15, 23, 30, 29, 21, 14);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (23, 1, 4, 1, 23, 2, 30, 22, 15);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (24, 1, 5, 17, 18, 25, 31, 24, 35);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (25, 1, 5, 18, 19, 26, 32, 31, 24);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (26, 1, 1, 19, 20, 27, 33, 32, 25);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (27, 1, 2, 20, 28, 35, 34, 33, 26);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (28, 1, 4, 21, 29, 36, 35, 27, 20);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (29, 1, 5, 22, 30, 37, 36, 28, 21);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (30, 1, 4, 23, 33, 5, 37, 29, 22);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (31, 1, 2, 24, 25, 32, 31, 31, 31);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (32, 1, 2, 25, 26, 33, 32, 32, 31);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (33, 1, 2, 26, 27, 34, 33, 33, 32);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (34, 1, 1, 27, 35, 31, 1, 34, 33);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (35, 1, 1, 28, 36, 24, 4, 34, 27);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (36, 1, 5, 29, 37, 17, 9, 35, 28);

INSERT INTO weatherHexes (id, weatherSystemId, weatherTypeId, one, two, three, four, five, six)
VALUES (37, 1, 1, 37, 34, 10, 16, 36, 29);

INSERT INTO weatherTables (id, name, weatherSystemId, diceType, diceNumber)
VALUES(1, "OSE Ocean Table", 2, 6, 2);


INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(2, 1, 7);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(3, 1, 8);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(4, 1, 9);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(5, 1, 10);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(6, 1, 11);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(7, 1, 11);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(8, 1, 11);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(9, 1, 12);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(10, 1, 13);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(11, 1, 14);

INSERT INTO weatherTableEntries (diceResult, tableId, weatherTypeId)
VALUES(12, 1, 15);

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (1, "The Haunted Jungles Of Lahag", "The air is humid and sticky.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (1, "The Topaz Isles", "The waves are calm, gulls soar freely in the cloudless skies above.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (1, "The Yellow City", "The city is pungent with the smells of filth, spices and incense.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (2, "The Haunted Jungles Of Lahag", "No torchers need in Silent Jungle Hexes during the day.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (2, "The Topaz Isles", "Doldrums oppress the silent skies. All sail powered boats come to a standstill.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (2, "The Yellow City", "The city takes a siesta. The markets are only open at night.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (3, "The Haunted Jungles Of Lahag", "Theres a cool and skin numbing fog in the air.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (3, "The Topaz Isles", "Clouds above dump their load into the sea, visibility is scant.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (3, "The Yellow City", "Cold winds from the upper Druk Yul bring snowflakes with them. An auspices event!");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (4, "The Haunted Jungles Of Lahag", "The trees sway noisely, checks relying on stealth are rolled with advantage.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (4, "The Topaz Isles", "The skies are excited, the blow unto the sea. All sail powered ships travel twice as fast.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (4, "The Yellow City", "The city smells fresh.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (5, "The Haunted Jungles Of Lahag", "The pitter patter of rain echos through the jungle and the ground becomes soft and muddy. Travel distance by foot or mount is halved.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (5, "The Topaz Isles", "A strong storm builds in the skies above, the sea is restless, gulls fly north hoping to avoid the potential Typhoon.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (5, "The Yellow City", "The city is pungent with the smells of filth, incense and brackish rainwater.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (6, "The Haunted Jungles Of Lahag", "Flash flood, Save vs Paralysis per hour.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (6, "The Topaz Isles", "The mother of storms sits in the sky above, tossing the sea into a tantrum. Ships caught in the storm have a 4/6 chance of sinking or being beached.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (6, "The Yellow City", "Canals flood, 4/6 chance of moored boats being smashed to pieces.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (7, "Ocean", "Sailing impossible. Movement by oar at 1/3 rate due to fatigue.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (8, "Ocean", "Sailing movement rate reduced to 1/3 normal.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (9, "Ocean", "Sailing movement rate reduced to 1/2 normal.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (10, "Ocean", "Sailing movement rate reduced to 2/3 normal.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (11, "Ocean", "Normal sailing movement rate.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (12, "Ocean", "Sailing movement rate increated by 1/3.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (13, "Ocean", "Sailing movement rate increated by 1/2.");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (14, "Ocean", "Sailing movement rate doubled. See **Near Gales**");

INSERT INTO weatherDetails (weatherTypeId, location, details)
VALUES (15, "Ocean", "Sailing movement rate tripled. **See gales and storms**");

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