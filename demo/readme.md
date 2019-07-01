# Example CLI
Simple example to show how this library is used in a CLI context.

## Prepare configs
 * Create a Pintrest app -> https://developers.pinterest.com/apps/ 
    * click on "Create app" and fill the form. This will get you the App ID and the App Secret. Keep those from the public. 
 * Rename `credentials.go.sample` to `credentials.go.sample`.
 * Enter your App ID ("APPID") and App Secret ("APPSECRET"). Keep the "  
    ``` 
        return credentials{
            AppId:"APPID",
            AppSecret:"APPSECRET",
        }
    ```
 
