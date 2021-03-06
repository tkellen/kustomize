# Feature Test for TestCase simple


This folder contains files for old test-case simple

## Setup the workspace

First, define a place to work:

<!-- @makeWorkplace @test -->
```bash
DEMO_HOME=$(mktemp -d)
```

## Preparation

<!-- @makeDirectories @test -->
```bash
mkdir -p ${DEMO_HOME}
```
## Execution

<!-- @build @test -->
```bash
# kustomize build ${DEMO_HOME}/in -o ${DEMO_HOME}/actual.yaml
```

## Verification


### Verification Step Expected0

<!-- @createExpected0 @test -->
```bash
cat <<'EOF' >${DEMO_HOME}/expected.yaml
apiVersion: v1
data:
  app-init.ini: |
    FOO=bar
    BAR=baz
kind: ConfigMap
metadata:
  annotations:
    note: This is a test annotation
  labels:
    app: mungebot
    org: kubernetes
    repo: test-infra
  name: test-infra-app-config-hf5424hg8g
---
apiVersion: v1
data:
  DB_PASSWORD: somepw
  DB_USERNAME: admin
kind: ConfigMap
metadata:
  annotations:
    note: This is a test annotation
  labels:
    app: mungebot
    org: kubernetes
    repo: test-infra
  name: test-infra-app-env-bh449c299k
---
apiVersion: v1
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUIwekNDQVgyZ0F3SUJBZ0lKQUkvTTdCWWp3Qit1TUEwR0NTcUdTSWIzRFFFQkJRVUFNRVV4Q3pBSkJnTlYKQkFZVEFrRlZNUk13RVFZRFZRUUlEQXBUYjIxbExWTjBZWFJsTVNFd0h3WURWUVFLREJoSmJuUmxjbTVsZENCWAphV1JuYVhSeklGQjBlU0JNZEdRd0hoY05NVEl3T1RFeU1qRTFNakF5V2hjTk1UVXdPVEV5TWpFMU1qQXlXakJGCk1Rc3dDUVlEVlFRR0V3SkJWVEVUTUJFR0ExVUVDQXdLVTI5dFpTMVRkR0YwWlRFaE1COEdBMVVFQ2d3WVNXNTAKWlhKdVpYUWdWMmxrWjJsMGN5QlFkSGtnVEhSa01Gd3dEUVlKS29aSWh2Y05BUUVCQlFBRFN3QXdTQUpCQU5MSgpoUEhoSVRxUWJQa2xHM2liQ1Z4d0dNUmZwL3Y0WHFoZmRRSGRjVmZIYXA2TlE1V29rLzR4SUErdWkzNS9NbU5hCnJ0TnVDK0JkWjF0TXVWQ1BGWmNDQXdFQUFhTlFNRTR3SFFZRFZSME9CQllFRkp2S3M4UmZKYVhUSDA4VytTR3YKelF5S24wSDhNQjhHQTFVZEl3UVlNQmFBRkp2S3M4UmZKYVhUSDA4VytTR3Z6UXlLbjBIOE1Bd0dBMVVkRXdRRgpNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUZCUUFEUVFCSmxmZkpIeWJqREd4Uk1xYVJtRGhYMCs2djAyVFVLWnNXCnI1UXVWYnBRaEg2dSswVWdjVzBqcDlRd3B4b1BUTFRXR1hFV0JCQnVyeEZ3aUNCaGtRK1YKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  tls.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCT3dJQkFBSkJBTkxKaFBIaElUcVFiUGtsRzNpYkNWeHdHTVJmcC92NFhxaGZkUUhkY1ZmSGFwNk5RNVdvCmsvNHhJQSt1aTM1L01tTmFydE51QytCZFoxdE11VkNQRlpjQ0F3RUFBUUpBRUoyTit6c1IwWG44L1E2dHdhNEcKNk9CMU0xV08rayt6dG5YLzFTdk5lV3U4RDZHSW10dXBMVFlnalpjSHVmeWtqMDlqaUhtakh4OHU4WlpCL28xTgpNUUloQVBXK2V5Wm83YXkzbE16MVYwMVdWak5LSzlRU24xTUpsYjA2aC9MdVl2OUZBaUVBMjVXUGVkS2dWeUNXClNtVXdiUHc4Zm5UY3BxRFdFM3lUTzN2S2NlYnFNU3NDSUJGM1VtVnVlOFlVM2p5YkMzTnh1WHEzd05tMzRSOFQKeFZMSHdEWGgvNk5KQWlFQWwyb0hHR0x6NjRCdUFmaktycXd6N3FNWXI5SENMSWUvWXNvV3Evb2x6U2NDSVFEaQpEMmxXdXNvZTIvbkVxZkRWVldHV2x5Sjd5T21xYVZtL2lOVU45QjJOMmc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
kind: Secret
metadata:
  annotations:
    note: This is a test annotation
  labels:
    app: mungebot
    org: kubernetes
    repo: test-infra
  name: test-infra-app-tls-6hkmhf2224
type: kubernetes.io/tls
---
apiVersion: v1
kind: Service
metadata:
  annotations:
    baseAnno: This is an base annotation
    note: This is a test annotation
  labels:
    app: mungebot
    foo: bar
    org: kubernetes
    repo: test-infra
  name: test-infra-baseprefix-mungebot-service
spec:
  ports:
  - port: 7002
  selector:
    app: mungebot
    foo: bar
    org: kubernetes
    repo: test-infra
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    baseAnno: This is an base annotation
    note: This is a test annotation
  labels:
    app: mungebot
    foo: bar
    org: kubernetes
    repo: test-infra
  name: test-infra-baseprefix-mungebot
spec:
  replicas: 2
  selector:
    matchLabels:
      app: mungebot
      foo: bar
      org: kubernetes
      repo: test-infra
  template:
    metadata:
      annotations:
        baseAnno: This is an base annotation
        note: This is a test annotation
      labels:
        app: mungebot
        foo: bar
        org: kubernetes
        repo: test-infra
    spec:
      containers:
      - env:
        - name: FOO
          valueFrom:
            configMapKeyRef:
              key: somekey
              name: test-infra-app-env-bh449c299k
        - name: BAR
          valueFrom:
            secretKeyRef:
              key: somekey
              name: test-infra-app-tls-6hkmhf2224
        - name: foo
          value: bar
        image: nginx:1.8.0
        name: nginx
        ports:
        - containerPort: 80
      - envFrom:
        - configMapRef:
            name: someConfigMap
        - configMapRef:
            name: test-infra-app-env-bh449c299k
        - secretRef:
            name: test-infra-app-tls-6hkmhf2224
        image: busybox
        name: busybox
        volumeMounts:
        - mountPath: /tmp/env
          name: app-env
        - mountPath: /tmp/tls
          name: app-tls
      volumes:
      - configMap:
          name: test-infra-app-env-bh449c299k
        name: app-env
      - name: app-tls
        secret:
          secretName: test-infra-app-tls-6hkmhf2224
EOF
```


<!-- @compareActualToExpected @test -->
```bash
```

