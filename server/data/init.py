import psycopg2
# from psycopg2 import sql
import argparse

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Connect to PostgreSQL database using a connection URL.")
    parser.add_argument("connection_url", help="PostgreSQL connection URL")
    args = parser.parse_args()

    try:
        # Connect to the database using the URL
        conn = psycopg2.connect(args.connection_url)

        # Create a cursor
        cur = conn.cursor()

        # Read and execute the SQL statements from the init.sql file
        with open('init.sql', 'r') as init_file:
            statements = init_file.read()
            cur.execute(statements)

        # Commit the changes
        conn.commit()

    except psycopg2.Error as e:
        print("Error:", e)
    finally:
        # Close the cursor and connection
        cur.close()
        conn.close()
