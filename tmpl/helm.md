# cilium

``` shell
helm --debug -v=100  install cilium cilium/cilium --version 1.8.1 \
    --namespace kube-system \
    --set global.debug.enabled=true \
    --set global.nodeinit.enabled=true \
    --set global.bpf.waitForMount=true \
    --set global.kubeProxyReplacement=partial \
    --set global.hostServices.enabled=false \
    --set global.externalIPs.enabled=true \
    --set global.nodePort.enabled=true \
    --set global.hostPort.enabled=true \
    --set global.pullPolicy=IfNotPresent \
    --set config.ipam=kubernetes  &>cilium.helm
```
