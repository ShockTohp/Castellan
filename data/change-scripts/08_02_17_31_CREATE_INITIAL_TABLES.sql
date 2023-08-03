PRAGMA foreign_keys = ON;
CREATE TABLE campaigns (
	id INTEGER,
	guildId INTEGER,
	name TEXT,
	PRIMARY KEY (id AUTOINCREMENT)
);

CREATE TABLE weatherSystems (
	id INTEGER,
	systemName TEXT,
	resolutionType TEXT,
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


CREATE TABLE weatherMarker (
	id INTEGER,
	campaignId INTEGER,
	weatherSystemId INTEGER,
    hexId INTEGER,
    weatherDate TEXT,
	PRIMARY KEY (id AUTOINCREMENT),
	FOREIGN KEY (weatherSystemId) REFERENCES weatherSystems (id),
	FOREIGN KEY (campaignId) REFERENCES campaigns (id)
);

