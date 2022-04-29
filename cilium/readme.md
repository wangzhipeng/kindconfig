
kind create cluster --name cilium --config ./kind-config.yaml


cilium install

cilium hubble enable --ui


helm install traefik traefik/traefik



kubectl patch deploy traefik -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'



# hubble-ui 
项目是一个包含前后端的网页端 
hubble-ui          quay.io/cilium/hubble-ui:v0.8.5@sha256:4eaca1ec1741043cfba6066a165b3bf251590cf4ac66371c4f63fbed2224ebb4: 1
hubble-ui          quay.io/cilium/hubble-ui-backend:v0.8.5@sha256:2bce50cf6c32719d072706f7ceccad654bfa907b2745a496da99610776fe31ed: 1
hubble-ui          docker.io/envoyproxy/envoy:v1.18.4@sha256:e5c2bb2870d0e59ce917a5100311813b4ede96ce4eb0c6bfa879e3fbe3e83935: 1



# hubble 
与hubble-ui 功能一样的命令行端(cilium hubble port-forward 开启hubble-relay 端口映射后才能使用 )

# cilium 
项目包含以下镜像
cilium             quay.io/cilium/cilium:v1.11.3@sha256:cb6aac121e348abd61a692c435a90a6e2ad3f25baa9915346be7b333de8a767f: 3
cilium-operator    quay.io/cilium/operator-generic:v1.11.3@sha256:5b81db7a32cb7e2d00bb3cf332277ec2b3be239d9e94a8d979915f4e6648c787: 1
hubble-relay       quay.io/cilium/hubble-relay:v1.11.3@sha256:7256ec111259a79b4f0e0f80ba4256ea23bd472e1fc3f0865975c2ed113ccb97: 1

# cilium-cli 
cilium 工具的命令行，复杂组件的安装配置


