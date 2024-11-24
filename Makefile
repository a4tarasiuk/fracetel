.DEFAULT_GOAL := help
.PHONY: checkout-master

checkout-master:
		git fetch && git checkout master && git pull origin master
