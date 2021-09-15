
## Quick Start

### Prerequisites

- [Go](https://golang.org/) version v1.15+

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) version v1.19+

- If not deploy EdgeX Foundry before, run

  ```shell
  kubectl apply -f config/samples/edgex.yaml
  ```

  

### Install

1. install our CRDs on prepared kubernetes cluster

   ```shell
   make install
   ```

2. Start controller 

   ```shell
   make run ENABLE_WEBHOOKS=false
   ```

3. Install device service

   ```shell
   kubectl apply -f config/samples/device-virtual.yaml
   ```

   

4. create CRD device, devicemodel and deviceaccess

   ```
   kubectl apply -f config/samples/devices_v1alpha3_deviceaccess.yaml -f config/samples/devices_v1alpha3_devicemodel.yaml -f devices_v1alpha3_device.yaml
   ```

   

5. check device status, 

   ```shell
   $ kubectl describe device random-boolean-device
   ...
   Spec:
     Commands:
       Name:   Bool
       Value:  true
     Device Access Ref:
       API Version:  devices.kubeedge.io/v1alpha3
       Kind:         DeviceAccess
       Name:         random-boolean-device-visitor
     Model:          devicemodel-sample
     Protocol:
       Address:  device-virtual-bool-01
       Name:     other
       Type:     300
     Service:
       Command:         edgex-core-command
       Device Service:  device-virtual
       Metadata:        edgex-core-metadata
   Status:
     Commands:
       Name:        Bool
       Read Write:  RW
       Value:       true
     Ready:         true
   ...
   ```

   