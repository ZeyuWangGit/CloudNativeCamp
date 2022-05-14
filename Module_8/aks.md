az login

az aks get-credentials --resource-group CloudNative --name akscluster

docker build -t httpserver:v0.1 -f Dockerfile .