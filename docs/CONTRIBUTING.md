# Contributing

Thanks for your interest in contributing! 🎉  
This document will guide you through setting up, understanding the design, and adding new features to the workflow.

---

## Table of Contents
- [Contributing](#contributing)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
    - [Installation](#installation)
  - [Finding Something to Work On 💡](#finding-something-to-work-on-)
  - [Design](#design)
  - [Adding a New Service or Subservice](#adding-a-new-service-or-subservice)
  - [Adding a New Searcher](#adding-a-new-searcher)
  - [Notes](#notes)
  - [Thank You! 🙏](#thank-you-)

---

## Requirements

- [Go](https://golang.org/doc/install) (version 1.22.4 or higher)

### Installation

1. Fork this repository.
2. Clone your fork locally:
   ```bash
   git clone github.com:<your-username>/alfred-gcp-workflow.git
   ```
3. Build the workflow:
   ```bash
   go build -o alfred-gcp-workflow
   ```
4. Update the Script Filter in Alfred:
    - Open Alfred GCP Workflow in Alfred.
    - Click on the **Script Filter** block.
    - Update the script path to point to your built binary.


## Finding Something to Work On 💡

We welcome contributions of all kinds — whether it's a small fix, a new feature, or just an idea! ✨
If you're looking for something to work on:
- Check out the [Issues tab](https://github.com/dineshgowda24/alfred-gcp-workflow/issues) on GitHub.
- If you are a first-time contributor, checkout some of the [good first issues](https://github.com/dineshgowda24/alfred-gcp-workflow/labels/good%20first%20issue)
- If you’d like to work on an issue, please leave a comment to let others know you're taking it.
- This helps avoid duplication and allows maintainers to assist you if needed!

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

1. **Update [services.yml](services.yml)**:
   - Add your new Service.
   - (Optional) Add Subservices under the Service.
2. **Icons**:
   - Use existing icons in the [icons](../icons) directory if available.
   - To add a new icon:
      - Place the image in [icons](../icons).
      - Reference the path inside `services.yml`.
   - Official GCP service icons can be found [here](https://cloud.google.com/icons).
3. **Sort Services Alphabetically**:
   - To keep `services.yml` clean and consistent, sort services by `id` after your changes.
   - You can use the `yq` tool for this:
     ```bash
     brew install yq
     yq '.|= sort_by(.id)' services.yml > tmp.yml && mv tmp.yml services.yml
     ```
   - This ensures all services remain consistently ordered.

## Adding a New Searcher
1. Create a new file in the [gcloud](gcloud/) folder:
   - The file should define the gcloud command to list resources for a subservice.
   - Only include necessary fields, and avoid listing sensitive information.
   - Example: See [firestore.go](gcloud/filestore.go).
2. Implement a new Searcher:
   - Create a struct that implements the `Searcher` interface.
   - Place it inside the [searchers](searchers/) folder.
3. Register the Searcher:
   - Update [searcher.go](searchers/searcher.go).
   - Use the key format: `service_id/subservice_id` when registering the searcher.

## Notes

- Keep contributions focused and atomic (one PR per logical change).
- Follow the existing code structure and style for consistency.
- Tests(if possible) and documentation updates are highly appreciated!


## Thank You! 🙏

*Every contribution, no matter how small, really helps us improve this workflow. We appreciate your time and effort—thank you so much!*

---