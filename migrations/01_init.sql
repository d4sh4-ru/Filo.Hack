-- Create Resident table
CREATE TABLE resident (
    resident_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    middle_name VARCHAR(50),
    age INT,
    house_number VARCHAR(10) NOT NULL,
    entrance INT,
    apartment INT
);

-- Create Interest table
CREATE TABLE interest (
    interest_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create EventType table
CREATE TABLE event_type (
    event_type_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create Event table
CREATE TABLE event (
    event_id SERIAL PRIMARY KEY,
    event_name VARCHAR(100) NOT NULL,
    event_date DATE NOT NULL,
    address VARCHAR(200) NOT NULL,
    event_type_id INT NOT NULL,
    organizer_resident_id INT NOT NULL,
    FOREIGN KEY (event_type_id) REFERENCES event_type(event_type_id),
    FOREIGN KEY (organizer_resident_id) REFERENCES resident(resident_id)
);

-- Create EventShow table
CREATE TABLE event_show (
    show_id SERIAL PRIMARY KEY,
    event_type_id INT NOT NULL,
    resident_id INT NOT NULL,
    FOREIGN KEY (event_type_id) REFERENCES event_type(event_type_id),
    FOREIGN KEY (resident_id) REFERENCES resident(resident_id)
);

-- Create EventParticipation table
CREATE TABLE event_participation (
    resident_id INT NOT NULL,
    event_id INT NOT NULL,
    PRIMARY KEY (resident_id, event_id),
    FOREIGN KEY (resident_id) REFERENCES resident(resident_id),
    FOREIGN KEY (event_id) REFERENCES event(event_id)
);

-- Create ResidentInterest junction table for many-to-many relationship
CREATE TABLE resident_interest (
    resident_id INT NOT NULL,
    interest_id INT NOT NULL,
    PRIMARY KEY (resident_id, interest_id),
    FOREIGN KEY (resident_id) REFERENCES resident(resident_id),
    FOREIGN KEY (interest_id) REFERENCES interest(interest_id)
);

-- Create indexes for better performance
CREATE INDEX idx_event_type ON event(event_type_id);
CREATE INDEX idx_event_organizer ON event(organizer_resident_id);
CREATE INDEX idx_show_resident ON event_show(resident_id);
CREATE INDEX idx_show_event_type ON event_show(event_type_id);
