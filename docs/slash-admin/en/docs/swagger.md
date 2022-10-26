## Use swagger

> Environment setting

[go-swagger](https://goswagger.io/install.html)

> In the root of project run slash-admin/

```shell
swagger generate spec --output=./core.yml --scan-models

swagger serve --no-open -F=swagger --port 36666 core.yaml
```

![pic](../../assets/swagger.png)

> Get Token
> Firstly, log in the system and press F12 to get authorization from any request

![pic](../../assets/get_token.png)

> Copy to swagger

![pic](../../assets/swagger_authority.png)