# Creating self-signed certs on Arch Linux

You can use `openSSL` to create SSL certificates. To do so, follow the steps below:

## Geneate CA (Certificate Authority)
- You need a certificate Authority that will be used to sign the certificate.
- You start by generating the RSA Key (Private Key) for the CA

```bash
openssl genrsa -aes256 -out ca-key.pem 4096
```

## Use the generated CA Key to create a CA Certificate

```bash
openssl req -new -x509 -sha256 -days 365 -key ca-key.pem -out ca.pem
```

## View the generated CA certificate

```bash
openssl x509 -in ca.pem -text
```

## Generate an RSA (Private) Key for your certificate
- Do not protect the certificate using a passphrase as you will be uploading it to the server

```bash
openssl genrsa -out cert-key.pem 4096
```

## Create a Certificate Sign Request

```bash
openssl req -new sha256 -subj "/CN=Anything" -key cert-key.pem -out cert.csr
```

## Create a file to configure the SubjectAlternativeNames

```bash
echo "subjectAltName=DNS:your.dns.record,IP:192.168.100.19" >> extfile.cnf
```

## Generate certificate from Certificate Sign Request
```bash
openssl x509 -req -sha256 -days 360 -in cert.csr -CA ca.pem -CAkey ca-key.pem -out cert.pem -extfile extfile.cnf -CAcreateserial
```

## Combine the ca.pem (Certificate Authority) and cert.pem file into one file

```bash
cat cert.pem > fullchain.pem
cat ca.pem >> fullchain.pem
```

Now you have the certificate which is "fullchain.pem" and the certificate key "cert-key.pem" which you should upload to your server