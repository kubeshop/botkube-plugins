Title: Discord for Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/installation/discord/cloud

Markdown Content:
*   Go to Botkube Cloud [Web App](https://app.botkube.io/) and create a new instance.
    
    You can do it by clicking "Create an Instance" button on Home Page or under this link [Create an Instance](https://app.botkube.io/instances/add)
    
*   Fill in the `Instance Display Name` and click `Next` button.
    
    ![Image 1: Instance Display Name](https://docs.botkube.io/assets/images/discord_instance_display_name-b35605d19eef1ecc93de54d6eefacae5.png)
    
*   Click `Add platform` dropdown, and select `Discord` option. ![Image 2: Select Platform](https://docs.botkube.io/assets/images/discord_platform_select-aac36ca4e34549bef88cc00b3603f4ac.png)
    
*   Create Botkube app at your Discord Server
    
    Reach [https://discordapp.com/developers/applications](https://discordapp.com/developers/applications).
    
    ![Image 3: discord_applications_portal](https://docs.botkube.io/assets/images/discord_applications_portal-a4e1b45cb3df4a271cbd599ec9f3b7ab.png)
    
*   Create a "New Application" named Botkube and add a bot named **Botkube** into the Application.
    
    ![Image 4: discord_create_new](https://docs.botkube.io/assets/images/discord_create_new-ba9152ffe6f7be4f64af374d836c7062.png)
    
*   Copy the Application **APPLICATION ID**
    
    ![Image 5: discord_copy_client_id](https://docs.botkube.io/assets/images/discord_copy_application_id-bf48ff3b0d9dc613c35d92dc287bd305.png)
    
    and paste it in the `BotID` field in the form.
    
    ![Image 6: bot_id_form](https://docs.botkube.io/assets/images/discord_bot_id_form-a9a0d728ad26361d5454b3eac4af8838.png)
    
*   Add a description - `Botkube helps you monitor your Kubernetes cluster, debug critical deployments and gives recommendations for standard practices by running checks on the Kubernetes resources.`.
    
    Set the Botkube icon (Botkube icon can be downloaded from [this link](https://github.com/kubeshop/botkube/blob/main/branding/logos/botkube-color-192x192.png)).
    
    Click on Save Changes to update the Bot.
    
*   Now, reach the **Bot** page and Click **Add Bot** to add a Discord Bot to your application.
    
    ![Image 7: discord_add_bot](https://docs.botkube.io/assets/images/discord_add_bot-867c43f73a079d08996072d3261d2fbc.png)
    
*   After Bot creation, now you can see a bot is added to your application. Click on the **Reset Token** button.
    
    ![Image 8: discord_bot_created](https://docs.botkube.io/assets/images/discord_bot_created-845172424d2066002bff223d9a3afd36.png)
    
*   Copy the Token and paste it in `Token` field the form.
    
    ![Image 9: discord_token_form](https://docs.botkube.io/assets/images/discord_token_form-0885aae7fa9636d7f896aae91ee10cdf.png)
    
*   Go to the **OAuth2** page. Generate the URL with suitable permissions using the **OAuth2 URL Generator** available under the OAuth2 section to add bot to your Discord server.
    

the generated URL contains **YOUR\_CLIENT\_ID**, Scope and permission details.

    https://discord.com/api/oauth2/authorize?client_id={YOUR_CLIENT_ID}&permissions={SET_OF_PERMISSIONS}&scope=bot
