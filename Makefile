RUN_TF = docker-compose run tf

set-token:
	echo 'credentials "app.terraform.io" { token = "${TF_TOKEN}" }' > ./terraform/.terraformrc

init:
	${RUN_TF} init

init-upgrade:
	${RUN_TF} init -upgrade

plan:
	${RUN_TF} plan

apply:
	${RUN_TF} apply

apply-refresh:
	${RUN_TF} apply -refresh-only

fmt:
	${RUN_TF} fmt -recursive

zip:
	zip -r terraform/modules/gcf/rss bots/rss -x "*.env"
