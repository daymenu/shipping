namesapce:
	sudo kubectl create -f namespace.yaml
build:
	docker build ../user -t shipping-user
	docker build ../container -t shipping-container
	docker build ../vessel -t shipping-vessel
	docker build ../api -t shipping-api
create:
	sudo kubectl create -f mariadb-deployment.yaml
	sudo kubectl create -f mariadb-svc.yaml
delete:
	sudo kubectl delete -f mariadb-deployment.yaml
	sudo kubectl delete -f mariadb-svc.yaml
show:
	sudo kubectl get deployment --namespace=shipping
	sudo kubectl get service --namespace=shipping