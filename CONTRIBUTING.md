# Contributing

Thanks for your interest in contributing! 🎉  
This document will guide you through setting up, understanding the design, and adding new features to the workflow.

---

## Table of Contents
- [Contributing](#contributing)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
    - [Installation](#installation)
  - [Design](#design)
  - [Adding a New Service or Subservice](#adding-a-new-service-or-subservice)
  - [Adding a New Searcher](#adding-a-new-searcher)
  - [Notes](#notes)
  - [Thank You! ♥️](#thank-you-️)

---

## Requirements

- [Go](https://golang.org/doc/install) (version 1.22.4 or higher)

### Installation

1. **Fork** this repository.
2. **Clone** your fork locally:
   ```bash
   git clone github.com:<your-username>/alfred-gcp-workflow.git
   ```
3. Build the workflow:
   ```bash
   go build -o alfred-gcp-workflow
   ```
4. Update the Script Filter in Alfred:
    - Open your Alfred GCP Workflow.
    - Click on the **Script Filter** block.
    - Update the script path to point to your built binary.

## Design

The workflow is organized around three main concepts:
1. **Service**
   - Represents a GCP service (e.g., Compute Engine, Cloud Storage).
   - Defined in [services.yml](services.yml).
2. **Subservice**
   - Represents a specific resource within a Service (e.g., a VM instance, a Storage bucket).
   - Listed under each Service in [services.yml](services.yml).
3. **Searcher**
   - Defines how to fetch resources for a Subservice using a gcloud command (e.g., gcloud compute instances list).
   - Implemented in [searchers/searcher.go](searchers/searcher.go).

## Adding a New Service or Subservice

1. Update [services.yml](services.yml):
   - Add your new Service.
   - (Optional) Add Subservices under the Service.
2. Icons:
   - Use existing icons in the [images/](images/) folder if available.
   - To add a new icon:
      - Place the image in [images/](images/).
      - Reference the path inside services.yml.
   - Official GCP service icons can be found [here](https://cloud.google.com/icons).

## Adding a New Searcher
1. Create a new file in the [gcloud/](gcloud/) folder:
   - The file should define the gcloud command to list resources for a Subservice.
   - Only include fields that are absolutely necessary — avoid listing sensitive information.
   - Example: See [firestore.go](gcloud/filestore.go).
2. Implement a new Searcher:
   - Create a struct that implements the Searcher interface.
   - Place it inside the [searchers/](searchers/) folder.
3. Register the Searcher:
   - Update [searcher.go](searchers/searcher.go).
   - Use the key format: `service_id/subservice_id` when registering the searcher.

## Notes

- Keep contributions focused and atomic (one PR per logical change).
- Follow the existing code structure and style for consistency.
- Tests(if possible) and documentation updates are highly appreciated!


## Thank You! ♥️

*Every contribution, big or small, helps improve the Alfred GCP Workflow. 
Excited to see what you build! 🚀*

---