import psycopg2
from psycopg2 import sql

# Connection URL string
connection_url = "postgresql://sctrack_pg_user:7jAZPc5kmiqafsK37zc5XA6LFx8XE9H3@dpg-ck19417dorps73c4k42g-a.oregon-postgres.render.com/sctrack_pg"

try:
    # Connect to the database using the URL
    conn = psycopg2.connect(connection_url)

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
