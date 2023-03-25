all: run

run: server_run

server_run:
	tilix -e 'bash -c "nats-streaming-server -cid prod | make publish_run | make subscribe_run; exec bash"'

publish_run:
	tilix -e 'bash -c "go run publish/publish.go; exec bash"'

subscribe_run:
	tilix -e 'bash -c "go run main_sub.go; exec bash"'

kill:
	pgrep nats-streaming- | xargs kill -KILL
	pgrep publish | xargs kill -KILL
	pgrep main_sub | xargs kill -KILL

postgre:
	psql -c "CREATE DATABASE wildberries" 
	psql -d wildberries -c "\i /home/andy/work/src/db/order.sql"
	psql -d wildberries -c "\d"
	psql -d wildberries -c "\d order_id"

del_db:
	psql -d wildberries -c "DROP TABLE order_id"
	psql -c "DROP DATABASE wildberries"

del_table:
	psql -d wildberries -c "DROP TABLE order_id"

table:
	psql -d wildberries -c "\i /home/andy/work/src/db/order.sql"

show_table:
	psql -d wildberries -c "SELECT * FROM order_id"

reset_table:
	psql -d wildberries -c "DROP TABLE order_id"
	psql -d wildberries -c "\i /home/andy/work/src/db/order.sql"
