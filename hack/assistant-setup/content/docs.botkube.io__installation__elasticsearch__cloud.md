Title: Elasticsearch for Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/installation/elasticsearch/cloud

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Installation](https://docs.botkube.io/)
*   [Elasticsearch](https://docs.botkube.io/installation/elasticsearch/)
*   Elasticsearch for Botkube Cloud

Version: 1.11

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   Botkube Cloud account which you can create [here](https://app.botkube.io/) for free.

Create a Botkube Cloud Instance with Elasticsearch[​](#create-a-botkube-cloud-instance-with-elasticsearch "Direct link to Create a Botkube Cloud Instance with Elasticsearch")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and create a new instance.
    
    You can do it by clicking "Create an Instance" button on Home Page or under this link [Create an Instance](https://app.botkube.io/instances/add)
    
2.  Fill in the `Instance Display Name` and click `Next` button.
    
    ![Image 1: Instance Display Name](https://docs.botkube.io/assets/images/els_instance_display_name-b35605d19eef1ecc93de54d6eefacae5.png)
    
3.  Click `Add platform` dropdown, and select `Elasticsearch` option.
    
    ![Image 2: Select Platform](https://docs.botkube.io/assets/images/els_platform_select-aac36ca4e34549bef88cc00b3603f4ac.png)
    
4.  Fill in all required data in the form
    
    ![Image 3: Form](https://docs.botkube.io/assets/images/els_form-940490b8840d8a700b57b3a803249bd9.png)
    
5.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 4: Plugins](https://docs.botkube.io/assets/images/els_add_plugins-dbf20e334cdca6198e7d9b0f8c68847f.png)
    
6.  Include optional `default aliases` and `default actions` and click `Create` button to create Botkube Cloud instance.
    
    ![Image 5: Create](https://docs.botkube.io/assets/images/els_create-4b637edb5bec18e1e53cf632d8bc6087.png)
    
    note
    
    If you don't include other platforms which use `Executor` plugins we recommend `default aliases` and `default actions` options unchecked
    
7.  Follow the instructions in the summary page to deploy Botkube into your environment.
    
    ![Image 6: Summary](https://docs.botkube.io/assets/images/els_summary-bfdc3ff0af6735b41a17d7219fd6b6f0.png)
    

Clean up[​](#clean-up "Direct link to Clean up")
------------------------------------------------

### Remove Botkube from Kubernetes cluster[​](#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 7: Delete](https://docs.botkube.io/assets/images/els_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
    

[Previous Elasticsearch for self-hosted Botkube](https://docs.botkube.io/installation/elasticsearch/self-hosted)[Next Outgoing webhook](https://docs.botkube.io/installation/webhook/)
