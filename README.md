# <img src="images/gcp.png" width="26"> alfred-gcp-workflow

An Alfred workflow to quickly open GCP services in your browser or search GCP resources with ease.

## Installation

1. Make sure you have the [Google Cloud Cli](https://cloud.google.com/sdk/docs/install) installed and authenticated.
2. Download the latest release from the [releases page](https://github.com/dineshgowda24/alfred-gcp-workflow/releases)
3. Unzip the downloaded file and open the `.alfredworkflow` file to import it into Alfred app.
4. Set the `ALFRED_GCP_WORKFLOW_GCLOUD_PATH` environment variable to point to your gcloud executable.
(This is required for the workflow to function properly.)

## Usage

1. Open Alfred and type `gcp` to see the available services and commands.
2. The home page will display useful links to Google Cloud:
![Home Page](images/docs/home.png)
3. Type `gcp` followed by a service name. For example, `gcp compute` will show the Compute Engine service.
4. If a service has 🗂️ in its subtitle, press  press <kbd>Tab</kbd> to autocomplete into the subservices section — to navigate to redis inside memorystore.
5. You can filter subservices directly by typing their name. For example, `gcp compute instances` will show Compute Engine instances.
6. If a subservice has 🔍⚡️ in its subtitle, it supports **resource search**. For example, after typing `gcp compute` you can <kbd>Tab</kbd> into `instances` to list them.
7. The workflow automatically uses your current active gcloud configuration.
To switch, run in the terminal:
```bash
gcloud config configurations activate <configuration-name>
```
and the workflow will pick up the new active configuration immediately.

## Contributing

Please read the [contributing guidelines](CONTRIBUTING.md) for details on how to set up your environment and submit changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for full license information.

## Acknowledgements 🙏

This workflow is inspired by the amazing [aws-alfred-workflow](https://github.com/rkoval/alfred-aws-console-services-workflow).
As a past user of that workflow before switching to GCP, I wanted to create a similar experience for GCP users.
Huge thanks to the original author for the idea and inspiration — without which this workflow wouldn't exist!