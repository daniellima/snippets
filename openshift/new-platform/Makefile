.PHONY: build debug deploy-platform deploy-app s2i debug-s2i run

platform_name=snippet-new-platform
platform_image_name=local/$(platform_name)
app_image_name=$(platform_image_name)-app
openshift_registry=172.30.1.1:5000
openshift_image_stream=$(openshift_registry)/myproject/$(platform_name)
openshift_user_token=ufNKeurwZWqZNbWZx5wQ7ux7udxBjYVaE8HizZJaobE
openshift_app_name=flask-app

build:
	docker build -t $(platform_image_name) .

debug:
	docker run --rm -it $(platform_image_name) /bin/bash

deploy:
	docker tag $(platform_image_name) $(openshift_image_stream)
	docker login -u openshift -p $(openshift_user_token) $(openshift_registry)
	docker push $(openshift_image_stream)

deploy-app:
	oc delete all -l app=$(openshift_app_name)
	oc new-app $(platform_name)~. --name=$(openshift_app_name)
	oc start-build $(openshift_app_name) --from-dir=.
	oc logs bc/$(openshift_app_name) -f
	oc logs dc/$(openshift_app_name) -f
	oc expose dc/$(openshift_app_name) --port 5000
	oc expose svc/$(openshift_app_name)

app-image: build
	s2i build . $(platform_image_name) $(app_image_name)

debug-app-image:
	docker run --rm -it $(app_image_name) /bin/bash

run-app:
	docker run --rm -it -p 5000:5000 $(app_image_name)
