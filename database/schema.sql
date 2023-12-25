--Lifts table schema
Create Table Lifts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(255) NOT NULL,
	compound BOOL,
	`upper` BOOL,
	`lower` BOOL
);

--Workouts table schema
CREATE TABLE Workouts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    `location` VARCHAR(255),
    notes VARCHAR(255),
    `date` DATE DEFAULT CURRENT_DATE
);

--Cardio table schema
Create Table Cardio (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(255) NOT NULL
);

Create Table LiftingLog (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    lift_id INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    sets INTEGER NOT NULL,
    reps INTEGER NOT NULL,
    workout_id INTEGER NOT NULL,
    FOREIGN KEY (workout_id) REFERENCES Workouts(id)
);

Create Table CardioLog (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cardio_id INTEGER NOT NULL,
    time INTEGER,
    distance INTEGER,
    workout_id INTEGER NOT NULL,
    FOREIGN KEY(workout_id) REFERENCES Workouts(id)
);
