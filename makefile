
EMAIL ?= $(shell bash -c 'read -p "Email: " email; echo $$email')

subscribe:
	@clear
	curl -X POST http://localhost:9000/api/v1/user/subscribe -d '{"email": "'$(EMAIL)'"}' -H 'Content-Type: application/json'

unsubscribe:
	@clear
	curl -X GET http://localhost:9000/api/v1/user/unsubscribe?email=$(EMAIL) -H 'Content-Type: application/json'

publish:
	@clear
	curl -X GET http://localhost:9000/api/v1/user/publish -H 'Content-Type: application/json'
