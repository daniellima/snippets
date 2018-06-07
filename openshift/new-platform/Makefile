.PHONY: build debug deploy s2i debug-s2i run

build:
	docker build -t snippets/ubuntu-sti .

debug:
	docker run --rm -it snippets/ubuntu-sti /bin/bash

deploy:
	docker tag snippets/ubuntu-sti 172.30.1.1:5000/myproject/ubuntu-sti
	docker login -u openshift -p wrdo4oSTBvCZghehDrZWWWoPKO58SOa9BdtZfZRDbjI 172.30.1.1:5000
	docker push 172.30.1.1:5000/myproject/ubuntu-sti

app:
	oc delete all -l app=flask-app
	oc new-app ubuntu-sti~. --name=flask-app
	oc start-build flask-app --from-dir=.
	oc logs bc/flask-app -f
	oc logs dc/flask-app -f

s2i: build
	s2i build . snippets/ubuntu-sti snippets/ubuntu-sti-builded

debug-s2i:
	docker run --rm -it snippets/ubuntu-sti-builded /bin/bash

run:
	docker run --rm -it -p 8000:5000 snippets/ubuntu-sti-builded
