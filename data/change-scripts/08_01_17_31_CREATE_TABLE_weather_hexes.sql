CREATE TABLE weatherhexes (
	id INTEGER,
	weather_id INTEGER,
	marked BOOLEAN NOT NULL CHECK (marked IN (0, 1)),
	one INTEGER,
	two INTEGER,
	three INTEGER,
	four INTEGER,
	five INTEGER,
	six INTEGER,
	PRIMARY KEY (id AUTOINCREMENT),
	FOREIGN KEY (weather_id) REFERENCES weathertypes (id)
);

