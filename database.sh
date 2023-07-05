#!/bin/bash

# Checks if PostgreSQL service is running
check_postgres_service() {
  pg_isready > /dev/null 2>&1
}

# Starts PostgreSQL service
start_postgres_service() {
  sudo service postgresql start
}

# Creates the chess database if it doesn't exist
create_database() {
  psql -c "CREATE DATABASE IF NOT EXISTS chess;"
}

# Creates the users table if it doesn't exist
create_table() {
  psql -d chess -c "
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL PRIMARY KEY,
      username VARCHAR(255),
      password VARCHAR(255)
    );
  "
}

# Check if PostgreSQL service is running
if ! check_postgres_service; then
  echo "PostgreSQL service is not running. Starting the service..."
  start_postgres_service
fi

# Create the chess database
echo "Creating 'chess' database..."
create_database

# Create the users table
echo "Creating 'users' table..."
create_table

echo "Setup complete."
