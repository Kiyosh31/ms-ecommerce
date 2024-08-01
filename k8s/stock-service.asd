apiVersion: apps/v1
kind: Deployment
metadata:
  name: stock-depl
spec:
  replicas: 1
  selector:
    matchLabels:
      app: stock
  template:
    metadata:
      labels:
        app: stock
    spec:
      containers:
        - name: stock
          image: stock-service
          imagePullPolicy: Never
          resources:
            limits:
              memory: 512Mi
              cpu: "1"
          env:
            - name: STOCK_SERVICE_HTTP_ADDR
              valueFrom:
                secretKeyRef:
                  name: secrets
                  key: STOCK_SERVICE_HTTP_ADDR
            - name: DB_CONNECTION_LINK
              valueFrom:
                secretKeyRef:
                  name: secrets
                  key: DB_CONNECTION_LINK
            - name: STOCK_SERVICE_DATABASE_NAME
              valueFrom:
                secretKeyRef:
                  name: secrets
                  key: STOCK_SERVICE_DATABASE_NAME
            - name: STOCK_SERVICE_DATABASE_COLLECTION
              valueFrom:
                secretKeyRef:
                  name: secrets
                  key: STOCK_SERVICE_DATABASE_COLLECTION
---
apiVersion: v1
kind: Service
metadata:
  name: stock-srv
spec:
  type: ClusterIP
  selector:
    app: stock
  ports:
    - name: http
      protocol: TCP
      port: 3004
      targetPort: 3004
