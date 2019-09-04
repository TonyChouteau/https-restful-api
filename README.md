<h2> GO API with SSL </h2>

This project show how to make a GO API with HTTPS access. 
The certificate is made by Let's Encrypt so web browser doesn't send Security warning.

<h4> 1) Add certbot to the system </h4>

```bash
git clone https://github.com/certbot/certbot
cd certbot
./certbot-auto --help
```

OR follow : 
https://certbot.eff.org/lets-encrypt/ubuntubionic-other

<h4> 2) Create a SSL Certificate with certbot</h4>

```bash
./certbot-auto certonly --preferred-challenges http-01 -d www.domain.com
```

OR (if you followed the installation using the certbot website) :

```bash
certbot certonly --preferred-challenges http-01 -d www.domain.com Saving debug log to /var/log/letsencrypt/letsencrypt.log
```

<h4> 3) Create you GO API </h4>

Using the "net/http" library to server on the desired port.

```go
PORT := "8081"
http.ListenAndServeTLS(":"+PORT, "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", nil)
```

and (to use gin-gonic) :

```go
PORT := "8081"
router := gin.Default()
http.ListenAndServeTLS(":"+PORT, "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", router)
```

<h2> Have Fun </h2>
