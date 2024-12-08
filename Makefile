gen_pb:
	protoc \
        	--go_out=./pkg/protocol/pb \
        	--go_opt=module=gchat/pkg/protocol/pb \
        	--go-grpc_out=./pkg/protocol/pb --go-grpc_opt=module=gchat/pkg/protocol/pb \
        	pkg/protocol/proto/*.proto

gcp_deploy_user:
	gcloud run deploy \
	--image  us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/user-service:latest \
	--platform managed \
    --allow-unauthenticated \
    --region us-central1 \
    --network projects/cs739-gchat/global/networks/default \
    --subnet projects/cs739-gchat/regions/us-central1/subnetworks/redis \
    --set-env-vars MYSQLADDR=/cloudsql/cs739-gchat:us-central1:cs739-gchat,MYSQLUSER=local-user,MYSQLPASSWORD=jason123456,DB_CONNECTOR=unix,REDISHOST=10.13.247.19,REDISPORT=6379

build_connect:
	docker build -t connect-service:latest --build-arg CMD_PATH=cmd/connect .
tag_connect: build_connect
	docker tag connect-service:latest us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/connect-service:latest
push_connect: tag_connect
	docker push us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/connect-service:latest

build_logic:
	docker build -t logic-service:latest --build-arg CMD_PATH=cmd/logic .
tag_logic: build_logic
	docker tag logic-service:latest us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/logic-service:latest
push_logic: tag_logic
	docker push us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/logic-service:latest

build_user:
	docker build -t user-service:latest --build-arg CMD_PATH=cmd/user .
tag_user: build_user
	docker tag user-service:latest us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/user-service:latest
push_user: tag_user
	docker push us-central1-docker.pkg.dev/cs739-gchat/gchat-repo/user-service:latest


gcp_cluster_cred:
	gcloud container clusters get-credentials gchat-connect --zone us-central1

gcp_push_helm:
	helm push gchat-1.4.0.tgz oci://us-central1-docker.pkg.dev/cs739-gchat/helm-repo

gcp_deploy_helm:
	helm upgrade gchat ./chart --version 1.11.0