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


