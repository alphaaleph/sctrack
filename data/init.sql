
-- \c postgres

-- Drop the sctrack-pg database if it exists
-- DROP DATABASE IF EXISTS sctrack-pg;

-- Create a new sctrack-pg database
-- CREATE DATABASE sctrack-pg;

-- Connect to the sctrack-pg database
-- \c sctrack-pg;

-- Create extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create enumeration type if it doesn't exist
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'action_enum') THEN
            CREATE TYPE action_enum AS ENUM ('completed', 'delivery', 'pickup', 'refuel', 'other');
        END IF;
    END $$;

-- Create tables

-- This table stores carrier information
CREATE TABLE IF NOT EXISTS carrier (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    telephone TEXT
);

-- This table stores enumeration values for actions
CREATE TABLE IF NOT EXISTS action_enum_table (
    action_value action_enum PRIMARY KEY
);

-- This table stores tasks
CREATE TABLE IF NOT EXISTS todos (
    uuid UUID DEFAULT uuid_generate_v4(),
    created TIMESTAMP DEFAULT NOW(),
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false,
    carrier_id TEXT, -- Changed back to TEXT
    action action_enum NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (carrier_id) REFERENCES carrier (id) ON DELETE CASCADE
);

-- This table stores journal events associated with tasks
CREATE TABLE IF NOT EXISTS journal (
    uuid UUID,
    index SERIAL,
    event TEXT NOT NULL,
    todo_uuid UUID NOT NULL,
    PRIMARY KEY (uuid, index),
    FOREIGN KEY (todo_uuid) REFERENCES todos (uuid) ON DELETE CASCADE
);

-- Populate the enumeration table
INSERT INTO action_enum_table (action_value) VALUES
    ('completed'),
    ('delivery'),
    ('pickup'),
    ('refuel'),
    ('other');

-- Generate dummy data for carrier table
INSERT INTO carrier (id, name, telephone) VALUES
    ('c0001', 'ABC Logistics', '123-456-7890'),
    ('c0002', 'XYZ Transport', '987-654-3210'),
    ('c0003', 'QQ QuickShip', '555-123-4567');

-- Insert data into the todos and journal tables, matching UUIDs and using auto-generated index
-- Insert dummy tasks associated with carriers and journal events
INSERT INTO todos (carrier_id, description, completed, action) VALUES
    ((SELECT id FROM carrier WHERE name = 'ABC Logistics'), 'Deliver to KS', false, 'delivery'),
    ((SELECT id FROM carrier WHERE name = 'QQ QuickShip'), 'Fill tank at I25', true, 'refuel'),
    ((SELECT id FROM carrier WHERE name = 'QQ QuickShip'), 'Collect load from PA', false, 'pickup');

-- Insert data into the journal table, matching UUIDs with todos and using auto-generated index
-- Insert journal events associated with tasks
INSERT INTO journal (uuid, event, todo_uuid) VALUES
    ((SELECT uuid FROM todos WHERE description = 'Deliver to KS'), '{"time":"2023-08-28T15:38:05.754673029Z", "action":"refuel"}', (SELECT uuid FROM todos WHERE description = 'Deliver to KS')),
    ((SELECT uuid FROM todos WHERE description = 'Deliver to KS'), '{"time":"2023-08-27T15:38:05.754673029Z", "action":"delivery"}', (SELECT uuid FROM todos WHERE description = 'Deliver to KS')),
    ((SELECT uuid FROM todos WHERE description = 'Fill tank at I25'), '{"time":"2023-07-27T15:38:05.754673029Z", "action":"delivery"}', (SELECT uuid FROM todos WHERE description = 'Fill tank at I25')),
    ((SELECT uuid FROM todos WHERE description = 'Fill tank at I25'), '{"time":"2023-06-28T15:38:05.754673029Z", "action":"refuel"}', (SELECT uuid FROM todos WHERE description = 'Fill tank at I25')),
    ((SELECT uuid FROM todos WHERE description = 'Collect load from PA'), '{"time":"2023-04-27T15:38:05.754673029Z", "action":"pickup"}', (SELECT uuid FROM todos WHERE description = 'Collect load from PA'));
